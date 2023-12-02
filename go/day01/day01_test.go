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

func TestReplaceSubstringNumbers(t *testing.T) {
	testCases := []utils.TestCase[string, string]{
		{Case: "two1nine", Expected: "219"},
		{Case: "eightwothree", Expected: "823"},
		{Case: "abcone2threexyz", Expected: "abc123xyz"},
		{Case: "xtwone3four", Expected: "x2134"},
		{Case: "4nineeightseven2", Expected: "49872"},
		{Case: "zoneight234", Expected: "z18234"},
		{Case: "7pqrstsixteen", Expected: "7pqrst6teen"},
		{Case: "8one2three", Expected: "8123"},
		{Case: "8one2thress", Expected: "812thress"},
		{Case: "8one2thr", Expected: "812thr"},
		{Case: "on8one2th", Expected: "on812th"},
		{Case: "8one2t", Expected: "812t"},
		{Case: "8one2o", Expected: "812o"},
		{Case: "8one2on", Expected: "812on"},
		{Case: "8one2one", Expected: "8121"},
		{Case: "87gfmklvnzfz4594qlksr5", Expected: "87gfmklvnzfz4594qlksr5"},
		{Case: "41seven", Expected: "417"},
		{Case: "bcplhrpvmkxkjxlzl7zkbpdnoneninethreej", Expected: "bcplhrpvmkxkjxlzl7zkbpdn193j"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Case, func(t *testing.T) {
			result := replaceSubstringNumbers(testCase.Case)
			if result != testCase.Expected {
				t.Errorf("expected %s got %s", testCase.Expected, result)
			}
		})
	}
}

func TestCalibrationWithNumbersReplacement(t *testing.T) {
	testCases := []utils.TestCase[string, int]{
		{Case: "two1nine", Expected: 29},
		{Case: "eightwothree", Expected: 83},
		{Case: "abcone2threexyz", Expected: 13},
		{Case: "xtwone3four", Expected: 24},
		{Case: "4nineeightseven2", Expected: 42},
		{Case: "zoneight234", Expected: 14},
		{Case: "7pqrstsixteen", Expected: 76},
		{Case: "fzrpfhbfvj6dbxbtfs7twofksfbshrzkdeightwoqg", Expected: 62},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Case, func(t *testing.T) {
			caseReplaced := replaceSubstringNumbers(testCase.Case)
			if result, err := getCalibrationValue(caseReplaced); err != nil {
				t.Errorf("got error %s", err)
			} else if result != testCase.Expected {
				t.Errorf("got %d expected %d", result, testCase.Expected)
			}
		})
	}
}

type ReplacementExpected struct {
	Replacement rune
	Length      int
}

func TestGetNumberReplacement(t *testing.T) {
	testCases := []utils.TestCase[string, ReplacementExpected]{
		{Case: "one", Expected: ReplacementExpected{Replacement: '1', Length: 3}},
		{Case: "onexyz", Expected: ReplacementExpected{Replacement: '1', Length: 3}},
		{Case: "twoxyz", Expected: ReplacementExpected{Replacement: '2', Length: 3}},
		{Case: "threexyz", Expected: ReplacementExpected{Replacement: '3', Length: 5}},
		{Case: "three", Expected: ReplacementExpected{Replacement: '3', Length: 5}},
		{Case: "fourxyz", Expected: ReplacementExpected{Replacement: '4', Length: 4}},
		{Case: "fivexyz", Expected: ReplacementExpected{Replacement: '5', Length: 4}},
		{Case: "sixxyz", Expected: ReplacementExpected{Replacement: '6', Length: 3}},
		{Case: "sevenxyz", Expected: ReplacementExpected{Replacement: '7', Length: 5}},
		{Case: "eightxyz", Expected: ReplacementExpected{Replacement: '8', Length: 5}},
		{Case: "ninexyz", Expected: ReplacementExpected{Replacement: '9', Length: 4}},
		{Case: "nine", Expected: ReplacementExpected{Replacement: '9', Length: 4}},
		{Case: "ninea", Expected: ReplacementExpected{Replacement: '9', Length: 4}},
		{Case: "nineab", Expected: ReplacementExpected{Replacement: '9', Length: 4}},
		{Case: "nineabc", Expected: ReplacementExpected{Replacement: '9', Length: 4}},
		{Case: "nineabcd", Expected: ReplacementExpected{Replacement: '9', Length: 4}},
		{Case: "nineabcde", Expected: ReplacementExpected{Replacement: '9', Length: 4}},
		{Case: "nineabcdef", Expected: ReplacementExpected{Replacement: '9', Length: 4}},
	}

	for _, testCase := range testCases {
		expectedLength := testCase.Expected.Length
		expectedReplacement := testCase.Expected.Replacement
		t.Run(testCase.Case, func(t *testing.T) {
			if replacement, length := getNumberReplacement(testCase.Case); length == 0 {
				t.Errorf("expected length %d got %d", expectedLength, length)
			} else if replacement != expectedReplacement {
				t.Errorf("expected %v got %v", expectedReplacement, replacement)
			}
		})
	}
}
