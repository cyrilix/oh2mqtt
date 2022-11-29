package oh2mqtt

import (
	"context"
	"fmt"
)

type productServiceMock struct {
	SourceTypes      map[uint32]string
	CurrentSourceIdx uint32
	Sources          map[uint32]string
}

func (p *productServiceMock) SourceCtx(ctx context.Context, Index uint32) (SystemName string, Type string, Name string, Visible bool, err error) {
	if _, ok := p.Sources[Index]; !ok {
		return "", "", "", false, fmt.Errorf("no source idx for value %v", Index)
	}
	return p.Sources[Index], p.SourceTypes[Index], p.Sources[Index], true, nil
}

func (p *productServiceMock) SourceIndexCtx(ctx context.Context) (Value uint32, err error) {
	return p.CurrentSourceIdx, nil
}

type volumeServiceMock struct {
	maxVolume     uint32
	currentVolume uint32
	mute          bool
}

func (v *volumeServiceMock) VolumeCtx(ctx context.Context) (Value uint32, err error) {
	return v.currentVolume, nil
}

func (v *volumeServiceMock) SetVolumeCtx(ctx context.Context, Value uint32) (err error) {
	v.currentVolume = Value
	return nil
}

func (v *volumeServiceMock) MuteCtx(ctx context.Context) (Value bool, err error) {
	return v.mute, nil
}

func (v *volumeServiceMock) SetMuteCtx(ctx context.Context, Value bool) (err error) {
	v.mute = Value
	return nil
}
func (v *volumeServiceMock) CharacteristicsCtx(ctx context.Context) (VolumeMax uint32, VolumeUnity uint32, VolumeSteps uint32,
	VolumeMilliDbPerStep uint32, BalanceMax uint32, FadeMax uint32, err error) {
	return v.maxVolume, 100, 100, 500, 0, 0, nil
}

type playerMock struct {
	state State
	id    uint32
}

func (p *playerMock) IdCtx(ctx context.Context) (Value uint32, err error) {
	return p.id, nil
}

func (p *playerMock) IdArrayCtx(ctx context.Context) (Token uint32, Array []byte, err error) {
	return 0, []byte{}, fmt.Errorf("not implemented")
}

func (p *playerMock) ProtocolInfoCtx(ctx context.Context) (Value string, err error) {
	return "", fmt.Errorf("not implemented")
}

func (p *playerMock) TransportStateCtx(ctx context.Context) (Value string, err error) {
	return p.state.String(), nil
}

type radioMock struct {
	playerMock
}

type playlistMock struct {
	playerMock
}

func (p *playlistMock) PlayCtx(ctx context.Context) (err error) {
	p.state = StatePlaying
	return nil
}
