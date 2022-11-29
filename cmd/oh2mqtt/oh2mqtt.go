package main

import (
	"flag"
	"github.com/cyrilix/oh2mqtt/pkg/cli"
	"github.com/cyrilix/oh2mqtt/pkg/events"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

const (
	DefaultClientId = "openhome2mqtt"
)

func main() {

	var mqttBroker, username, password, clientId string
	var intervalPublishSec int
	var topicBaseName string

	mqttQos := cli.InitIntFlag("MQTT_QOS", 0)
	_, mqttRetain := os.LookupEnv("MQTT_RETAIN")

	cli.InitMqttFlags(DefaultClientId, &mqttBroker, &username, &password, &clientId, &mqttQos, &mqttRetain)

	flag.IntVar(&intervalPublishSec, "interval-time-pub", 10, "time interval inb seconds between 2 msgs")
	flag.StringVar(&topicBaseName, "topic-base-name", "openhome/room/%s/%s", "Topcic name pattern to publish")

	logLevel := zap.LevelFlag("log", zap.InfoLevel, "log level")
	flag.Parse()

	if len(os.Args) <= 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(*logLevel)
	lgr, err := config.Build()
	if err != nil {
		log.Fatalf("unable to init logger: %v", err)
	}
	defer func() {
		if err := lgr.Sync(); err != nil {
			log.Printf("unable to Sync logger: %v\n", err)
		}
	}()
	zap.ReplaceGlobals(lgr)

	client, err := cli.Connect(mqttBroker, username, password, clientId)
	if err != nil {
		zap.S().Fatalf("unable to connect to mqtt bus: %v", err)
	}
	defer client.Disconnect(50)

	p := events.NewMqttPublisher(client)
	ohg := events.NewOHGateway(p, events.WithIntervalPublish(time.Duration(intervalPublishSec)*time.Second), events.WithTopicFormat(topicBaseName))
	defer func() { _ = ohg.Close() }()

	cli.HandleExit(ohg)

	ohg.Start()
}
