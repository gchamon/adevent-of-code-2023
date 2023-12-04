package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type ScratchCard struct {
	ID int
}

func NewScratchCard(input string) (scratchCard ScratchCard) {
	pattern := regexp.MustCompile("Card (?P<id>\\d+): (?P<winning>[\\d\\s]+) | (?P<candidates>\\d\\s+)")
	match := pattern.FindStringSubmatch(input)
	fmt.Printf("%+v\n", match)
	id, _ := strconv.Atoi(match[pattern.SubexpIndex("id")])
	scratchCard.ID = id
	return
}
