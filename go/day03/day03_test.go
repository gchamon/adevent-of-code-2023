package main

import (
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

func TestGetSchematicNumbersWithAdjacentSymbols(t *testing.T) {
	expect := []SchematicNumber{
		{
			Value: 35,
			AdjacentSymbols: map[rune]bool{
				'.': true,
				'*': true,
			},
		},
		{
			Value: 633,
			AdjacentSymbols: map[rune]bool{
				'.': true,
				'#': true,
			},
		},
		{
			Value: 617,
			AdjacentSymbols: map[rune]bool{
				'.': true,
				'*': true,
			},
		},
		{
			Value: 58,
			AdjacentSymbols: map[rune]bool{
				'.': true,
			},
		},
		{
			Value: 592,
			AdjacentSymbols: map[rune]bool{
				'.': true,
				'+': true,
			},
		},
		{
			Value: 755,
			AdjacentSymbols: map[rune]bool{
				'.': true,
				'*': true,
			},
		},
		{
			Value: 664,
			AdjacentSymbols: map[rune]bool{
				'.': true,
				'$': true,
			},
		},
		{
			Value: 598,
			AdjacentSymbols: map[rune]bool{
				'.': true,
				'*': true,
			},
		},
		{
			Value: 467,
			AdjacentSymbols: map[rune]bool{
				'.': true,
				'*': true,
			},
		},
		{
			Value: 114,
			AdjacentSymbols: map[rune]bool{
				'.': true,
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
