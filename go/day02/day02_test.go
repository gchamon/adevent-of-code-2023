package main

import (
	"adventOfCode/utils"
	"reflect"
	"testing"
)

func TestGameStringParse(t *testing.T) {
	testCases := []utils.TestCase[string, CubeGame]{
		{
			Case: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			Expected: CubeGame{Id: 1, CubesSubsets: []CubesSubset{
				{Red: 4, Blue: 3, Green: 0},
				{Red: 1, Blue: 6, Green: 2},
				{Red: 0, Blue: 0, Green: 2},
			}},
		},
		{
			Case: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			Expected: CubeGame{Id: 2, CubesSubsets: []CubesSubset{
				{Red: 0, Blue: 1, Green: 2},
				{Red: 1, Blue: 4, Green: 3},
				{Red: 0, Blue: 1, Green: 1},
			}},
		},
		{
			Case: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			Expected: CubeGame{Id: 3, CubesSubsets: []CubesSubset{
				{Red: 20, Blue: 6, Green: 8},
				{Red: 4, Blue: 5, Green: 13},
				{Red: 1, Blue: 0, Green: 5},
			}},
		},
		{
			Case: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			Expected: CubeGame{Id: 4, CubesSubsets: []CubesSubset{
				{Red: 3, Blue: 6, Green: 1},
				{Red: 6, Blue: 0, Green: 3},
				{Red: 14, Blue: 15, Green: 3},
			}},
		},
		{
			Case: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			Expected: CubeGame{Id: 5, CubesSubsets: []CubesSubset{
				{Red: 6, Blue: 1, Green: 3},
				{Red: 1, Blue: 2, Green: 2},
			}},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Case, func(t *testing.T) {
			result := parseGame(testCase.Case)
			if !reflect.DeepEqual(result, testCase.Expected) {
				t.Errorf("expected %+v, got %+v", testCase.Expected, result)
			}
		})
	}
}
