package main

import (
	"adventOfCode/utils"
	"errors"
	"strconv"
	"unicode"
)

var NoDigitFound = errors.New("unable to find a digit in provided string")

func getCalibrationValue(line string) (int, error) {
	firstDigit, err := getFirstDigit(line)
	if err != nil {
		return 0, err
	}
	lastDigit, err := getFirstDigit(utils.Reverse(line))
	if err != nil {
		return 0, err
	}
	return firstDigit*10 + lastDigit, nil
}

func getFirstDigit(line string) (int, error) {
	for _, c := range line {
		if unicode.IsDigit(c) {
			digit, _ := strconv.Atoi(string(c))
			return digit, nil
		}
	}
	return 0, NoDigitFound
}
