package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Seeds []int

type SrcDestMap map[int]int

type ResourcesMap struct {
	From string
	To   string
	Map  SrcDestMap
}

var NewSeedsValidationError = errors.New("invalid input")

func MakeMap(dest, src, rng int) (resMap SrcDestMap) {
	resMap = make(SrcDestMap, rng)
	for i := 0; i < rng; i++ {
		resMap[src+i] = dest + i
	}
	return
}

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
