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

	return Filter[string](lines, func(line string, index int) bool {
		return len(line) != 0
	})
}
