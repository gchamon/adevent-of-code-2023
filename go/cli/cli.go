package main

import (
	"adventOfCode/utils"
	"flag"
)

func main() {
	year := flag.Int("year", 2023, "Year to download inputs")
	flag.Parse()

	for day := 1; day <= utils.GetAvailableDays(*year); day++ {
		fileName := utils.GetInputFileName(*year, day)
		aocInput := utils.GetAOCInput(*year, day)
		utils.WriteToFile(aocInput, fileName)
	}
}
