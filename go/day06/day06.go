package main

import (
	"adventOfCode/utils"
	"strings"
)

type Race struct {
	Time     int
	Distance int
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
