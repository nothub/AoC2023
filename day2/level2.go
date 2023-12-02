package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	byt, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	raw := string(byt)
	lines := strings.Split(raw, "\n")

	result := 0
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}

		minReds := 0
		minGreens := 0
		minBlues := 0

		split := strings.Split(line, ": ")
		rounds := strings.Split(split[1], "; ")
		for _, game := range rounds {
			throws := strings.Split(game, ", ")
			for _, throw := range throws {
				spl := strings.Split(throw, " ")
				count, _ := strconv.Atoi(spl[0])
				switch spl[1] {
				case "red":
					if count > minReds {
						minReds = count
					}
				case "green":
					if count > minGreens {
						minGreens = count
					}
				case "blue":
					if count > minBlues {
						minBlues = count
					}
				}
			}
		}
		result = result + (minReds * minGreens * minBlues)
	}

	fmt.Println(result)
}
