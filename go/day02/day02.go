package main

import (
	"adventOfCode/utils"
	"fmt"
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

func main() {
	fmt.Println("Day 2")
	fmt.Println("first part:")
	input := utils.Reader(2023, 02)
	sumPossibleIds := 0
	for _, line := range strings.Split(input, "\n") {
		game := parseGame(line)
		if game.IsPossible() {
			sumPossibleIds += game.Id
		}
	}
	fmt.Println(sumPossibleIds)
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

func (g CubeGame) IsPossible() bool {
	MaxCubeSubset := CubesSubset{
		Red: 12, Green: 13, Blue: 14,
	}

	for _, subset := range g.CubesSubsets {
		if subset.Red > MaxCubeSubset.Red || subset.Green > MaxCubeSubset.Green || subset.Blue > MaxCubeSubset.Blue {
			return false
		}
	}
	return true
}

func (g CubeGame) GetMinimumSubset() (minimumGamePossible CubesSubset) {
	for _, subset := range g.CubesSubsets {
		if subset.Red > minimumGamePossible.Red {
			minimumGamePossible.Red = subset.Red
		}
		if subset.Green > minimumGamePossible.Green {
			minimumGamePossible.Green = subset.Green
		}
		if subset.Blue > minimumGamePossible.Blue {
			minimumGamePossible.Blue = subset.Blue
		}
	}
	return
}

func (s CubesSubset) CalculatePower() int {
	return 0
}
