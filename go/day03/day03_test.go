package main

import (
	"adventOfCode/utils"
	"fmt"
	"sort"
	"strings"
	"testing"
)

var testInput = NewSchematic(strings.TrimSpace(`
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`))

func sortSchematicNumbers(schematicNumbers *[]SchematicNumber) {
	sort.SliceStable(*schematicNumbers, func(i, j int) bool {
		return (*schematicNumbers)[i].Value < (*schematicNumbers)[j].Value
	})
}

func TestGetSchematicNumbers(t *testing.T) {
	expect := []SchematicNumber{
		{Value: 35},
		{Value: 633},
		{Value: 617},
		{Value: 58},
		{Value: 592},
		{Value: 755},
		{Value: 664},
		{Value: 598},
		{Value: 467},
		{Value: 114},
	}
	result, _ := testInput.GetSchematicParts()
	sortSchematicNumbers(&result)
	sortSchematicNumbers(&expect)
	utils.AssertInt(t, len(result), len(expect))

	for i := 0; i < len(expect); i++ {
		if expect[i].Value != result[i].Value {
			t.Errorf("expect %+v, got %+v", expect[i], result[i])
		}
	}
}

func TestGetSchematicDimentions(t *testing.T) {
	expectDimensions := SchematicDimensions{Width: 10, Length: 10}
	utils.AssertDeepEqual(t, testInput.Dimensions, expectDimensions)
}

func TestGetSchematicSymbol(t *testing.T) {
	type coordinates struct {
		x int
		y int
	}
	testCases := []utils.TestCase[coordinates, rune]{
		{Case: coordinates{x: 3, y: 8}, Expected: '$'},
		{Case: coordinates{x: 3, y: 1}, Expected: '*'},
		{Case: coordinates{x: 5, y: 5}, Expected: '+'},
		{Case: coordinates{x: 9, y: 9}, Expected: '.'},
		{Case: coordinates{x: 8, y: 5}, Expected: '8'},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%+v", testCase.Case), func(t *testing.T) {
			result, err := testInput.GetSymbol(testCase.Case.x, testCase.Case.y)
			utils.AssertNotError(t, err)
			utils.AssertRune(t, result, testCase.Expected)
		})
	}

	outOfBoundsCases := []coordinates{
		{x: -1, y: -1},
		{x: -1, y: 10},
		{x: 10, y: -1},
		{x: 10, y: 10},
	}
	for _, testCase := range outOfBoundsCases {
		t.Run(fmt.Sprintf("out of bounds: %+v", testCase), func(t *testing.T) {
			_, err := testInput.GetSymbol(testCase.x, testCase.y)
			utils.AssertExpectError(t, err, ErrOutOfBounds)
		})
	}
}

func TestGetSchematicNumbersWithAdjacentSymbols(t *testing.T) {
	expect := []SchematicNumber{
		{
			Value: 35,
			AdjacentSymbols: map[string]bool{
				".": true,
				"*": true,
			},
		},
		{
			Value: 633,
			AdjacentSymbols: map[string]bool{
				".": true,
				"#": true,
			},
		},
		{
			Value: 617,
			AdjacentSymbols: map[string]bool{
				".": true,
				"*": true,
			},
		},
		{
			Value: 58,
			AdjacentSymbols: map[string]bool{
				".": true,
			},
		},
		{
			Value: 592,
			AdjacentSymbols: map[string]bool{
				".": true,
				"+": true,
			},
		},
		{
			Value: 755,
			AdjacentSymbols: map[string]bool{
				".": true,
				"*": true,
			},
		},
		{
			Value: 664,
			AdjacentSymbols: map[string]bool{
				".": true,
				"$": true,
			},
		},
		{
			Value: 598,
			AdjacentSymbols: map[string]bool{
				".": true,
				"*": true,
			},
		},
		{
			Value: 467,
			AdjacentSymbols: map[string]bool{
				".": true,
				"*": true,
			},
		},
		{
			Value: 114,
			AdjacentSymbols: map[string]bool{
				".": true,
			},
		},
	}
	result, _ := testInput.GetSchematicParts()
	sortSchematicNumbers(&result)
	sortSchematicNumbers(&expect)

	for i := 0; i < len(expect); i++ {
		utils.AssertDeepEqual(t, result[i].AdjacentSymbols, expect[i].AdjacentSymbols)
	}
}

func TestIsPartNumber(t *testing.T) {
	testCases := []utils.TestCase[SchematicNumber, bool]{
		{
			Case: SchematicNumber{
				Value: 35,
				AdjacentSymbols: map[string]bool{
					".": true,
					"*": true,
				},
			},
			Expected: true,
		},
		{
			Case: SchematicNumber{
				Value: 633,
				AdjacentSymbols: map[string]bool{
					".": true,
					"#": true,
				},
			},
			Expected: true,
		},
		{
			Case: SchematicNumber{
				Value: 58,
				AdjacentSymbols: map[string]bool{
					".": true,
				},
			},
			Expected: false,
		},
		{
			Case: SchematicNumber{
				Value: 592,
				AdjacentSymbols: map[string]bool{
					".": true,
					"+": true,
				},
			},
			Expected: true,
		},
		{
			Case: SchematicNumber{
				Value: 664,
				AdjacentSymbols: map[string]bool{
					".": true,
					"$": true,
				},
			},
			Expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%+v", testCase.Case), func(t *testing.T) {
			utils.AssertBool(t, testCase.Case.IsPartNumber(), testCase.Expected)
		})
	}
}

func TestPartNumbersSum(t *testing.T) {
	expectedPartNumbersSum := 4361
	numbers, _ := testInput.GetSchematicParts()
	partNumbersSum := sumPartNumbers(numbers)
	utils.AssertInt(t, partNumbersSum, expectedPartNumbersSum)
}

func TestGearCandidates(t *testing.T) {
	expectedSumGearRatios := 467835
	_, gearCandidates := testInput.GetSchematicParts()
	sumGearRatios := gearCandidates.SumAllGearRatios()
	utils.AssertInt(t, sumGearRatios, expectedSumGearRatios)
}
