package day3

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func init() {
	log.SetFlags(0)
}

func input(path string) [][]rune {
	byt, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	raw := string(byt)
	lines := strings.Split(raw, "\n")
	var chars [][]rune
	for _, line := range lines {
		if len(line) > 0 {
			chars = append(chars, []rune(line))
		}
	}
	return chars
}

func level1(path string) (result int) {
	schematic := input(path)

	for i := 0; i < len(schematic); i++ {
		for j := 0; j < len(schematic[0]); j++ {
			if isNumber(schematic[i][j]) {
				l := 1
				for {
					if j+l >= len(schematic[0]) {
						break
					}
					if isNumber(schematic[i][j+l]) {
						l++
					} else {
						break
					}
				}

				start := j
				end := j + l - 1

				if isPartNumber(schematic, i, start, end) {
					n, err := strconv.Atoi(string(schematic[i][start : end+1]))
					if err != nil {
						log.Fatalln(err.Error())
					}
					result = result + n
				}

				j = end
			}
		}
	}

	return result
}

func isPartNumber(schematic [][]rune, line int, start int, end int) bool {
	for _, ri := range []int{-1, 0, 1, 2} {
		for j := start - 1; j <= end+1; j++ {
			i := line + ri
			if i < 0 || i >= len(schematic) {
				continue
			}
			if j < 0 || j >= len(schematic[0]) {
				continue
			}
			if isSymbol(schematic[i][j]) {
				return true
			}
		}
	}
	return false
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isSymbol(r rune) bool {
	if r == '.' {
		return false
	}
	return !isNumber(r)
}
