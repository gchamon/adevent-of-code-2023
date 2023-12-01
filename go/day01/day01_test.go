package main

import "testing"

type input struct {
	Line     string
	Expected int
}

func TestFindFirstDigit(t *testing.T) {
	inputs := []input{
		{Line: "1abc2", Expected: 1},
		{Line: "pqr3stu8vwx", Expected: 3},
		{Line: "a1b2c3d4e5f", Expected: 1},
		{Line: "treb7uchet", Expected: 7},
	}
	for _, input := range inputs {
		result, _ := getFirstDigit(input.Line)

		if result != input.Expected {
			t.Errorf("from %+v got %d", input, result)
		}
	}
}

func TestPanicFindFirstDigit(t *testing.T) {
	line := "absajberakb"
	want := NoDigitFound
	if _, err := getFirstDigit(line); err != want {
		t.Errorf("should have panicked with '%v', instead got '%v'", want, err)
	}
}
