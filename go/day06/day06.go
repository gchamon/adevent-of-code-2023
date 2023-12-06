package main

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("Day 6")
	fmt.Println("first part:")
	input := utils.Reader(2023, 06)
	races := parseInput(input)
	result := 1
	for _, race := range races {
		result = result * race.GetWaysToWin()
	}
	fmt.Println(result)
	fmt.Println("second part:")
	race := parseInput(strings.ReplaceAll(input, " ", ""))[0]
	fmt.Println(race.GetWaysToWin())
}

type Race struct {
	Time     int
	Distance int
}

// this leverages a second degree inequality equation that rises from the problem
// for every time X spent pressing the button, the resulting speed increases to x mm/ms
// the remaining time it is propelled, covering a total distance in time Y = T - X
// where T is total time
// total distance is Speed * Time = X*Y = S
// Total time is X + Y = T
// X*Y > S in order to win
// X*(T - X) > S
// -X² +XT -S > 0
// X1 = (T - √(T²-4S))/2
// X2 = (T + √(T²-4S))/2
// total ways to win are the integest between the bounds of X1 and X2
func (r Race) GetWaysToWin() int {
	getBound := func(c float64) float64 {
		return (float64(r.Time) + c*math.Sqrt(math.Pow(float64(r.Time), 2)-4*float64(r.Distance))) / 2
	}
	timeLowerBound := getBound(-1)
	timeUpperBound := getBound(+1)
	return int(math.Ceil(timeUpperBound) - math.Floor(timeLowerBound) - 1)
}

func parseInput(input string) []Race {
	timesAndDistancesString := strings.Split(input, "\n")
	times := utils.ParseIntList(strings.Split(timesAndDistancesString[0], ":")[1])
	distances := utils.ParseIntList(strings.Split(timesAndDistancesString[1], ":")[1])

	races := []Race{}
	for i, time := range times {
		races = append(races, Race{Time: time, Distance: distances[i]})
	}

	return races
}
