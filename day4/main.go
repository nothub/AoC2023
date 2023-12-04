package day4

import (
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

func init() {
	log.SetFlags(0)
}

var spaceinator = regexp.MustCompile(`\s+`)

func level1(path string) (result int) {

	byt, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	raw := string(byt)
	lines := strings.Split(raw, "\n")

	for _, line := range lines {

		if len(line) == 0 {
			continue
		}

		line = spaceinator.ReplaceAllString(line, " ")

		split := strings.Split(strings.Split(line, ": ")[1], " | ")
		left := strings.Split(split[0], " ")
		right := strings.Split(split[1], " ")

		points := 0

		for _, num := range left {
			if slices.Contains(right, num) {
				if points == 0 {
					points = 1
				} else {
					points = points * 2
				}
			}
		}

		result = result + points
	}

	return result
}
