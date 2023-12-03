package main

import (
	"adventOfCode/utils"
	"errors"
	"fmt"
	"reflect"
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
	result := testInput.GetSchematicNumbers()
	sortSchematicNumbers(&result)
	sortSchematicNumbers(&expect)
	if len(expect) != len(result) {
		t.Errorf("expected %d elements, got %d elements", len(expect), len(result))
	}

	for i := 0; i < len(expect); i++ {
		if expect[i].Value != result[i].Value {
			t.Errorf("expect %+v, got %+v", expect[i], result[i])
		}
	}
}

func TestGetSchematicDimentions(t *testing.T) {
	expectDimensions := SchematicDimensions{Width: 10, Length: 10}
	if !reflect.DeepEqual(expectDimensions, testInput.Dimensions) {
		t.Errorf("expect %+v, got %+v", expectDimensions, testInput.Dimensions)
	}
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
			if err != nil {
				t.Errorf("wasn't expecting error %s", err)
			}
			if result != testCase.Expected {
				t.Errorf("expected %v, got %v", testCase.Expected, result)
			}
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
			result, err := testInput.GetSymbol(testCase.x, testCase.y)
			if err == nil {
				t.Errorf("was expecting error %s", err)
			}
			if !errors.Is(err, ErrOutOfBounds) {
				t.Errorf("expected %v, got %v", ErrOutOfBounds, result)
			}
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
	result := testInput.GetSchematicNumbers()
	sortSchematicNumbers(&result)
	sortSchematicNumbers(&expect)

	for i := 0; i < len(expect); i++ {
		if !reflect.DeepEqual(expect[i].AdjacentSymbols, result[i].AdjacentSymbols) {
			t.Errorf("expect %+v, got %+v", expect[i], result[i])
		}
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
			result := testCase.Case.IsPartNumber()
			if result != testCase.Expected {
				t.Errorf("expect part number %v, got %v", testCase.Expected, result)
			}
		})
	}
}
