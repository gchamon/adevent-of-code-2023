package utils

import (
	"log"
	"os"
	"path/filepath"
)

func WriteToFile(fileContents []byte, fileName string) {
	err := os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
	if err != nil {
		log.Fatalf("Error making dir for %s: %s", fileName, err)
	}
	err = os.WriteFile(fileName, fileContents, os.FileMode(0644))
	if err != nil {
		log.Fatalf("Error writting to file %s: %s", fileName, err)
	}
}
