package two

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
			name: "shall find five numbers and no symbols",
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
		},
		{
			name: "shall find three numbers and one symbol",
			args: args{
				input: `1.2.
.*.
.3.`,
			},
			wantNumbers: map[int]map[int]int{
				0: {0: 1, 2: 2},
				2: {1: 3},
			},
			wantSymbols: map[int]map[int]struct{}{
				1: {1: {}},
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
			name: "shall find two touching numbers",
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
			want: []int{242, 21},
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
			want: 0,
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
			want: 467835,
		},
		{
			name: "edge",
			args: args{
				input: `.1
.*
2.`,
			},
			want: 2,
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
