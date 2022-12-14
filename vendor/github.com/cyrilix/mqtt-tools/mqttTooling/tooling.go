package mqttTooling

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"os"
	"time"
)

const (
	MinTLSVersion = tls.VersionTLS12
)

func newTlsConfig(cafile string, certfile string, keyfile string) (*tls.Config, error) {

	cert, err := tls.LoadX509KeyPair(certfile, keyfile)
	if err != nil {
		zap.S().Panicf("unable to load certificate files: %v", err)
	}
	cacert, err := ioutil.ReadFile(cafile)
	if err != nil {
		return nil, fmt.Errorf("unable to load cacert file: %v", err)
	}
	cacertPool := x509.NewCertPool()
	if !cacertPool.AppendCertsFromPEM(cacert) {
		return nil, fmt.Errorf("unable to load cacert file: %v", err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      cacertPool,
		//ServerName:         "",
		InsecureSkipVerify: true,
		MinVersion:         MinTLSVersion,
	}, nil
}

func Connect(params *MqttCliParameters) (MQTT.Client, error) {
	//create a ClientOptions struct setting the Broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker(params.Broker)
	opts.SetUsername(params.Username)
	opts.SetPassword(params.Password)
	opts.SetClientID(params.ClientId)
	opts.SetAutoReconnect(true)
	opts.SetCleanSession(params.Clean)
	opts.SetDefaultPublishHandler(
		//define a function for the default message handler
		func(client MQTT.Client, msg MQTT.Message) {
			fmt.Printf("TOPIC: %s\n", msg.Topic())
			fmt.Printf("MSG: %s\n", msg.Payload())
		})
	if params.HasTLSConfig() {
		zap.S().Info("mqtt: enable x509 authentication")
		tlsConfig, err := params.TLSConfig()
		if err != nil {
			return nil, fmt.Errorf("unable to configure tls parameters: %v", err)
		}
		opts.SetTLSConfig(tlsConfig)
	}

	//create and start a client using the above ClientOptions
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("unable to connect to mqtt bus: %v", token.Error())
	}
	return client, nil
}

type MqttCliParameters struct {
	Broker   string
	Username string
	Password string
	CAFile   string
	CertFile string
	KeyFile  string
	ClientId string
	Qos      int
	Retain   bool
	Clean    bool
}

func (p *MqttCliParameters) HasTLSConfig() bool {
	return p.CAFile != "" && p.CertFile != "" && p.KeyFile != ""
}

func (p *MqttCliParameters) TLSConfig() (*tls.Config, error) {
	if p.HasTLSConfig() {
		zap.S().Info("mqtt: enable x509 authentication")
		tlsConfig, err := newTlsConfig(p.CAFile, p.CertFile, p.KeyFile)
		if err != nil {
			return nil, fmt.Errorf("unable to configure tls parameters: %v", err)
		}
		return tlsConfig, nil
	}
	return nil, errors.New("unable to configure tls connection, no certificates/key files defined")
}

func InitMqttFlagSet(parameters *MqttCliParameters) {
	setDefaultValueFromEnv(&parameters.ClientId, "MQTT_CLIENT_ID", parameters.ClientId)
	setDefaultValueFromEnv(&parameters.Broker, "MQTT_BROKER", "tcp://127.0.0.1:1883")

	flag.StringVar(&parameters.Broker, "mqtt-broker", parameters.Broker, "Broker Uri, use MQTT_BROKER env if arg not set")
	flag.StringVar(&parameters.Username, "mqtt-username", os.Getenv("MQTT_USERNAME"), "Broker Username, use MQTT_USERNAME env if arg not set")
	flag.StringVar(&parameters.Password, "mqtt-password", os.Getenv("MQTT_PASSWORD"), "Broker Password, MQTT_PASSWORD env if args not set")
	flag.StringVar(&parameters.ClientId, "mqtt-client-id", parameters.ClientId, "Mqtt client id, use MQTT_CLIENT_ID env if args not set")
	flag.IntVar(&parameters.Qos, "mqtt-qos", parameters.Qos, "Qos to pusblish message, use MQTT_QOS env if arg not set")
	flag.BoolVar(&parameters.Retain, "mqtt-retain", parameters.Retain, "Retain mqtt message, if not set, true if MQTT_RETAIN env variable is set")
	flag.BoolVar(&parameters.Clean, "mqtt-clean", parameters.Clean, "set the \"Clean session\" flag in the connect message")
	flag.StringVar(&parameters.CAFile, "mqtt-cafile", "", "CA pem file for x509 authentication")
	flag.StringVar(&parameters.CertFile, "mqtt-certfile", "", "X509 certificate pem file for x509 authentication")
	flag.StringVar(&parameters.KeyFile, "mqtt-keyfile", "", "rsa key pem file for x509 authentication")
}

func setDefaultValueFromEnv(value *string, key string, defaultValue string) {
	if os.Getenv(key) != "" {
		*value = os.Getenv(key)
	} else {
		*value = defaultValue
	}
}

type Topic string

type TopicFormatter string

func (t TopicFormatter) Apply(room string, topic string) Topic {
	return Topic(fmt.Sprintf(string(t), room, topic))
}

type Publisher interface {
	Publish(topic Topic, payload []byte) error
	PublishAsString(topic Topic, payload fmt.Stringer) error
}

type Subscriber interface {
	Subscribe(topic Topic, mh MQTT.MessageHandler)
}

type MQTTPubSub interface {
	Publisher
	Subscriber
	io.Closer
}

// NewMqttPublisher is deprecated. Use NewMQTTPubSub
func NewMqttPublisher(client MQTT.Client) *MqttPublisher {
	return &MqttPublisher{client: client}
}

func NewMQTTPubSub(client MQTT.Client) MQTTPubSub {
	return &MqttPublisher{client: client}
}

type MqttPublisher struct {
	client   MQTT.Client
	qos      byte
	retained bool
}

func (m *MqttPublisher) Publish(topic Topic, payload []byte) error {
	token := m.client.Publish(string(topic), 0, false, payload)
	token.WaitTimeout(10 * time.Millisecond)
	if err := token.Error(); err != nil {
		return fmt.Errorf("unable to events to topic: %v", err)
	}
	return nil
}

func (m *MqttPublisher) PublishAsString(topic Topic, payload fmt.Stringer) error {
	return m.Publish(topic, []byte(payload.String()))
}

// Subscribe register func to execute on message
func (m *MqttPublisher) Subscribe(topic Topic, callback MQTT.MessageHandler) {
	tokenResp := m.client.Subscribe(string(topic), m.qos, callback)
	if tokenResp.Error() != nil {
		zap.S().Fatalf("%+v", tokenResp.Error())
	}
}

// Close connection to broker
func (m *MqttPublisher) Close() error {
	m.client.Disconnect(500)
	return nil
}
