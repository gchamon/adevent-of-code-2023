package main

import (
	"adventOfCode/utils"
	"errors"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Day 3")
	fmt.Println("first part:")
	input := utils.Reader(2023, 03)
	schematic := NewSchematic(input)
	schematicNumbers := schematic.GetSchematicNumbers()
	partNumbersSum := sumPartNumbers(schematicNumbers)
	fmt.Println(partNumbersSum)

}

type SchematicNumber struct {
	Value           int
	AdjacentSymbols map[string]bool
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

func NewSchematicNumber(value int) (s SchematicNumber) {
	s.Value = value
	s.AdjacentSymbols = make(map[string]bool)
	return
}

func (s Schematic) GetSchematicNumbers() []SchematicNumber {
	numbers := []SchematicNumber{}
	currentValue := 0
	numberLength := 0

	// while the schematic contents is a single string, it is
	// useful to maintain logical coordinates to symbols in order to
	// map adjacent symbols
	x, y := 0, 0
	incrementCoordinates := func(c rune) {
		if c == rune('\n') {
			x, y = 0, y+1
		} else {
			x++
		}
	}
	newSchematicNumberWithAdjacentSymbols := func() (number SchematicNumber) {
		number = NewSchematicNumber(currentValue)
		getSymbolIfPossible := func(cx, cy int) {
			if symbol, err := s.GetSymbol(cx, cy); err == nil {
				number.AdjacentSymbols[string(symbol)] = true
			}
		}
		getSymbolIfPossible(x, y)                // symbol in front of number
		getSymbolIfPossible(x-numberLength-1, y) // symbol behind number
		for i := x - numberLength - 1; i <= x; i++ {
			getSymbolIfPossible(i, y-1) // symbols in the row on top of number
			getSymbolIfPossible(i, y+1) // symbols in the row under the bumber
		}
		return
	}

	for _, char := range s.Contents {
		if digit, err := strconv.Atoi(string(char)); err == nil {
			currentValue = currentValue*10 + digit
			numberLength++
		} else if currentValue > 0 {
			currentNumber := newSchematicNumberWithAdjacentSymbols()
			numbers = append(numbers, currentNumber)
			currentValue = 0
			numberLength = 0
		}
		incrementCoordinates(char)
	}
	return numbers
}

func (n SchematicNumber) IsPartNumber() bool {
	for symbol := range n.AdjacentSymbols {
		if symbol != "." {
			return true
		}
	}
	return false
}

func sumPartNumbers(schematicNumbers []SchematicNumber) (partNumbersSum int) {
	for _, schematicNumber := range schematicNumbers {
		if schematicNumber.IsPartNumber() {
			partNumbersSum += schematicNumber.Value
		}
	}
	return
}
