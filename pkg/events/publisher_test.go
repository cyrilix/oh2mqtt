package events

import "testing"

func TestTopicFormatter_Apply(t *testing.T) {
	type args struct {
		room  string
		topic string
	}
	tests := []struct {
		name string
		t    TopicFormatter
		args args
		want Topic
	}{
		{name: "default", t: TopicFormatter("room/%s/topic/%s"), args: args{"bedroom", "state"}, want: "room/bedroom/topic/state"},
		{name: "setter", t: TopicFormatter("room/%s/topic/%s/set"), args: args{"bedroom", "state"}, want: "room/bedroom/topic/state/set"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Apply(tt.args.room, tt.args.topic); got != tt.want {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}
