package utils

import (
	"strconv"
	"strings"
)

// takes a string which is a list of integers separated by spaces and returns the corresponding array of ints
// "1 2 3" -> []int{1,2,3}
func ParseIntList(input string) (output []int) {
	maybeInts := strings.Split(input, " ")
	for _, maybeInt := range maybeInts {
		if intToAdd, err := strconv.Atoi(maybeInt); err == nil {
			output = append(output, intToAdd)
		}
	}
	return
}
