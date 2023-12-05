package main

import (
	"errors"
	"fmt"
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

type ResourcesMaps []ResourcesMap

func (r *ResourcesMaps) Traverse(seed int) (location int) {
	source := seed
	for _, resourcesMap := range *r {
		destination := resourcesMap.GetDestination(source)
		fmt.Println(resourcesMap.From, source, resourcesMap.To, destination)
		source = destination
	}
	fmt.Println("")
	location = source
	return
}

func (r *ResourcesMap) GetDestination(source int) int {
	if destination, ok := r.Map[source]; ok {
		return destination
	} else {
		return source
	}
}

func GetResourcesMaps(inputs []string) (resourcesMaps ResourcesMaps) {
	for _, input := range inputs {
		resourcesMaps = append(resourcesMaps, NewResourcesMap(input))
	}
	return
}

func NewResourcesMap(input string) ResourcesMap {
	mapInputs := strings.Split(input, "\n")
	labelPattern := regexp.MustCompile("(?P<source>\\w+)-to-(?P<destination>\\w+) map:")
	match := labelPattern.FindStringSubmatch(mapInputs[0])

	newMap := NewMap()
	AddAllToMap(newMap, mapInputs[1:])

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

func AddAllToMap(srcDestMap *SrcDestMap, ranges []string) *SrcDestMap {
	for _, rng := range ranges {
		rngInt := parseIntList(rng)
		destination, source, rangMax := rngInt[0], rngInt[1], rngInt[2]
		srcDestMap.AddToMap(destination, source, rangMax)
	}
	return srcDestMap
}

var NewSeedsValidationError = errors.New("invalid input")

func NewSeeds(input string) (seeds Seeds) {
	pattern := regexp.MustCompile("seeds: ([\\d\\s]+)")
	match := pattern.FindStringSubmatch(input)
	seeds = parseIntList(match[1])
	return
}

// takes a string which is a list of integers separated by spaces and returns the corresponding array of ints
// "1 2 3" -> []int{1,2,3}
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
