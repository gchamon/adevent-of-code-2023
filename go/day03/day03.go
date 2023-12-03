package main

import (
	"errors"
	"sort"
	"strconv"
)

type SchematicNumber struct {
	Value           int
	AdjacentSymbols map[rune]bool
}

type SchematicDimensions struct {
	Width  int
	Length int
}

type Schematic struct {
	Contents   string
	Dimensions SchematicDimensions
}

func NewSchematic(input string) (schematic Schematic) {
	schematic.Contents = input
	for i, c := range input {
		if c == rune('\n') {
			schematic.Dimensions.Width = i
			schematic.Dimensions.Length = len(input) / schematic.Dimensions.Width
			break
		}
	}
	return
}

var ErrOutOfBounds = errors.New("symbol out of bounds")

// Returns the symbol referenced by the coordinates x, y.
// Coordinates are 0-based indexes
// Consider the following schematic:
// 467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..
// y will start mapping from top to bottom, and x will start mapping
// from left to right. So the pair (6, 3) references '#'
func (s Schematic) GetSymbol(x, y int) (r rune, e error) {
	if x >= s.Dimensions.Width || y >= s.Dimensions.Length || x < 0 || y < 0 {
		r, e = 0, ErrOutOfBounds
	} else {
		// each line accounts for Width amount of characters +1 (the new line character)
		symbolIndex := x + s.Dimensions.Width*y + y
		r, e = rune(s.Contents[symbolIndex]), nil
	}
	return
}

func sortSchematicNumbers(schematicNumbers *[]SchematicNumber) {
	sort.SliceStable(*schematicNumbers, func(i, j int) bool {
		return (*schematicNumbers)[i].Value < (*schematicNumbers)[j].Value
	})
}

func (s Schematic) GetSchematicNumbers() []SchematicNumber {
	numbers := []SchematicNumber{}
	curValue := 0
	numberLength := 0
	for _, char := range s.Contents {
		if digit, err := strconv.Atoi(string(char)); err == nil {
			curValue = curValue*10 + digit
			numberLength++
		} else if curValue > 0 {
			numbers = append(numbers, SchematicNumber{Value: curValue})
			curValue = 0
		}
	}
	return numbers
}
