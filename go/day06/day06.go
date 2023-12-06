package main

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"strings"
)

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
	fmt.Printf("%+v\n", r)
	getBound := func(c float64) float64 {
		return (float64(r.Time) + c*math.Sqrt(math.Pow(float64(r.Time), 2)-4*float64(r.Distance))) / 2
	}
	getBoundInt := func(bound float64, round func(float64) float64, c int) int {
		boundRound := round(bound)
		if bound == boundRound {
			return int(bound) + c
		} else {
			return int(boundRound)
		}
	}
	timeLowerBound := getBound(-1)
	timeUpperBound := getBound(+1)
	fmt.Println(timeUpperBound, timeLowerBound)
	return getBoundInt(timeUpperBound, math.Floor, -1) - getBoundInt(timeLowerBound, math.Ceil, 1) + 1
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
