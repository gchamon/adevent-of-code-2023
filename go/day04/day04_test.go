package main

import (
	"adventOfCode/utils"
	"strings"
	"testing"
)

var inputTest = strings.Split(strings.TrimSpace(`
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`), "\n")

func TestCardParseID(t *testing.T) {
	type TestCaseID utils.TestCase[string, int]
	testCases := []TestCaseID{}
	for i, input := range inputTest {
		testCases = append(testCases, TestCaseID{Case: input, Expected: i + 1})
	}
	for _, testCase := range testCases {
		t.Run(testCase.Case, func(t *testing.T) {
			result := NewScratchCard(testCase.Case)
			utils.AssertInt(t, result.ID, testCase.Expected)
		})
	}
}
