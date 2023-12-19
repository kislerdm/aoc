package main

import "testing"

func Test_parseNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantShift int
	}{
		{
			name: "one",
			args: args{
				"one1s",
			},
			want:      1,
			wantShift: 3,
		},
		{
			name: "three",
			args: args{
				"three2sa4as",
			},
			want:      3,
			wantShift: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, gotShift := parseNumber(tt.args.s); got != tt.want || gotShift != tt.wantShift {
				t.Errorf("parseNumber() = %v, %v, want %v, %v", got, gotShift, tt.want, tt.wantShift)
			}
		})
	}
}
