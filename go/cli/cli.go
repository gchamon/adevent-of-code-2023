package main

import (
	"adventOfCode/utils"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
)

func main() {
	year := flag.Int("year", 2023, "Year to download inputs")
	flag.Parse()

	curPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working dir: %s", err)
	}

	for day := 1; day <= utils.GetAvailableDays(*year); day++ {
		fileName := path.Join(curPath, "../inputs", strconv.Itoa(*year), fmt.Sprintf("%02d", day), "input.txt")
		aocInput := utils.GetAOCInput(*year, day)
		utils.WriteToFile(aocInput, fileName)
	}
}
