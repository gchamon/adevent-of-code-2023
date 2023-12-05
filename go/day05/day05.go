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

func NewResourcesMap(input string) ResourcesMap {
	mapInputs := strings.Split(input, "\n")
	labelPattern := regexp.MustCompile("(?P<source>\\w+)-to-(?P<destination>\\w+) map:")
	match := labelPattern.FindStringSubmatch(mapInputs[0])

	newMap := NewMap()
	resourcesMap := ResourcesMap{
		From: match[labelPattern.SubexpIndex("source")],
		To:   match[labelPattern.SubexpIndex("destination")],
		Map:  *newMap,
	}
	return resourcesMap
}

func NewMap() *SrcDestMap {
	m := make(SrcDestMap)
	return &m
}

func (m *SrcDestMap) AddToMap(dest, src, rng int) {
	for i := 0; i < rng; i++ {
		(*m)[src+i] = dest + i
	}
}

func AddAllToMap(srcDestMap *SrcDestMap, ranges [][3]int) *SrcDestMap {
	for _, rng := range ranges {
		destination, source, rangMax := rng[0], rng[1], rng[2]
		srcDestMap.AddToMap(destination, source, rangMax)
	}
	return srcDestMap
}

var NewSeedsValidationError = errors.New("invalid input")

func NewSeeds(input string) (seeds Seeds, err error) {
	pattern := regexp.MustCompile("seeds: ([\\d\\s]+)")
	if match := pattern.FindStringSubmatch(input); len(match) == 0 {
		seeds = Seeds{}
		err = NewSeedsValidationError
	} else {
		seeds = parseIntList(match[1])
		err = nil
	}
	return
}

func parseIntList(input string) (output []int) {
	maybeInts := strings.Split(input, " ")
	for _, maybeInt := range maybeInts {
		if intToAdd, err := strconv.Atoi(maybeInt); err == nil {
			output = append(output, intToAdd)
		}
	}
	return
}

func splitInput(input string) []string {
	return strings.Split(input, "\n\n")
}
