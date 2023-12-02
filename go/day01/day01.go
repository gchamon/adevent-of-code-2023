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
	calibrationValues := []int{}
	stringsList := strings.Split(input, "\n")
	for _, line := range stringsList {
		if line != "" {
			if calibrationValue, err := getCalibrationValue(line); err == nil {
				calibrationValues = append(calibrationValues, calibrationValue)
			} else {
				log.Fatalf("error processing %s: %s", line, err)
			}
		}
	}
	fmt.Println(sumListOfInt(calibrationValues))

	fmt.Println("second part:")
	calibrationValues = []int{}
	fmt.Println(calibrationValues)
	for _, line := range stringsList {
		if line != "" {
			replacedLine := replaceSubstringNumbers(line)
			if calibrationValue, err := getCalibrationValue(replacedLine); err == nil {
				fmt.Printf("%s: %s %d\n", line, replacedLine, calibrationValue)
				calibrationValues = append(calibrationValues, calibrationValue)
			} else {
				log.Fatalf("error processing %s: %s", line, err)
			}
		}
	}
	fmt.Println(sumListOfInt(calibrationValues))
}

func replaceSubstringNumbers(line string) string {
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
	runes := []rune(line)
	resultRunes := []rune{}
	minWindowSize := 3
	maxWindowSize := 5
	for len(runes) > 0 {
		pushRune := true
		for j := minWindowSize; j <= int(min(maxWindowSize, len(runes))); j++ {
			windowSubstring := string(runes[0:j])
			if number, ok := numbers[windowSubstring]; ok {
				resultRunes = append(resultRunes, number)
				runes = runes[j:]
				pushRune = false
				break
			}
		}
		if pushRune == true {
			resultRunes = append(resultRunes, runes[0])
			runes = runes[1:]
		}
		if len(runes) < minWindowSize {
			resultRunes = append(resultRunes, runes...)
			runes = []rune{}
		}
	}
	return string(resultRunes)
}

func sumListOfInt(array []int) int {
	sum := 0
	for _, n := range array {
		sum += n
	}

	return sum
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
