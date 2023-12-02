package utils

import (
	"os"
	"strings"
)

func Reader(year, day int) string {
	fileName := GetInputFileName(year, day)
	content, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(content))
}
