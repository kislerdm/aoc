package one

import (
	"strings"
)

func Run(input string) int {
	// find coordinates of the numbers and symbols
	numbers, symbols := parseSchema(input)

	found := findPortNumbers(symbols, numbers)

	var o int
	for _, n := range found {
		o += n
	}
	return o
}

func parseSchema(input string) (map[int]map[int]int, map[int]map[int]struct{}) {
	lines := strings.Split(input, "\n")

	// defines the matrix verticalCoordinate -> {horizontalCoordinate -> Value}
	numbers := map[int]map[int]int{}

	// defines the matrix indicating symbols: verticalCoordinate -> {horizontalCoordinate -> Value}
	symbols := map[int]map[int]struct{}{}

	for verticalCoordinate, line := range lines {
		var numbTmp int
		var horizontalCoordinateTmp []int

		for horizontalCoordinate, val := range line {
			if v, ok := readNumber(val); ok {
				if _, ok := numbers[verticalCoordinate]; !ok {
					numbers[verticalCoordinate] = map[int]int{}
				}

				horizontalCoordinateTmp = append(horizontalCoordinateTmp, horizontalCoordinate)
				extNumber(&numbTmp, v)
				for _, c := range horizontalCoordinateTmp {
					numbers[verticalCoordinate][c] = numbTmp
				}

				continue
			}

			numbTmp = 0
			horizontalCoordinateTmp = nil

			if val == '.' {
				continue
			}

			if _, ok := symbols[verticalCoordinate]; !ok {
				symbols[verticalCoordinate] = map[int]struct{}{}
			}
			symbols[verticalCoordinate][horizontalCoordinate] = struct{}{}
		}
	}
	return numbers, symbols
}

func extNumber(tmp *int, v int) {
	if *tmp > 0 {
		*tmp = 10 * *tmp
	}
	*tmp += v
}

// finds all numbers adjacent to all the symbols
// for every given symbol, the numbers map will be checked in the following order:
// one. directly above the symbol from left to right;
// two. in the line of the symbol location from left to right;
// 3. directly below the symbol from left to right.
func findPortNumbers(symbols map[int]map[int]struct{}, numbers map[int]map[int]int) []int {
	var touching []int

	for verticalCoordinate := range symbols {
		for horizontalCoordinate := range symbols[verticalCoordinate] {
			// temp touching numbers "set" to avoid double counting in case several digits of a number touch the symbol
			// example:
			// GIVEN
			// 101
			// .#.
			// WHEN check the line above the symbol '#'
			// THEN the number 101 is expected only
			tmp := map[int]struct{}{}

			// check the line above the symbol
			// diagonal: on the left from the symbol
			if v, ok := exists(numbers, verticalCoordinate-1, horizontalCoordinate-1); ok {
				tmp[v] = struct{}{}
			}
			// directly above
			if v, ok := exists(numbers, verticalCoordinate-1, horizontalCoordinate); ok {
				tmp[v] = struct{}{}
			}
			// diagonal: on the right from the symbol
			if v, ok := exists(numbers, verticalCoordinate-1, horizontalCoordinate+1); ok {
				tmp[v] = struct{}{}
			}

			for v := range tmp {
				touching = append(touching, v)
				delete(tmp, v)
			}

			// check in the line of the symbol location
			// check on the left from the symbol
			if v, ok := exists(numbers, verticalCoordinate, horizontalCoordinate-1); ok {
				touching = append(touching, v)
			}

			// check on the right from the symbol
			if v, ok := exists(numbers, verticalCoordinate, horizontalCoordinate+1); ok {
				touching = append(touching, v)
			}

			// check the line below the symbol
			// diagonal: on the left from the symbol
			if v, ok := exists(numbers, verticalCoordinate+1, horizontalCoordinate-1); ok {
				tmp[v] = struct{}{}
			}
			// directly below
			if v, ok := exists(numbers, verticalCoordinate+1, horizontalCoordinate); ok {
				tmp[v] = struct{}{}
			}
			// diagonal: on the right from the symbol
			if v, ok := exists(numbers, verticalCoordinate+1, horizontalCoordinate+1); ok {
				tmp[v] = struct{}{}
			}
			for v := range tmp {
				touching = append(touching, v)
			}
		}
	}

	return touching
}

func exists(numbers map[int]map[int]int, probeVerticalCoordinate, probeHorizontalCoordinate int) (int, bool) {
	v, ok := numbers[probeVerticalCoordinate]
	if !ok {
		return 0, false
	}
	o, ok := v[probeHorizontalCoordinate]
	return o, ok
}

func readNumber(v rune) (int, bool) {
	switch v {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return int(v - '0'), true
	}
	return 0, false
}
