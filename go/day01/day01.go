package main

import (
	"adventOfCode/utils"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var NoDigitFound = errors.New("unable to find a digit in provided string")

func main() {
	fmt.Println("Day 1")
	fmt.Println("first part:")
	input := utils.Reader(2023, 01)
	sumCalibrationValues := 0
	stringsList := strings.Split(input, "\n")
	for _, line := range stringsList {
		if line != "" {
			if calibrationValue, err := getCalibrationValue(line); err == nil {
				sumCalibrationValues += calibrationValue
			} else {
				log.Fatalf("error processing %s: %s", line, err)
			}
		}
	}
	fmt.Println(sumCalibrationValues)

	fmt.Println("second part:")
	sumCalibrationValues = 0
	for _, line := range stringsList {
		if line != "" {
			replacedLine := replaceSubstringNumbers(line)
			if calibrationValue, err := getCalibrationValue(replacedLine); err == nil {
				sumCalibrationValues += calibrationValue
			} else {
				log.Fatalf("error processing %s: %s", line, err)
			}
		}
	}
	fmt.Println(sumCalibrationValues)
}

func replaceSubstringNumbers(line string) string {
	runes := []rune(line)
	resultRunes := []rune{}
	for len(runes) > 0 {
		pushRune := true
		replacement, length := getNumberReplacement(runes)
		for length > 0 {
			runes = runes[length-1:]
			resultRunes = append(resultRunes, replacement)
			replacement, length = getNumberReplacement(runes)
			pushRune = false
		}
		if pushRune == true {
			resultRunes = append(resultRunes, runes[0])
		}
		runes = runes[1:]
		if len(runes) < 3 {
			resultRunes = append(resultRunes, runes...)
			runes = []rune{}
		}
	}
	return string(resultRunes)
}

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

func getNumberReplacement(input []rune) (replacement rune, length int) {
	replacement, length = 0, 0
	runes := []rune(input)
	numbers := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	minWindowSize := 3
	maxWindowSize := 5
	for j := minWindowSize; j <= min(maxWindowSize, len(runes)); j++ {
		windowSubstring := string(runes[0:j])
		if number, ok := numbers[windowSubstring]; ok {
			replacement, length = number, j
			break
		}
	}

	return
}
