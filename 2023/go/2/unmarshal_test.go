package main

import (
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Attempt
	}{
		{
			name: "1 blue, 2 green, 3 red",
			args: args{"1 blue, 2 green, 3 red"},
			want: Attempt{
				R: 3,
				G: 2,
				B: 1,
			},
		},
		{
			name: "1 blue, 3 red",
			args: args{"1 blue, 3 red"},
			want: Attempt{
				R: 3,
				B: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unmarshal(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
