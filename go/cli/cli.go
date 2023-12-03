package main

import (
	"adventOfCode/utils"
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	year := flag.Int("year", 2023, "Year to download inputs")
	flag.Parse()

	for day := 1; day <= utils.GetAvailableDays(*year); day++ {
		fileName := utils.GetInputFileName(*year, day)
		if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Downloading day %d...\n", day)
			aocInput := utils.GetAOCInput(*year, day)
			utils.WriteToFile(aocInput, fileName)
		} else {
			fmt.Printf("Input for day %d already exists. Skipping...\n", day)
		}
	}
	fmt.Println("done")
}
