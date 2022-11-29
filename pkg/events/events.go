package events

import (
	"context"
	"fmt"
	"github.com/cyrilix/oh2mqtt/pkg/oh2mqtt"
	"go.uber.org/zap"
	"time"
)

type Option func(o *OHGateway)

func WithTopicFormat(topicFormat string) Option {
	tf := TopicFormatter(topicFormat)
	return func(o *OHGateway) {
		o.tf = &tf
	}
}

func WithIntervalPublish(interval time.Duration) Option {
	return func(o *OHGateway) {
		o.intervalPublish = interval
	}
}

func WithDeviceDiscover(d oh2mqtt.DevicesDiscover) Option {
	return func(o *OHGateway) {
		o.d = d
	}
}

func NewOHGateway(publisher Publisher, opts ...Option) *OHGateway {
	defaultTopicFormatter := TopicFormatter("room/%s/%s")
	o := &OHGateway{
		tf:              &defaultTopicFormatter,
		d:               &oh2mqtt.UpnpDevicesDiscover{},
		p:               publisher,
		intervalPublish: 30 * time.Second,
		cancel:          make(chan interface{}),
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type OHGateway struct {
	tf              *TopicFormatter
	d               oh2mqtt.DevicesDiscover
	p               Publisher
	intervalPublish time.Duration
	cancel          chan interface{}
}

func (o *OHGateway) Close() error {
	if o.cancel != nil {
		o.cancel <- struct{}{}
	}
	return nil
}

func (o *OHGateway) Start() {
	ctx := context.Background()
	ticker := time.NewTicker(o.intervalPublish)

	for {
		select {
		case <-o.cancel:
			zap.S().Infof("stop Openhome gateway run")
			return
		case <-ticker.C:
		}
		mcs, err := o.d.ListMusicController(ctx)
		if err != nil {
			zap.S().Errorf("unable to list devices: %v", err)
			continue
		}

		for _, mc := range mcs {
			mcTmp := mc
			go func() {
				o.publishState(mcTmp, ctx)
				o.publishSourceType(mcTmp, ctx)
				o.publishSourceName(mcTmp, ctx)
				o.publishVolumeLevel(mcTmp, ctx)
				o.publishMuteState(mcTmp, ctx)
			}()
		}
	}
}

func (o *OHGateway) publishState(mc oh2mqtt.MusicController, ctx context.Context) {
	state, err := mc.State(ctx)
	if err != nil {
		zap.S().Errorf("unable read state for device '%v': %v", mc.Room(), err)
	}
	err = o.p.PublishAsString(o.tf.Apply(mc.Room(), "state"), &state)
	if err != nil {
		zap.S().Errorf("unable to publish state '%v' for device '%s': %v", state, mc.Room(), err)
	}
}

func (o *OHGateway) publishSourceType(mc oh2mqtt.MusicController, ctx context.Context) {
	src, err := mc.SourceType(ctx)
	if err != nil {
		zap.S().Errorf("unable read state for device '%v': %v", mc.Room(), err)
	}
	err = o.p.PublishAsString(o.tf.Apply(mc.Room(), "source/type"), &src)
	if err != nil {
		zap.S().Errorf("unable to publish type source '%v' for device '%s': %v", src, mc.Room(), err)
	}
}

func (o *OHGateway) publishSourceName(mc oh2mqtt.MusicController, ctx context.Context) {
	srcName, err := mc.SourceName(ctx)
	if err != nil {
		zap.S().Errorf("unable read source name for device '%v': %v", mc.Room(), err)
		return
	}
	err = o.p.Publish(o.tf.Apply(mc.Room(), "source/name"), []byte(srcName))
	if err != nil {
		zap.S().Errorf("unable to publish source name '%v' for device '%s': %v", srcName, mc.Room(), err)
	}
}

func (o *OHGateway) publishVolumeLevel(mc oh2mqtt.MusicController, ctx context.Context) {
	volume, err := mc.Volume(ctx)
	if err != nil {
		zap.S().Errorf("unable read volume level for device '%v': %v", mc.Room(), err)
		return
	}
	err = o.p.Publish(o.tf.Apply(mc.Room(), "volume"), []byte(fmt.Sprintf("%v", volume)))
	if err != nil {
		zap.S().Errorf("unable to publish volume level '%v' for device '%s': %v", volume, mc.Room(), err)
	}
}

func (o *OHGateway) publishMuteState(mc oh2mqtt.MusicController, ctx context.Context) {
	mute, err := mc.Mute(ctx)
	if err != nil {
		zap.S().Errorf("unable read mute state for device '%v': %v", mc.Room(), err)
		return
	}

	err = o.p.Publish(o.tf.Apply(mc.Room(), "volume/mute"), []byte(fmt.Sprintf("%v", mute.String())))
	if err != nil {
		zap.S().Errorf("unable to publish volume level '%v' for device '%s': %v", mute.String(), mc.Room(), err)
	}
}
