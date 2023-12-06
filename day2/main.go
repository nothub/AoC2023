package day2

import (
	"github.com/nothub/AoC2023/utils"
	"log"
	"strconv"
	"strings"
)

const (
	maxReds   = 12
	maxGreens = 13
	maxBlues  = 14
)

func init() {
	log.SetFlags(0)
}

func level1() (result int) {
	for _, line := range utils.ReadLines("input.txt") {
		if len(line) < 1 {
			continue
		}
		legit := true
		split := strings.Split(line, ": ")
		gameId, _ := strconv.Atoi(strings.Split(split[0], " ")[1])
		rounds := strings.Split(split[1], "; ")
		for _, game := range rounds {
			throws := strings.Split(game, ", ")
			for _, throw := range throws {
				spl := strings.Split(throw, " ")
				count, _ := strconv.Atoi(spl[0])
				switch spl[1] {
				case "red":
					if count > maxReds {
						legit = false
					}
				case "green":
					if count > maxGreens {
						legit = false
					}
				case "blue":
					if count > maxBlues {
						legit = false
					}
				}
			}
		}
		if legit {
			result = result + gameId
		}
	}

	return result
}

func level2() (result int) {
	for _, line := range utils.ReadLines("input.txt") {
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

	return result
}
