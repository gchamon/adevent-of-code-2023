package main

import (
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)

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
	result := getSchematicNumbers(testInput)
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
