package main

import (
	"adventOfCode/utils"
	"fmt"
	"strings"
	"testing"
)

var inputTest = strings.TrimSpace(`
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`)

func TestGetSeeds(t *testing.T) {
	seeds := NewSeeds(splitInput(inputTest)[0])
	expect := Seeds{79, 14, 55, 13}
	utils.AssertDeepEqual(t, seeds, expect)
}

func TestMakeMap(t *testing.T) {
	testCases := []utils.TestCase[[]string, SrcDestMap]{
		{
			Case: []string{
				"50 98 2",
			},
			Expected: SrcDestMap{
				{50, 98, 2},
			},
		},
		{
			Case: []string{
				"50 98 2",
				"42 0 7",
			},
			Expected: SrcDestMap{
				{50, 98, 2},
				{42, 0, 7},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%+v", testCase.Case), func(t *testing.T) {
			srcDestMap := NewMap()
			AddAllToMap(srcDestMap, testCase.Case)
			utils.AssertDeepEqual(t, *srcDestMap, testCase.Expected)
		})
	}
}

func TestGetResourcesMap(t *testing.T) {
	inputTestSplit := splitInput(inputTest)[1:]
	testCases := []utils.TestCase[string, ResourcesMap]{
		{
			Case: inputTestSplit[0],
			Expected: ResourcesMap{
				From: "seed",
				To:   "soil",
				Map: *AddAllToMap(
					NewMap(),
					[]string{
						"50 98 2",
						"52 50 48",
					},
				),
			},
		},
		{
			Case: inputTestSplit[1],
			Expected: ResourcesMap{
				From: "soil",
				To:   "fertilizer",
				Map: *AddAllToMap(
					NewMap(),
					[]string{
						"0 15 37",
						"37 52 2",
						"39 0 15",
					},
				),
			},
		},
		{
			Case: inputTestSplit[2],
			Expected: ResourcesMap{
				From: "fertilizer",
				To:   "water",
				Map: *AddAllToMap(
					NewMap(),
					[]string{
						"49 53 8",
						"0 11 42",
						"42 0 7",
						"57 7 4",
					},
				),
			},
		},
		{
			Case: inputTestSplit[3],
			Expected: ResourcesMap{
				From: "water",
				To:   "light",
				Map: *AddAllToMap(
					NewMap(),
					[]string{
						"88 18 7",
						"18 25 70",
					},
				),
			},
		},
		{
			Case: inputTestSplit[4],
			Expected: ResourcesMap{
				From: "light",
				To:   "temperature",
				Map: *AddAllToMap(
					NewMap(),
					[]string{
						"45 77 23",
						"81 45 19",
						"68 64 13",
					},
				),
			},
		},
		{
			Case: inputTestSplit[5],
			Expected: ResourcesMap{
				From: "temperature",
				To:   "humidity",
				Map: *AddAllToMap(
					NewMap(),
					[]string{
						"0 69 1",
						"1 0 69",
					},
				),
			},
		},
		{
			Case: inputTestSplit[6],
			Expected: ResourcesMap{
				From: "humidity",
				To:   "location",
				Map: *AddAllToMap(
					NewMap(),
					[]string{
						"60 56 37",
						"56 93 4",
					},
				),
			},
		},
	}

	for _, testCase := range testCases {
		result := NewResourcesMap(testCase.Case)
		utils.AssertDeepEqual(t, result, testCase.Expected)
	}
}

func TestShallowMapTraversal(t *testing.T) {
	inputs := GetResourcesMaps(splitInput(inputTest)[1:])
	type TestCase struct {
		resourceMap ResourcesMap
		sources     []int
	}
	testCases := []utils.TestCase[TestCase, []int]{
		{
			Case: TestCase{
				resourceMap: inputs[0],
				sources:     []int{98, 51, 2},
			},
			Expected: []int{50, 53, 2},
		},
		{
			Case: TestCase{
				resourceMap: inputs[1],
				sources:     []int{16, 53, 14, 55},
			},
			Expected: []int{1, 38, 53, 55},
		},
		{
			Case: TestCase{
				resourceMap: inputs[2],
				sources:     []int{54, 12, 0, 8, 61},
			},
			Expected: []int{50, 1, 42, 58, 61},
		},
		{
			Case: TestCase{
				resourceMap: inputs[3],
				sources:     []int{24, 94, 2},
			},
			Expected: []int{94, 87, 2},
		},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprint(testCase.Case), func(t *testing.T) {
			results := []int{}
			for _, source := range testCase.Case.sources {
				results = append(results, testCase.Case.resourceMap.GetDestination(source))
			}
			utils.AssertDeepEqual(t, results, testCase.Expected)
		})
	}
}

func TestFullTraversal(t *testing.T) {
	inputTestSplit := splitInput(inputTest)
	seeds := NewSeeds(inputTestSplit[0])
	resourcesMaps := GetResourcesMaps(inputTestSplit[1:])
	testCases := []utils.TestCase[int, int]{
		{Case: seeds[0], Expected: 82},
		{Case: seeds[1], Expected: 43},
		{Case: seeds[2], Expected: 86},
		{Case: seeds[3], Expected: 35},
	}
	for _, testCase := range testCases {
		result := resourcesMaps.Traverse(testCase.Case)
		utils.AssertInt(t, result, testCase.Expected)
	}
}

func TestLowestLocation(t *testing.T) {
	inputTestSplit := splitInput(inputTest)
	seeds := NewSeeds(inputTestSplit[0])
	resourcesMaps := GetResourcesMaps(inputTestSplit[1:])
	expected := 35
	result := getLowestLocation(seeds, resourcesMaps)
	utils.AssertInt(t, result, expected)
}
