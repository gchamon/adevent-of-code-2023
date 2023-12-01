package utils

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"time"
)

func GetAOCInput(year, day int) []byte {
	aocCookie := os.Getenv("AOC_SESSION_COOKIE")
	aocUrl := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", aocUrl, nil)
	if err != nil {
		log.Fatalf("Error making new request: %s", err)
	}
	sessionCookie := http.Cookie{
		Name:  "session",
		Value: aocCookie,
	}
	req.AddCookie(&sessionCookie)
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("error executing request: %s", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading response body: %s", err)
	}

	return body
}

func GetAvailableDays(year int) int {
	today := time.Now()
	aocBegin := fmt.Sprintf("%d-12-01", year)
	minDate, _ := time.Parse("2006-01-02", aocBegin)
	diff := today.Sub(minDate)
	availableDays := int(math.Min(math.Ceil(diff.Hours()/24), 24))

	return availableDays
}
