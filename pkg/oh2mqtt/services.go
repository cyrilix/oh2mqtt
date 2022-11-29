package oh2mqtt

import "context"

type ProductService interface {
	SourceCtx(ctx context.Context, Index uint32) (SystemName string, Type string, Name string, Visible bool, err error)
	SourceIndexCtx(ctx context.Context) (Value uint32, err error)
}

type VolumeService interface {
	// VolumeCtx The Volume action reports the current Volume setting.
	VolumeCtx(ctx context.Context) (Value uint32, err error)
	SetVolumeCtx(ctx context.Context, Value uint32) (err error)
	MuteCtx(ctx context.Context) (Value bool, err error)
	SetMuteCtx(ctx context.Context, Value bool) (err error)
	CharacteristicsCtx(ctx context.Context) (VolumeMax uint32, VolumeUnity uint32, VolumeSteps uint32,
		VolumeMilliDbPerStep uint32, BalanceMax uint32, FadeMax uint32, err error)
}

type PlayerService interface {
	// IdCtx report The id of the currently selected preset (as set by the last call to SetId).
	// Or 0 if no preset has ever been selected or SetChannel was called more recently than SetId.
	IdCtx(ctx context.Context) (Value uint32, err error)

	// IdArrayCtx is Base64-encoded array of 32-bit big endian unsigned integers. Will always be (ChannelsMax * 4) bytes long.
	// Each id represents a preset id. A value of 0 implies that no preset is set for that slot.
	IdArrayCtx(ctx context.Context) (Token uint32, Array []byte, err error)

	// ProtocolInfoCtx return Audio formats and networking protocols supported by the device.
	ProtocolInfoCtx(ctx context.Context) (Value string, err error)

	// TransportStateCtx return One of Stopped, Playing, Paused or Buffering.
	TransportStateCtx(ctx context.Context) (Value string, err error)
}

type RadioService interface {
	PlayerService
}

type PlaylistService interface {
	PlayerService

	// PlayCtx If TransportState is Stopped or Paused, start playing the track indicated by the Id state variable.
	// If TransportState is Playing or Buffering, re-start playing the current track from the beginning.
	// If the Playlist source is not active, deactivates the current source and activates Playlist.
	PlayCtx(ctx context.Context) (err error)
}
