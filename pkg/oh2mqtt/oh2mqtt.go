package oh2mqtt

import (
	"context"
	"fmt"
	"git.cyrilix.bzh/cyrilix/goupnp/dcps/openhome"
	"go.uber.org/zap"
	"strings"
)

const (
	SourceTypeUnknown SourceType = iota
	SourceTypePlaylist
	SourceTypeRadio
)
const (
	StateUnknown State = iota
	StateStopped
	StatePlaying
	StatePaused
	StateBuffering
)

const (
	MuteUnknown Mute = iota
	MuteOn
	MuteOff
)

type Mute int

func muteFromBool(v bool) Mute {
	if v {
		return MuteOn
	}
	return MuteOff
}

func (m Mute) String() string {
	switch m {
	case MuteOn:
		return "ON"
	case MuteOff:
		return "OFF"
	}
	return "unknown"
}

func (m Mute) Bool() bool {
	switch m {
	case MuteOn:
		return true
	case MuteOff:
		return false
	}
	return false
}

type State int

func stateFromString(v string) State {
	switch strings.ToLower(v) {
	case "stopped":
		return StateStopped
	case "playing":
		return StatePlaying
	case "paused":
		return StatePaused
	case "buffering":
		return StateBuffering
	}
	return StateUnknown
}

func (s State) String() string {
	switch s {
	case StatePlaying:
		return "Playing"
	case StateStopped:
		return "Stopped"
	case StatePaused:
		return "Paused"
	case StateBuffering:
		return "Buffering"
	}
	return "unknown"
}

type SourceType int

func (s SourceType) String() string {
	switch s {
	case SourceTypeUnknown:
		return "unknown"
	case SourceTypePlaylist:
		return "Playlist"
	case SourceTypeRadio:
		return "Radio"
	}
	return "unknown"
}

func sourceTypeFromType(srcType string) SourceType {
	switch srcType {
	case "Radio":
		return SourceTypeRadio
	case "Playlist":
		return SourceTypePlaylist
	}
	return SourceTypeUnknown
}

type MusicController interface {
	Room() string
	SourceName(ctx context.Context) (string, error)
	SourceType(ctx context.Context) (SourceType, error)
	Volume(ctx context.Context) (int, error)
	SetVolume(ctx context.Context, level int) error
	Mute(ctx context.Context) (Mute, error)
	SetMute(ctx context.Context, mute Mute) error
	State(ctx context.Context) (State, error)
}

type UpnpMusicController struct {
	roomName string
	svcs     Services
}

func (m *UpnpMusicController) Room() string {
	return m.roomName
}

// SourceName return display name of the current source
func (m *UpnpMusicController) SourceName(ctx context.Context) (string, error) {
	_, name, err := m.sourceInfo(ctx)
	if err != nil {
		return "", fmt.Errorf("unable to find current source infos for '%v': %w", m.roomName, err)
	}
	return name, nil
}

// SourceType return type of the current source
func (m *UpnpMusicController) SourceType(ctx context.Context) (SourceType, error) {
	srcType, _, err := m.sourceInfo(ctx)
	if err != nil {
		return SourceTypeUnknown, fmt.Errorf("unable to find current source type for '%v': %w", m.roomName, err)
	}
	return srcType, nil
}

func (m *UpnpMusicController) sourceInfo(ctx context.Context) (SourceType, string, error) {
	idx, err := m.svcs.product.SourceIndexCtx(ctx)
	if err != nil {
		return SourceTypeUnknown, "", fmt.Errorf("unable to find current source type for '%v': %w", m.roomName, err)
	}
	_, srcType, name, _, err := m.svcs.product.SourceCtx(ctx, idx)
	if err != nil {
		return SourceTypeUnknown, "", fmt.Errorf("unable to process current source type for '%v': %w", m.roomName, err)
	}
	return sourceTypeFromType(srcType), name, nil
}

// Volume return the current Volume setting in percent
func (m *UpnpMusicController) Volume(ctx context.Context) (int, error) {
	v, err := m.svcs.volume.VolumeCtx(ctx)
	if err != nil {
		return -1, fmt.Errorf("unable to read current volume for '%v': %w", m.roomName, err)
	}
	return int(v), err
}

func (m *UpnpMusicController) SetVolume(ctx context.Context, level int) error {
	err := m.svcs.volume.SetVolumeCtx(ctx, uint32(level))
	if err != nil {
		return fmt.Errorf("unable to set volume level for '%v': %w", m.roomName, err)
	}
	return nil
}

func (m *UpnpMusicController) Mute(ctx context.Context) (Mute, error) {
	mute, err := m.svcs.volume.MuteCtx(ctx)
	if err != nil {
		return MuteUnknown, fmt.Errorf("unable to read mute status for '%v': %w", m.roomName, err)
	}
	return muteFromBool(mute), nil
}

func (m *UpnpMusicController) SetMute(ctx context.Context, mute Mute) error {
	err := m.svcs.volume.SetMuteCtx(ctx, mute.Bool())
	if err != nil {
		return fmt.Errorf("unable to mute volume for '%v': %w", m.roomName, err)
	}
	return nil
}

func (m *UpnpMusicController) State(ctx context.Context) (State, error) {
	st, err := m.SourceType(ctx)
	if err != nil {
		return StateUnknown, fmt.Errorf("unable to detect current source for '%v': %w", m.roomName, err)
	}
	var rawState string
	switch st {
	case SourceTypeRadio:
		rawState, err = m.svcs.radio.TransportStateCtx(ctx)
	case SourceTypePlaylist:
		rawState, err = m.svcs.playlist.TransportStateCtx(ctx)
	default:
		return StateUnknown, fmt.Errorf("unknown state '%v' for '%v'", st, m.roomName)
	}
	if err != nil {
		return StateUnknown, fmt.Errorf("unable to read TransportState from '%s' for '%v'", st, m.roomName)
	}
	return stateFromString(rawState), nil
}

type Services struct {
	info     *openhome.Info1
	product  ProductService
	volume   VolumeService
	radio    RadioService
	playlist PlaylistService
}

type DevicesDiscover interface {
	ListMusicController(ctx context.Context) ([]MusicController, error)
}

type UpnpDevicesDiscover struct{}

func (d *UpnpDevicesDiscover) ListMusicController(ctx context.Context) ([]MusicController, error) {
	var mcs []MusicController

	products, _, err := openhome.NewProduct2ClientsCtx(ctx)
	if err != nil {
		return []MusicController{}, fmt.Errorf("unable to list openhome products: %w", err)
	}
	for _, p := range products {
		infos, err := openhome.NewInfo1ClientsFromRootDevice(p.RootDevice, p.Location)
		if err != nil {
			return []MusicController{}, fmt.Errorf("unable to list info service for device %s: %v",
				p.RootDevice.Device.FriendlyName, err)
		}

		volumes, err := openhome.NewVolume1ClientsFromRootDevice(p.RootDevice, p.Location)
		if err != nil {
			return []MusicController{}, fmt.Errorf("unable to list volume service for device %s: %v",
				p.RootDevice.Device.FriendlyName, err)
		}

		radios, err := openhome.NewRadio1ClientsFromRootDevice(p.RootDevice, p.Location)
		if err != nil {
			return []MusicController{}, fmt.Errorf("unable to list radios service for device %s: %v",
				p.RootDevice.Device.FriendlyName, err)
		}

		playlists, err := openhome.NewPlaylist1ClientsFromRootDevice(p.RootDevice, p.Location)
		if err != nil {
			return []MusicController{}, fmt.Errorf("unable to list playlists service for device %s: %v",
				p.RootDevice.Device.FriendlyName, err)
		}

		svc := Services{
			info:     infos[0],
			product:  p,
			volume:   volumes[0],
			radio:    radios[0],
			playlist: playlists[0],
		}
		roomName, _, _, _, _, err := p.Product()
		if err != nil {
			zap.S().Warnf("unable to read Product Room Name, use device FriendlyName in place: %v", err)
			roomName = p.RootDevice.Device.FriendlyName
		}
		mcs = append(mcs, &UpnpMusicController{
			roomName,
			svc,
		})
	}
	return mcs, nil
}
