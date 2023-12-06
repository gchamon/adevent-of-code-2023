package main

import (
	"adventOfCode/utils"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Day 5")
	fmt.Println("first part:")
	input := utils.Reader(2023, 05)
	inputSplit := splitInput(input)
	seeds := NewSeeds(inputSplit[0])
	resourcesMaps := GetResourcesMaps(inputSplit[1:])
	lowestLocation := getLowestLocation(seeds, resourcesMaps)
	fmt.Println(lowestLocation)

	fmt.Println("second part:")
	lowestLocationRange := getLowestLocationRange(seeds, resourcesMaps)
	fmt.Println(lowestLocationRange)
}

type Seeds []int

type SrcDest struct {
	Destination int
	Source      int
	Range       int
}

type SrcDestMap []SrcDest
type ResourcesMap struct {
	From string
	To   string
	Map  SrcDestMap
}

type ResourcesMaps []ResourcesMap

func getLowestLocationRange(seeds Seeds, maps ResourcesMaps) (lowestLocation int) {
	lowestLocationFloat := math.Inf(0)
	for i := 0; i < len(seeds)/2; i++ {
		seedStart := seeds[i*2]
		seedRange := seeds[i*2+1]
		for seed := seedStart; seed < seedStart+seedRange; seed++ {
			location := float64(maps.Traverse(seed))
			if location < lowestLocationFloat {
				lowestLocationFloat = location
			}
		}
	}
	return int(lowestLocationFloat)
}

func getLowestLocation(seeds Seeds, maps ResourcesMaps) (lowestLocation int) {
	lowestLocationFloat := math.Inf(0)
	for _, seed := range seeds {
		location := float64(maps.Traverse(seed))
		if location < lowestLocationFloat {
			lowestLocationFloat = location
		}
	}
	return int(lowestLocationFloat)
}

func (r *ResourcesMaps) Traverse(seed int) (location int) {
	source := seed
	for _, resourcesMap := range *r {
		destination := resourcesMap.GetDestination(source)
		source = destination
	}
	location = source
	return
}

func (r *ResourcesMap) GetDestination(source int) (destination int) {
	for _, srcDest := range r.Map {
		if source >= srcDest.Source && source < srcDest.Source+srcDest.Range {
			destination = srcDest.Destination + (source - srcDest.Source)
			return
		}
	}
	destination = source
	return
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
	m := SrcDestMap{}
	return &m
}

func AddAllToMap(srcDestMap *SrcDestMap, ranges []string) *SrcDestMap {
	for _, rng := range ranges {
		rngInt := utils.ParseIntList(rng)
		destination, source, rangMax := rngInt[0], rngInt[1], rngInt[2]
		(*srcDestMap) = append((*srcDestMap), SrcDest{destination, source, rangMax})
	}
	return srcDestMap
}

var NewSeedsValidationError = errors.New("invalid input")

func NewSeeds(input string) (seeds Seeds) {
	pattern := regexp.MustCompile("seeds: ([\\d\\s]+)")
	match := pattern.FindStringSubmatch(input)
	seeds = utils.ParseIntList(match[1])
	return
}

func splitInput(input string) []string {
	return strings.Split(input, "\n\n")
}
