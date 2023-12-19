package main

import (
	"strconv"
	"strings"
)

type Attempt struct {
	R, G, B int
}

func Unmarshal(s string) Attempt {
	o := Attempt{}

	for _, colVal := range strings.Split(s, ", ") {
		els := strings.SplitN(colVal, " ", 2)
		if len(els) != 2 {
			panic("cannot unmarshal " + s)
		}

		switch colCode := els[1]; colCode {
		case "red":
			o.R = extractColorVal(els[0])
		case "green":
			o.G = extractColorVal(els[0])
		case "blue":
			o.B = extractColorVal(els[0])
		default:
			panic("unknown color code " + colCode)
		}

	}

	return o
}

func extractColorVal(s string) int {
	o, err := strconv.ParseInt(s, 10, 32)
	if err != nil || o == 0 {
		panic("cannot parse value " + s)
	}
	return int(o)
}
