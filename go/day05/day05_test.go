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
	seeds, err := NewSeeds(splitInput(inputTest)[0])
	expect := Seeds{79, 14, 55, 13}
	utils.AssertNotError(t, err)
	utils.AssertDeepEqual(t, seeds, expect)
}

func TestMakeMap(t *testing.T) {
	testCases := []utils.TestCase[[][3]int, map[int]int]{
		{
			Case: [][3]int{
				{50, 98, 2},
			},
			Expected: map[int]int{
				98: 50,
				99: 51,
			},
		},
		{
			Case: [][3]int{
				{42, 0, 7},
			},
			Expected: map[int]int{
				0: 42,
				1: 43,
				2: 44,
				3: 45,
				4: 46,
				5: 47,
				6: 48,
			},
		},
		{
			Case: [][3]int{
				{50, 98, 2},
				{42, 0, 7},
			},
			Expected: map[int]int{
				0:  42,
				1:  43,
				2:  44,
				3:  45,
				4:  46,
				5:  47,
				6:  48,
				98: 50,
				99: 51,
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%+v", testCase.Case), func(t *testing.T) {
			srcDestMap := NewMap()
			AddAllToMap(&srcDestMap, testCase.Case)
			utils.AssertDeepEqual(t, srcDestMap, testCase.Expected)
		})
	}
}

// func TestGetResourcesMap(t *testing.T) {
// 	inputTestSplit := splitInput(inputTest)[1:]
// 	testCases := []utils.TestCase[string, ResourcesMap]{
// 		{
// 			Case: inputTestSplit[0],
// 			Expected: ResourcesMap{
// 				From: "seed",
// 				To: "soil",
// 				Map:
// 			}
// 		}
// 	}
// }
