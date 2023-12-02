package main

import (
	"adventOfCode/utils"
	"testing"
)

func TestFindFirstDigit(t *testing.T) {
	inputs := []utils.TestCase[string, int]{
		{Case: "1abc2", Expected: 1},
		{Case: "pqr3stu8vwx", Expected: 3},
		{Case: "a1b2c3d4e5f", Expected: 1},
		{Case: "treb7uchet", Expected: 7},
	}
	for _, input := range inputs {
		t.Run(input.Case, func(t *testing.T) {
			result, _ := getFirstDigit(input.Case)

			if result != input.Expected {
				t.Errorf("from %+v got %d", input, result)
			}
		})
	}
}

func TestPanicFindFirstDigit(t *testing.T) {
	line := "absajberakb"
	want := NoDigitFound
	if _, err := getFirstDigit(line); err != want {
		t.Errorf("should have panicked with '%v', instead got '%v'", want, err)
	}
}

func TestGetCalibrationValue(t *testing.T) {
	inputs := []utils.TestCase[string, int]{
		{Case: "1abc2", Expected: 12},
		{Case: "pqr3stu8vwx", Expected: 38},
		{Case: "a1b2c3d4e5f", Expected: 15},
		{Case: "treb7uchet", Expected: 77},
	}

	for _, input := range inputs {
		t.Run(input.Case, func(t *testing.T) {
			result, _ := getCalibrationValue(input.Case)
			if result != input.Expected {
				t.Errorf("expected %d got %d", input.Expected, result)
			}
		})
	}
}
