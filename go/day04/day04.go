package main

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type ScratchCard struct {
	ID               int
	WinningNumbers   utils.Set[int]
	CandidateNumbers utils.Set[int]
}

func parseNumbers(input string) (numbers utils.Set[int]) {
	numbers = utils.NewSet[int]()

	for _, maybeWinningNumberStr := range strings.Split(input, " ") {
		if winningNumber, err := strconv.Atoi(maybeWinningNumberStr); err == nil {
			numbers.Add(winningNumber)
		}
	}

	return
}

func NewScratchCard(input string) (scratchCard ScratchCard) {
	pattern := regexp.MustCompile("Card (?P<id>\\d+): (?P<winning>[\\d\\s]+) \\| (?P<candidates>[\\d\\s]+)")
	match := pattern.FindStringSubmatch(input)
	fmt.Printf("%+v\n", match)
	id, _ := strconv.Atoi(match[pattern.SubexpIndex("id")])
	scratchCard.ID = id
	scratchCard.WinningNumbers = parseNumbers(match[pattern.SubexpIndex("winning")])
	scratchCard.CandidateNumbers = parseNumbers(match[pattern.SubexpIndex("candidates")])
	return
}
