package events

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type Topic string

type TopicFormatter string

func (t TopicFormatter) Apply(room string, topic string) Topic {
	return Topic(fmt.Sprintf(string(t), room, topic))
}

type Publisher interface {
	Publish(topic Topic, payload []byte) error
	PublishAsString(topic Topic, payload fmt.Stringer) error
}

func NewMqttPublisher(client mqtt.Client) *MqttPublisher {
	return &MqttPublisher{client: client}
}

type MqttPublisher struct {
	client mqtt.Client
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
