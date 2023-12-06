package main

import (
	"adventOfCode/utils"
	"fmt"
	"strings"
	"testing"
)

var inputTest = strings.TrimSpace(`
Time:      7  15   30
Distance:  9  40  200
`)

func TestParseInput(t *testing.T) {
	expected := []Race{
		{Time: 7, Distance: 9},
		{Time: 15, Distance: 40},
		{Time: 30, Distance: 200},
	}
	result := parseInput(inputTest)
	utils.AssertDeepEqual(t, result, expected)
}

func TestGetWaysToWin(t *testing.T) {
	races := parseInput(inputTest)
	testCases := []utils.TestCase[Race, int]{
		{Case: races[0], Expected: 4},
		{Case: races[1], Expected: 8},
		{Case: races[2], Expected: 9},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%+v", testCase.Case), func(t *testing.T) {
			result := testCase.Case.GetWaysToWin()
			utils.AssertInt(t, result, testCase.Expected)
		})
	}
}

func TestGetWaysToWinSingleRace(t *testing.T) {
	expect := 71503
	result := parseInput(strings.ReplaceAll(inputTest, " ", ""))[0].GetWaysToWin()
	utils.AssertInt(t, result, expect)
}
