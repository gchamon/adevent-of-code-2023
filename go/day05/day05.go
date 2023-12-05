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

func NewMap() (m SrcDestMap) {
	m = make(SrcDestMap)
	return
}

func (m *SrcDestMap) AddToMap(dest, src, rng int) {
	for i := 0; i < rng; i++ {
		(*m)[src+i] = dest + i
	}
}

func AddAllToMap(srcDestMap *SrcDestMap, ranges [][3]int) {
	for _, rng := range ranges {
		destination, source, rangMax := rng[0], rng[1], rng[2]
		srcDestMap.AddToMap(destination, source, rangMax)
	}
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
