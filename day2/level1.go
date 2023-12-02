package main

import (
	"fmt"
	"log"
	"os"
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
		legit := true
		split := strings.Split(line, ": ")
		gameId, _ := strconv.Atoi(strings.Split(split[0], " ")[1])
		games := strings.Split(split[1], "; ")
		for _, game := range games {
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

	fmt.Println(result)
}
