/*
https://adventofcode.com/2023/day/1

Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?

Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

What is the sum of all of the calibration values?

*/

package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")

	var cnt int64
	for _, line := range lines {
		var l, r int
		var i int
		for i < len(line) {
			if v, shift := parseNumber(line[i:]); v > 0 {
				r = v
				i += shift
			} else {
				i++
			}

			if l == 0 {
				l = r
			}
		}

		if r == 0 {
			r = l
		}

		lineCnt := l*10 + r

		cnt += int64(lineCnt)
	}

	log.Println(cnt)
}

func parseNumber(s string) (val int, shift int) {
	switch s[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return int(s[0] - '0'), 1

	case 'o', 't', 's', 'e':
		if len(s) < 3 {
			return 0, 1
		}

		shift = 3
		switch s[:shift] {
		case "one":
			return 1, shift
		case "two":
			return 2, shift
		case "six":
			return 6, shift
		}

		if len(s) < 5 {
			return 0, 1
		}

		shift = 5
		switch s[:shift] {
		case "three":
			return 3, shift
		case "seven":
			return 7, shift
		case "eight":
			return 8, shift
		}

	case 'f', 'n':
		shift = 4
		if len(s) < shift {
			return 0, 1
		}

		switch s[:4] {
		case "four":
			return 4, shift
		case "five":
			return 5, shift
		case "nine":
			return 9, shift
		}
	}

	return 0, 1
}
