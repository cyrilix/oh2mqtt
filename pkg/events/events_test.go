package events

import (
	"context"
	"fmt"
	"github.com/cyrilix/mqtt-tools/mqttTooling"
	"github.com/cyrilix/oh2mqtt/pkg/oh2mqtt"
	"sync"
	"testing"
	"time"
)

type MusicControllerMock struct {
	room       string
	sourceName string
	sourceType oh2mqtt.SourceType
	volume     int
	state      oh2mqtt.State
}

func (m *MusicControllerMock) Room() string {
	return m.room
}

func (m *MusicControllerMock) SourceName(_ context.Context) (string, error) {
	return m.sourceName, nil
}

func (m *MusicControllerMock) SourceType(_ context.Context) (oh2mqtt.SourceType, error) {
	return m.sourceType, nil
}

func (m *MusicControllerMock) Volume(_ context.Context) (int, error) {
	return m.volume, nil
}

func (m *MusicControllerMock) SetVolume(_ context.Context, level int) error {
	m.volume = level
	return nil
}

func (m *MusicControllerMock) Mute(_ context.Context) (oh2mqtt.Mute, error) {
	if m.volume == 0 {
		return oh2mqtt.MuteOn, nil
	}
	return oh2mqtt.MuteOff, nil
}

func (m *MusicControllerMock) SetMute(_ context.Context, mute oh2mqtt.Mute) error {
	if mute.Bool() {
		m.volume = 0
	} else {
		m.volume = 100
	}
	return nil
}

func (m *MusicControllerMock) State(_ context.Context) (oh2mqtt.State, error) {
	return m.state, nil
}

type DeviceDiscoverMock struct {
	controllers []oh2mqtt.MusicController
}

func (d *DeviceDiscoverMock) ListMusicController(_ context.Context) ([]oh2mqtt.MusicController, error) {
	return d.controllers, nil
}

type PublisherMock struct {
	muMsgs sync.RWMutex
	msgs   map[mqttTooling.Topic][]string

	wait *sync.WaitGroup
}

func (p *PublisherMock) GetMsgs(topic mqttTooling.Topic) []string {
	p.muMsgs.RLock()
	defer p.muMsgs.RUnlock()
	msgs, ok := p.msgs[topic]
	if !ok {
		return []string{}
	}
	return msgs
}

func (p *PublisherMock) Wait(n int) *sync.WaitGroup {
	p.muMsgs.Lock()
	defer p.muMsgs.Unlock()
	p.wait = &sync.WaitGroup{}
	p.wait.Add(n)
	return p.wait
}

func (p *PublisherMock) Publish(topic mqttTooling.Topic, payload []byte) error {
	return p.publishString(topic, string(payload))
}

func (p *PublisherMock) PublishAsString(topic mqttTooling.Topic, payload fmt.Stringer) error {
	return p.publishString(topic, payload.String())
}

func (p *PublisherMock) publishString(topic mqttTooling.Topic, payload string) error {
	p.muMsgs.Lock()
	defer p.muMsgs.Unlock()

	if p.msgs == nil {
		p.msgs = make(map[mqttTooling.Topic][]string)
	}

	if _, ok := p.msgs[topic]; !ok {
		p.msgs[topic] = []string{}
	}

	p.msgs[topic] = append(p.msgs[topic], payload)
	p.wait.Done()
	return nil
}

func TestOHGateway(t *testing.T) {
	type fields struct {
		publisher *PublisherMock
		opts      []Option
	}
	tests := []struct {
		name   string
		fields fields
		want   map[mqttTooling.Topic]string
	}{
		{
			name: "publish source name",
			fields: fields{
				&PublisherMock{},
				[]Option{
					WithIntervalPublish(5 * time.Millisecond),
					WithTopicFormat("test/room/%s/%s"),
					WithDeviceDiscover(
						&DeviceDiscoverMock{
							controllers: []oh2mqtt.MusicController{
								&MusicControllerMock{
									room:       "room1",
									sourceName: "radio",
								},
							},
						},
					),
				},
			},
			want: map[mqttTooling.Topic]string{mqttTooling.Topic("test/room/room1/source/name"): "radio"},
		},
		{
			name: "publish source type",
			fields: fields{
				&PublisherMock{},
				[]Option{
					WithIntervalPublish(5 * time.Millisecond),
					WithTopicFormat("test/room/%s/%s"),
					WithDeviceDiscover(
						&DeviceDiscoverMock{
							controllers: []oh2mqtt.MusicController{
								&MusicControllerMock{
									room:       "room1",
									sourceName: "radio",
									sourceType: oh2mqtt.SourceTypeRadio,
								},
							},
						},
					),
				},
			},
			want: map[mqttTooling.Topic]string{mqttTooling.Topic("test/room/room1/source/type"): oh2mqtt.SourceTypeRadio.String()},
		},
		{
			name: "publish source state",
			fields: fields{
				&PublisherMock{},
				[]Option{
					WithIntervalPublish(5 * time.Millisecond),
					WithTopicFormat("test/room/%s/%s"),
					WithDeviceDiscover(
						&DeviceDiscoverMock{
							controllers: []oh2mqtt.MusicController{
								&MusicControllerMock{
									room:       "room1",
									sourceName: "radio",
									state:      oh2mqtt.StatePlaying,
								},
							},
						},
					),
				},
			},
			want: map[mqttTooling.Topic]string{mqttTooling.Topic("test/room/room1/state"): oh2mqtt.StatePlaying.String()},
		},
		{
			name: "publish volume level",
			fields: fields{
				&PublisherMock{},
				[]Option{
					WithIntervalPublish(5 * time.Millisecond),
					WithTopicFormat("test/room/%s/%s"),
					WithDeviceDiscover(
						&DeviceDiscoverMock{
							controllers: []oh2mqtt.MusicController{
								&MusicControllerMock{
									room:       "room1",
									sourceName: "radio",
									volume:     10,
								},
							},
						},
					),
				},
			},
			want: map[mqttTooling.Topic]string{mqttTooling.Topic("test/room/room1/volume"): "10"},
		},
		{
			name: "publish mute state",
			fields: fields{
				&PublisherMock{},
				[]Option{
					WithIntervalPublish(5 * time.Millisecond),
					WithTopicFormat("test/room/%s/%s"),
					WithDeviceDiscover(
						&DeviceDiscoverMock{
							controllers: []oh2mqtt.MusicController{
								&MusicControllerMock{
									room:       "room1",
									sourceName: "radio",
									volume:     0,
								},
							},
						},
					),
				},
			},
			want: map[mqttTooling.Topic]string{mqttTooling.Topic("test/room/room1/volume/mute"): "ON"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := NewOHGateway(tt.fields.publisher, tt.fields.opts...)

			// Wait number message between 2 intervals
			w := tt.fields.publisher.Wait(5)
			go o.Start()
			defer o.Close()

			w.Wait()

			for topic, wanted := range tt.want {
				got := tt.fields.publisher.GetMsgs(topic)
				if len(got) == 0 || got[0] != wanted {
					t.Errorf("bad value for topic '%v' = %v, want '%v'", topic, got, wanted)
				}
			}
		})
	}
}
