package cli

import (
	"flag"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"io"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func SetDefaultValueFromEnv(value *string, key string, defaultValue string) {
	if os.Getenv(key) != "" {
		*value = os.Getenv(key)
	} else {
		*value = defaultValue
	}
}
func SetIntDefaultValueFromEnv(value *int, key string, defaultValue int) error {
	var sVal string
	if os.Getenv(key) != "" {
		sVal = os.Getenv(key)
		val, err := strconv.Atoi(sVal)
		if err != nil {
			zap.S().Errorf("unable to convert string to int: %v", err)
			return err
		}
		*value = val
	} else {
		*value = defaultValue
	}
	return nil
}
func SetFloat64DefaultValueFromEnv(value *float64, key string, defaultValue float64) error {
	var sVal string
	if os.Getenv(key) != "" {
		sVal = os.Getenv(key)
		val, err := strconv.ParseFloat(sVal, 64)
		if err != nil {
			zap.S().Errorf("unable to convert string to float: %v", err)
			return err
		}
		*value = val
	} else {
		*value = defaultValue
	}
	return nil
}

func HandleExit(c io.Closer) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Kill, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signals
		_ = c.Close()
		os.Exit(0)
	}()
}

func InitMqttFlagSet(flagSet *flag.FlagSet, defaultClientId string, mqttBroker, username, password, clientId *string, mqttQos *int, mqttRetain *bool) {
	SetDefaultValueFromEnv(clientId, "MQTT_CLIENT_ID", defaultClientId)
	SetDefaultValueFromEnv(mqttBroker, "MQTT_BROKER", "tcp://127.0.0.1:1883")

	flagSet.StringVar(mqttBroker, "mqtt-broker", *mqttBroker, "Broker Uri, use MQTT_BROKER env if arg not set")
	flagSet.StringVar(username, "mqtt-username", os.Getenv("MQTT_USERNAME"), "Broker Username, use MQTT_USERNAME env if arg not set")
	flagSet.StringVar(password, "mqtt-password", os.Getenv("MQTT_PASSWORD"), "Broker Password, MQTT_PASSWORD env if args not set")
	flagSet.StringVar(clientId, "mqtt-client-id", *clientId, "Mqtt client id, use MQTT_CLIENT_ID env if args not set")
	flagSet.IntVar(mqttQos, "mqtt-qos", *mqttQos, "Qos to pusblish message, use MQTT_QOS env if arg not set")
	flagSet.BoolVar(mqttRetain, "mqtt-retain", *mqttRetain, "Retain mqtt message, if not set, true if MQTT_RETAIN env variable is set")
}

func InitMqttFlags(defaultClientId string, mqttBroker, username, password, clientId *string, mqttQos *int, mqttRetain *bool) {
	InitMqttFlagSet(flag.CommandLine, defaultClientId, mqttBroker, username, password, clientId, mqttQos, mqttRetain)
}

func InitIntFlag(key string, defValue int) int {
	var value int
	err := SetIntDefaultValueFromEnv(&value, key, defValue)
	if err != nil {
		zap.S().Panicf("invalid int value: %v", err)
	}
	return value
}

func InitFloat64Flag(key string, defValue float64) float64 {
	var value float64
	err := SetFloat64DefaultValueFromEnv(&value, key, defValue)
	if err != nil {
		zap.S().Panicf("invalid value: %v", err)
	}
	return value
}

func Connect(uri, username, password, clientId string) (MQTT.Client, error) {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker(uri)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	opts.SetAutoReconnect(true)
	opts.SetDefaultPublishHandler(
		//define a function for the default message handler
		func(client MQTT.Client, msg MQTT.Message) {
			zap.S().Infof("TOPIC: %s", msg.Topic())
			zap.S().Infof("MSG: %s", msg.Payload())
		})

	//create and start a client using the above ClientOptions
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("unable to connect to mqtt bus: %v", token.Error())
	}
	return client, nil
}
