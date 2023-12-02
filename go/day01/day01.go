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
		if calibrationValue, err := getCalibrationValue(line); err == nil {
			sumCalibrationValues += calibrationValue
		} else {
			log.Fatalf("error processing %s: %s", line, err)
		}
	}
	fmt.Println(sumCalibrationValues)

	fmt.Println("second part:")
	sumCalibrationValues = 0
	for _, line := range stringsList {
		replacedLine := replaceSubstringNumbers(line)
		if calibrationValue, err := getCalibrationValue(replacedLine); err == nil {
			sumCalibrationValues += calibrationValue
		} else {
			log.Fatalf("error processing %s: %s", line, err)
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
		for length > 0 { // while there are replacements...
			runes = runes[length-1:]                          // pop the first `length`-most chars but the last
			resultRunes = append(resultRunes, replacement)    // and add the replacement to the resulting stack
			replacement, length = getNumberReplacement(runes) // and try again
			pushRune = false                                  // in case there are replacements, the first char in the window won't be added to the resulting stack
		}
		if pushRune == true { // in case no replacement is found add the current rune to the resulting stack
			resultRunes = append(resultRunes, runes[0])
		}
		runes = runes[1:]   // and finally pop the first item from the search stack
		if len(runes) < 3 { // in case there are only 2 or less runes, there can't be any more replacements, and we add the rest to the resulting stack
			resultRunes = append(resultRunes, runes...)
			runes = []rune{} // clear the search stack so that the for exists. Could be done with a break and an empty for, but I think this is more readable
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
