package utils

import (
	"log"
	"os"
	"strings"
)

func ReadLines(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err.Error())
	}

	lines := strings.Split(string(data), "\n")
	for i := range lines {
		if len(lines[i]) == 0 {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}

	return lines
}
