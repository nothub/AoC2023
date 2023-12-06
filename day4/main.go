package day4

import (
	"github.com/nothub/AoC2023/utils"
	"log"
	"regexp"
	"slices"
	"strings"
)

func init() {
	log.SetFlags(0)
}

type card struct {
	left  []string
	right []string
	count int
}

var spaceinator = regexp.MustCompile("\\s+")

func level1(path string) (result int) {
	for _, card := range readCards(path) {
		points := 0
		for _, num := range card.left {
			if slices.Contains(card.right, num) {
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

func level2(path string) (result int) {
	cards := readCards(path)
	for i, card := range cards {
		points := 0
		for _, num := range card.left {
			if slices.Contains(card.right, num) {
				if points == 0 {
					points = 1
				} else {
					points = points + 1
				}
			}
		}
		for j := 1; j <= points; j++ {
			cards[i+j].count = cards[i+j].count + card.count
		}
	}
	for _, card := range cards {
		result = result + card.count
	}
	return result
}

func readCards(path string) (cards []card) {
	lines := utils.ReadLines(path)

	for _, line := range lines {

		if len(line) == 0 {
			continue
		}
		line = spaceinator.ReplaceAllString(line, " ")

		split := strings.Split(strings.Split(line, ": ")[1], " | ")
		left := strings.Split(split[0], " ")
		right := strings.Split(split[1], " ")

		cards = append(cards, card{
			left:  left,
			right: right,
			count: 1,
		})

	}

	return cards
}
