package oh2mqtt

import (
	"context"
	"reflect"
	"testing"
)

func TestUpnpDevicesDiscover_ListMusicController(t *testing.T) {
	t.Skip("Run only to debug on network")
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []UpnpMusicController
		wantErr bool
	}{
		{
			name:    "default",
			args:    args{ctx: context.Background()},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &UpnpDevicesDiscover{}
			got, err := d.ListMusicController(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMusicController() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListMusicController() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sourceTypeFromType(t *testing.T) {
	type args struct {
		srcType string
	}
	tests := []struct {
		name string
		args args
		want SourceType
	}{
		{name: "radio", args: args{srcType: "Radio"}, want: SourceTypeRadio},
		{name: "playlist", args: args{srcType: "Playlist"}, want: SourceTypePlaylist},
		{name: "empty", args: args{srcType: ""}, want: SourceTypeUnknown},
		{name: "invalid", args: args{srcType: "invalid"}, want: SourceTypeUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sourceTypeFromType(tt.args.srcType); got != tt.want {
				t.Errorf("sourceTypeFromType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSourceType_String(t *testing.T) {
	tests := []struct {
		name string
		s    SourceType
		want string
	}{
		{name: "radio", s: SourceTypeRadio, want: "Radio"},
		{name: "playlist", s: SourceTypePlaylist, want: "Playlist"},
		{name: "unknown", s: SourceTypeUnknown, want: "unknown"},
		{name: "invalid", s: 100, want: "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stateFromString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want State
	}{
		{name: "stopped", args: args{value: "Stopped"}, want: StateStopped},
		{name: "playing", args: args{value: "Playing"}, want: StatePlaying},
		{name: "paused", args: args{value: "Paused"}, want: StatePaused},
		{name: "buffering", args: args{value: "Buffering"}, want: StateBuffering},
		{name: "STOPPED", args: args{value: "STOPPED"}, want: StateStopped},
		{name: "PLAYING", args: args{value: "PLAYING"}, want: StatePlaying},
		{name: "PAUSED", args: args{value: "PAUSED"}, want: StatePaused},
		{name: "BUFFERING", args: args{value: "BUFFERING"}, want: StateBuffering},
		{name: "empty", args: args{value: ""}, want: StateUnknown},
		{name: "invalid", args: args{value: "invalid"}, want: StateUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stateFromString(tt.args.value); got != tt.want {
				t.Errorf("stateFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestState_String(t *testing.T) {
	tests := []struct {
		name string
		s    State
		want string
	}{
		{name: "paused", s: StatePaused, want: "Paused"},
		{name: "buffering", s: StateBuffering, want: "Buffering"},
		{name: "playing", s: StatePlaying, want: "Playing"},
		{name: "stopped", s: StateStopped, want: "Stopped"},
		{name: "unknown", s: StateUnknown, want: "unknown"},
		{name: "invalid", s: 100, want: "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_muteFromBool(t *testing.T) {
	type args struct {
		value bool
	}
	tests := []struct {
		name string
		args args
		want Mute
	}{
		{name: "on", args: args{value: true}, want: MuteOn},
		{name: "off", args: args{value: false}, want: MuteOff},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := muteFromBool(tt.args.value); got != tt.want {
				t.Errorf("stateFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMute_String(t *testing.T) {
	tests := []struct {
		name string
		m    Mute
		want string
	}{
		{name: "on", m: MuteOn, want: "ON"},
		{name: "buffering", m: MuteOff, want: "OFF"},
		{name: "unknown", m: MuteUnknown, want: "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMute_Bool(t *testing.T) {
	tests := []struct {
		name string
		m    Mute
		want bool
	}{
		{name: "on", m: MuteOn, want: true},
		{name: "buffering", m: MuteOff, want: false},
		{name: "unknown", m: MuteUnknown, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Bool(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMusicController_Room(t *testing.T) {
	type fields struct {
		roomName string
		svcs     Services
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "default",
			fields: fields{
				roomName: "room1",
				svcs:     Services{},
			},
			want: "room1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UpnpMusicController{
				roomName: tt.fields.roomName,
				svcs:     tt.fields.svcs,
			}
			if got := m.Room(); got != tt.want {
				t.Errorf("Room() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMusicController_SourceName(t *testing.T) {
	type fields struct {
		roomName string
		svcs     Services
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{name: "default",
			fields: fields{
				roomName: "room1",
				svcs: Services{
					info: nil,
					product: &productServiceMock{
						SourceTypes:      map[uint32]string{0: "Playlist", 1: "Radio"},
						CurrentSourceIdx: 1,
						Sources:          map[uint32]string{0: "Playlist Name", 1: "Radio Name"},
					},
				},
			},
			want:    "Radio Name",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UpnpMusicController{
				roomName: tt.fields.roomName,
				svcs:     tt.fields.svcs,
			}
			got, err := m.SourceName(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SourceName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SourceName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMusicController_Volume(t *testing.T) {
	type fields struct {
		roomName string
		svcs     Services
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				roomName: "room1",
				svcs:     Services{volume: &volumeServiceMock{maxVolume: 100, currentVolume: 75, mute: false}},
			},
			args: args{
				ctx: context.Background(),
			},
			want:    75,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UpnpMusicController{
				roomName: tt.fields.roomName,
				svcs:     tt.fields.svcs,
			}
			got, err := m.Volume(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Volume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Volume() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMusicController_SetVolume(t *testing.T) {
	vsm := &volumeServiceMock{100, 75, false}
	type fields struct {
		roomName string
		svcs     Services
	}
	type args struct {
		ctx   context.Context
		level int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				roomName: "room1",
				svcs:     Services{volume: vsm},
			},
			args: args{
				ctx:   context.Background(),
				level: 80,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UpnpMusicController{
				roomName: tt.fields.roomName,
				svcs:     tt.fields.svcs,
			}
			if err := m.SetVolume(tt.args.ctx, tt.args.level); (err != nil) != tt.wantErr {
				t.Errorf("SetVolume() error = %v, wantErr %v", err, tt.wantErr)
			}
			if int(vsm.currentVolume) != tt.args.level {
				t.Errorf("SetVolume(): bad value %v, want %v", int(vsm.currentVolume), tt.args.level)
			}
		})
	}
}

func TestMusicController_Mute(t *testing.T) {
	type fields struct {
		roomName string
		svcs     Services
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Mute
		wantErr bool
	}{
		{
			name:   "mute on",
			fields: fields{roomName: "room1", svcs: Services{volume: &volumeServiceMock{maxVolume: 100, currentVolume: 75, mute: true}}},
			args: args{
				ctx: context.Background(),
			},
			want:    MuteOn,
			wantErr: false,
		},
		{
			name:   "mute off",
			fields: fields{roomName: "room1", svcs: Services{volume: &volumeServiceMock{maxVolume: 100, currentVolume: 75, mute: false}}},
			args: args{
				ctx: context.Background(),
			},
			want:    MuteOff,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UpnpMusicController{
				roomName: tt.fields.roomName,
				svcs:     tt.fields.svcs,
			}
			got, err := m.Mute(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Mute() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMusicController_SetMute(t *testing.T) {
	vsm := &volumeServiceMock{100, 75, false}
	type fields struct {
		roomName string
		svcs     Services
	}
	type args struct {
		ctx  context.Context
		mute Mute
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		initStatus bool
		want       bool
		wantErr    bool
	}{
		{
			name: "enable mute",
			fields: fields{
				roomName: "room1",
				svcs:     Services{volume: vsm},
			},
			args: args{
				ctx:  context.Background(),
				mute: MuteOn,
			},
			initStatus: false,
			want:       MuteOn.Bool(),
			wantErr:    false,
		},
		{
			name: "disable mute",
			fields: fields{
				roomName: "room1",
				svcs:     Services{volume: vsm},
			},
			args: args{
				ctx:  context.Background(),
				mute: MuteOff,
			},
			initStatus: true,
			want:       MuteOff.Bool(),
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vsm.mute = tt.initStatus
			m := &UpnpMusicController{
				roomName: tt.fields.roomName,
				svcs:     tt.fields.svcs,
			}
			if err := m.SetMute(tt.args.ctx, tt.args.mute); (err != nil) != tt.wantErr {
				t.Errorf("SetMute() error = %v, wantErr %v", err, tt.wantErr)
			}
			if vsm.mute != tt.want {
				t.Errorf("SetMute(): bad value %v, want %v", vsm.mute, tt.want)
			}
		})
	}
}

func TestMusicController_State(t *testing.T) {
	const sourceRadioIdx uint32 = 0
	const sourcePlaylistIdx uint32 = 1

	type fields struct {
		playlistService PlaylistService
		radioService    RadioService
		sourceIndex     uint32
	}
	tests := []struct {
		name    string
		fields  fields
		want    State
		wantErr bool
	}{
		{
			name: "playlist paused",
			fields: fields{
				sourceIndex:     sourcePlaylistIdx,
				playlistService: &playlistMock{playerMock: playerMock{state: StatePaused}},
				radioService:    &radioMock{},
			},
			want:    StatePaused,
			wantErr: false,
		},
		{
			name: "playlist playing",

			fields: fields{
				sourceIndex:     sourcePlaylistIdx,
				playlistService: &playlistMock{playerMock: playerMock{state: StatePlaying}},
				radioService:    &radioMock{},
			},
			want:    StatePlaying,
			wantErr: false,
		},
		{
			name: "playlist buffering",
			fields: fields{
				sourceIndex:     sourcePlaylistIdx,
				playlistService: &playlistMock{playerMock: playerMock{state: StateBuffering}},
				radioService:    &radioMock{},
			},
			want:    StateBuffering,
			wantErr: false,
		},
		{
			name: "playlist stopped",
			fields: fields{
				sourceIndex:     sourcePlaylistIdx,
				playlistService: &playlistMock{playerMock: playerMock{state: StateStopped}},
				radioService:    &radioMock{},
			},
			want:    StateStopped,
			wantErr: false,
		},
		{
			name: "source unknown",
			fields: fields{
				sourceIndex:     1000,
				playlistService: &playlistMock{},
				radioService:    &radioMock{},
			},
			want:    StateUnknown,
			wantErr: true,
		},
		{
			name: "radio paused",
			fields: fields{
				sourceIndex:     sourceRadioIdx,
				playlistService: &playlistMock{},
				radioService:    &radioMock{playerMock{state: StatePaused}},
			},
			want:    StatePaused,
			wantErr: false,
		},
		{
			name: "radio played",
			fields: fields{
				sourceIndex:     sourceRadioIdx,
				playlistService: &playlistMock{},
				radioService:    &radioMock{playerMock{state: StatePlaying}},
			},
			want:    StatePlaying,
			wantErr: false,
		},
		{
			name: "radio buffering",
			fields: fields{
				sourceIndex:     sourceRadioIdx,
				playlistService: &playlistMock{},
				radioService:    &radioMock{playerMock{state: StateBuffering}},
			},
			want:    StateBuffering,
			wantErr: false,
		},
		{
			name: "radio stopped",
			fields: fields{
				sourceIndex:     sourceRadioIdx,
				playlistService: &playlistMock{},
				radioService:    &radioMock{playerMock{state: StateStopped}},
			},
			want:    StateStopped,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UpnpMusicController{
				roomName: "room1",
				svcs: Services{
					product: &productServiceMock{
						SourceTypes:      map[uint32]string{sourceRadioIdx: "Radio", sourcePlaylistIdx: "Playlist"},
						CurrentSourceIdx: tt.fields.sourceIndex,
						Sources:          map[uint32]string{sourceRadioIdx: "Radio", sourcePlaylistIdx: "Playlist"},
					},
					radio:    tt.fields.radioService,
					playlist: tt.fields.playlistService,
				},
			}
			got, err := m.State(context.TODO())
			if (err != nil) != tt.wantErr {
				t.Errorf("State() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("State() got = %v, want %v", got, tt.want)
			}
		})
	}
}
