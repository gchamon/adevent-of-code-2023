package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Seeds []int

var NewSeedsValidationError = errors.New("invalid input")

func NewSeeds(input string) (seeds Seeds, err error) {
	pattern := regexp.MustCompile("seeds: ([\\d\\s]+)")
	if match := pattern.FindStringSubmatch(input); len(match) == 0 {
		seeds = Seeds{}
		err = NewSeedsValidationError
	} else {
		maybeSeeds := strings.Split(match[1], " ")
		err = nil
		for _, maybeSeed := range maybeSeeds {
			if seed, err := strconv.Atoi(maybeSeed); err == nil {
				seeds = append(seeds, seed)
			}
		}
	}
	return
}

func splitInput(input string) []string {
	return strings.Split(input, "\n\n")
}
