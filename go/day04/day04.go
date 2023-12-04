package main

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 4")
	fmt.Println("first part:")
	input := utils.Reader(2023, 04)
	scratchCards := getScratchCards(strings.Split(input, "\n"))
	totalPoints := sumCardsPoints(scratchCards)
	fmt.Println(totalPoints)

	fmt.Println("second part:")
	fmt.Println(sumTotalScratchCardCopies(scratchCards))
}

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
		if number, err := strconv.Atoi(maybeWinningNumberStr); err == nil {
			numbers.Add(number)
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
	pattern := regexp.MustCompile("Card\\s+(?P<id>\\d+): (?P<winning>[\\d\\s]+) \\| (?P<candidates>[\\d\\s]+)")
	match := pattern.FindStringSubmatch(input)
	id, _ := strconv.Atoi(match[pattern.SubexpIndex("id")])
	scratchCard.ID = id
	scratchCard.WinningNumbers = parseNumbers(match[pattern.SubexpIndex("winning")])
	scratchCard.CandidateNumbers = parseNumbers(match[pattern.SubexpIndex("candidates")])
	return
}

func sumTotalScratchCardCopies(scratchCards []ScratchCard) (totalCopies int) {
	scratchCardCopies := make(map[int]int)
	var lastScratchCard ScratchCard
	for _, scratchCard := range scratchCards {
		scratchCardCopies[scratchCard.ID]++
		winners := scratchCard.WinningNumbers.Intersection(scratchCard.CandidateNumbers)
		for i := scratchCard.ID + 1; i <= scratchCard.ID+winners.Len(); i++ {
			scratchCardCopies[i] += scratchCardCopies[scratchCard.ID]
		}
		lastScratchCard = scratchCard
	}
	for i := 1; i <= lastScratchCard.ID; i++ {
		totalCopies += scratchCardCopies[i]
	}
	return
}
