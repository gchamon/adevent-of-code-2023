package main

import (
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
