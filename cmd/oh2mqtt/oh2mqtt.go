package main

import (
	"context"
	"flag"
	"github.com/cyrilix/mqtt-tools/mqttTooling"
	"github.com/cyrilix/oh2mqtt/pkg/events"
	"github.com/hellofresh/health-go/v5"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	DefaultClientId = "openhome2mqtt"
)

func HandleExit(c io.Closer) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Kill, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signals
		_ = c.Close()
		os.Exit(0)
	}()
}

func main() {

	var intervalPublishSec int
	var topicBaseName string
	var probeListen, probePath string

	parameters := mqttTooling.MqttCliParameters{
		ClientId: DefaultClientId,
	}

	mqttTooling.InitMqttFlagSet(&parameters)

	flag.IntVar(&intervalPublishSec, "interval-time-pub", 10, "time interval inb seconds between 2 msgs")
	flag.StringVar(&topicBaseName, "topic-base-name", "openhome/room/%s/%s", "Topcic name pattern to publish")
	flag.StringVar(&probeListen, "probe-listen", ":8080", "Port and ip to listen for probes")
	flag.StringVar(&probePath, "probe-path", "/status", "HTTP path for probes")

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

	client, err := mqttTooling.Connect(&parameters)
	if err != nil {
		zap.S().Fatalf("unable to connect to mqtt bus: %v", err)
	}
	defer client.Disconnect(50)

	p := mqttTooling.NewMQTTPubSub(client)
	ohg := events.NewOHGateway(p, events.WithIntervalPublish(time.Duration(intervalPublishSec)*time.Second), events.WithTopicFormat(topicBaseName))
	defer func() { _ = ohg.Close() }()

	HandleExit(ohg)

	healthz, _ := health.New(
		health.WithChecks(health.Config{
			Name:      "oh2mqtt",
			Timeout:   5 * time.Second,
			SkipOnErr: false,
			Check: func(ctx context.Context) error {
				return nil
			},
		}),
	)
	http.Handle(probePath, healthz.Handler())

	zap.S().Debug("run status handler")
	go func() {
		log.Fatal(http.ListenAndServe(probeListen, nil))
	}()

	ohg.Start()
}
