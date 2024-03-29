package day6

import (
	"github.com/nothub/AoC2023/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	log.SetFlags(0)
}

var spaceinator = regexp.MustCompile("\\s+")

func level1(path string) (result int) {
	lines := utils.ReadLines(path)

	var time []int
	var dist []int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		line = spaceinator.ReplaceAllString(line, " ")
		vals := strings.Split(line, " ")
		for _, val := range vals {
			num, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			if strings.HasPrefix(line, "Time") {
				time = append(time, num)
			}
			if strings.HasPrefix(line, "Distance") {
				dist = append(dist, num)
			}
		}
	}

	for i := 0; i < len(time); i++ {
		t := time[i]
		d := dist[i]
		wins := 0
		for step := 1; step < t-1; step++ {
			if (t-step)*step > d {
				wins++
			}
		}
		if result == 0 {
			result = wins
		} else {
			result = result * wins
		}
	}

	return result
}

func level2(path string) (result int) {
	lines := utils.ReadLines(path)

	var time int
	var dist int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		line = spaceinator.ReplaceAllString(line, "")
		val := strings.Split(line, ":")[1]
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalln(err.Error())
		}
		if strings.HasPrefix(line, "Time") {
			time = num
		}
		if strings.HasPrefix(line, "Distance") {
			dist = num
		}
	}

	for step := 1; step < time-1; step++ {
		if (time-step)*step > dist {
			result++
		}
	}

	return result
}
