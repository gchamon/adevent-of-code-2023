package main

import (
	"sort"
	"strconv"
	"strings"
)

type SchematicNumber struct {
	Value           int
	AdjacentSymbols map[rune]bool
}

func sortSchematicNumbers(schematicNumbers *[]SchematicNumber) {
	sort.SliceStable(*schematicNumbers, func(i, j int) bool {
		return (*schematicNumbers)[i].Value < (*schematicNumbers)[j].Value
	})
}

func getSchematicNumbers(input string) []SchematicNumber {
	numbers := []SchematicNumber{}
	curValue := 0
	for _, line := range strings.Split(input, "\n") {
		for _, char := range line {
			if digit, err := strconv.Atoi(string(char)); err == nil {
				curValue = curValue*10 + digit
			} else if curValue > 0 {
				numbers = append(numbers, SchematicNumber{Value: curValue})
				curValue = 0
			}
		}
		if curValue > 0 {
			numbers = append(numbers, SchematicNumber{Value: curValue})
			curValue = 0
		}
	}
	return numbers
}
