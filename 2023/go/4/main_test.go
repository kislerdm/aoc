package main

import (
	"reflect"
	"testing"
)

func Test_readNumbers(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[int]struct{}
	}{
		{
			name: "",
			args: args{
				s: " 9 39 27",
			},
			want: map[int]struct{}{9: {}, 39: {}, 27: {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readNumbers(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
