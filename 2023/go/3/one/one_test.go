package one

import (
	"reflect"
	"testing"
)

func Test_parseSchema(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name        string
		args        args
		wantNumbers map[int]map[int]int
		wantSymbols map[int]map[int]struct{}
	}{
		{
			name: "shall find five numbers and one symbol",
			args: args{
				input: `.11.22.
..7!3..
..101..`,
			},
			wantNumbers: map[int]map[int]int{
				0: {1: 11, 2: 11, 4: 22, 5: 22},
				1: {2: 7, 4: 3},
				2: {2: 101, 3: 101, 4: 101},
			},
			wantSymbols: map[int]map[int]struct{}{
				1: {3: {}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNumbers, gotSymbols := parseSchema(tt.args.input)

			for k, v := range tt.wantNumbers {
				for i, el := range v {
					if got, ok := gotNumbers[k][i]; !ok || got != el {
						t.Errorf("parseSchema() want numbers = %v, got numbers = %v", tt.wantNumbers, gotNumbers)
						return
					}
				}
			}

			for k, v := range gotNumbers {
				for i, el := range v {
					if got, ok := tt.wantNumbers[k][i]; !ok || got != el {
						t.Errorf("parseSchema() want numbers = %v, got numbers = %v", tt.wantNumbers, gotNumbers)
						return
					}
				}
			}

			for k, v := range tt.wantSymbols {
				for i := range v {
					if _, ok := gotSymbols[k][i]; !ok {
						t.Errorf("parseSchema() want symbols = %v, got symbols = %v", tt.wantSymbols, gotSymbols)
						return
					}
				}
			}

			for k, v := range gotSymbols {
				for i := range v {
					if _, ok := tt.wantSymbols[k][i]; !ok {
						t.Errorf("parseSchema() want symbols = %v, got symbols = %v", tt.wantSymbols, gotSymbols)
						return
					}
				}
			}
		})
	}
}

func Test_findTouching(t *testing.T) {
	type args struct {
		symbols map[int]map[int]struct{}
		numbers map[int]map[int]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "shall find five touching numbers",
			args: args{
				symbols: map[int]map[int]struct{}{
					1: {3: {}},
				},
				numbers: map[int]map[int]int{
					0: {1: 11, 2: 11, 4: 22, 5: 22},
					1: {2: 7, 4: 3},
					2: {2: 101, 3: 101, 4: 101},
				},
			},
			want: []int{11, 22, 7, 3, 101},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPortNumbers(tt.args.symbols, tt.args.numbers)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPortNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "trivial",
			args: args{
				input: `.11.22.
..7!3..
..101..`,
			},
			want: 144,
		},
		{
			name: "base examples",
			args: args{
				input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			},
			want: 4361,
		},
		{
			name: "first 3 lines or input",
			args: args{
				input: `...766.......821.547.....577......................................387.....................56..........446.793..........292..................
...........................%...../.....981..........627..../..........-.....623......610..-..............*..................16......891.....
...$...........716..&336.......470.325.................*.84........$..34....*.....+.....#.....*76....#.........303.433........-........&....`,
			},
			want: 5639,
		},
		{
			name: "last 3 lines of input",
			args: args{
				input: `......538.581........&....*............%......10.....168....537&....296..*......177...192................-.......470........................
..................661......496.346*.....870............*................958....-......*......-....@......101.....+..........................
..808..............................365..................195.........................90......482.837............................404.214......`,
			},
			want: 6945,
		},
		{
			name: "edge",
			args: args{
				input: `.1
.*`,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.input); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
