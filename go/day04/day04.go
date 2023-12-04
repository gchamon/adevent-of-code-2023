package main

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type ScratchCard struct {
	ID               int
	WinningNumbers   utils.Set[int]
	CandidateNumbers utils.Set[int]
}

func (s ScratchCard) GetPoints() int {
	winners := s.WinningNumbers.Intersection(s.CandidateNumbers)
	return int(math.Pow(2, float64(winners.Len())-1))
}

func sumCardsPoints(cards []ScratchCard) (points int) {
	for _, card := range cards {
		points += card.GetPoints()
	}
	return
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

func getScratchCards(input []string) (scratchCards []ScratchCard) {
	scratchCards = []ScratchCard{}

	for _, scratchCardString := range input {
		scratchCards = append(scratchCards, newScratchCard(scratchCardString))
	}

	return
}

func newScratchCard(input string) (scratchCard ScratchCard) {
	pattern := regexp.MustCompile("Card (?P<id>\\d+): (?P<winning>[\\d\\s]+) \\| (?P<candidates>[\\d\\s]+)")
	match := pattern.FindStringSubmatch(input)
	fmt.Printf("%+v\n", match)
	id, _ := strconv.Atoi(match[pattern.SubexpIndex("id")])
	scratchCard.ID = id
	scratchCard.WinningNumbers = parseNumbers(match[pattern.SubexpIndex("winning")])
	scratchCard.CandidateNumbers = parseNumbers(match[pattern.SubexpIndex("candidates")])
	return
}
