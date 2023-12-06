package main

import (
	"adventOfCode/utils"
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

// func TestGetWaysToWin(t *testing.T) {
// 	testCase := []utils.TestCase[[2]int, int]{

// 	}
// }
