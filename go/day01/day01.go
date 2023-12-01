package main

import (
	"errors"
	"strconv"
	"unicode"
)

var NoDigitFound = errors.New("unable to find a digit in provided string")

func getFirstDigit(line string) (int, error) {
	for _, c := range line {
		if unicode.IsDigit(c) {
			digit, _ := strconv.Atoi(string(c))
			return digit, nil
		}
	}
	return 0, NoDigitFound
}
