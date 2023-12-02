package main

import (
	"regexp"
	"strconv"
	"strings"
)

type CubesSubset struct {
	Red   int
	Blue  int
	Green int
}

type CubeGame struct {
	Id           int
	CubesSubsets []CubesSubset
}

func parseGame(input string) CubeGame {
	cubeGame := CubeGame{}
	pattern := regexp.MustCompile("Game (?P<id>[0-9]+): (?P<subsets>.*)")
	match := pattern.FindStringSubmatch(input)

	gameId := match[pattern.SubexpIndex("id")]
	cubeGame.Id, _ = strconv.Atoi(gameId)

	gameSubsets := match[pattern.SubexpIndex("subsets")]
	for _, gameSubset := range strings.Split(gameSubsets, ";") {
		cubesSubset := CubesSubset{}
		for _, gameColor := range strings.Split(gameSubset, ",") {
			gameColorSplit := strings.Split(strings.TrimSpace(gameColor), " ")
			switch gameColorSplit[1] {
			case "red":
				cubesSubset.Red, _ = strconv.Atoi(gameColorSplit[0])
			case "blue":
				cubesSubset.Blue, _ = strconv.Atoi(gameColorSplit[0])
			case "green":
				cubesSubset.Green, _ = strconv.Atoi(gameColorSplit[0])
			}
		}
		cubeGame.CubesSubsets = append(cubeGame.CubesSubsets, cubesSubset)
	}

	return cubeGame
}
