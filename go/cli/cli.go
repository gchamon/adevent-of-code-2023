package main

import (
	"adventOfCode/utils"
	"flag"
	"fmt"
)

func main() {
	year := flag.Int("year", 2023, "Year to download inputs")
	flag.Parse()

	for day := 1; day <= utils.GetAvailableDays(*year); day++ {
		fmt.Printf("Downloading day %d...\n", day)
		fileName := utils.GetInputFileName(*year, day)
		aocInput := utils.GetAOCInput(*year, day)
		utils.WriteToFile(aocInput, fileName)
	}
	fmt.Print("done")
}
