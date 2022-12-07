package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("day4/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	log.Println("Day 4 [1]: ", partOne(lines))
	log.Println("Day 4 [2]: ", partTwo(lines))
}

func partOne(lines []string) int {
	count := 0

	for _, line := range lines {
		var startAreaOne, endAreaOne, startAreaTwo, endAreaTwo int

		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &startAreaOne, &endAreaOne, &startAreaTwo, &endAreaTwo)

		if err != nil {
			return 0
		}

		if startAreaOne <= startAreaTwo && endAreaOne >= endAreaTwo || startAreaTwo <= startAreaOne && endAreaTwo >= endAreaOne {
			count++
		}
	}

	return count
}

func partTwo(lines []string) int {
	count := 0

	for _, line := range lines {
		var startAreaOne, endAreaOne, startAreaTwo, endAreaTwo int

		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &startAreaOne, &endAreaOne, &startAreaTwo, &endAreaTwo)

		if err != nil {
			return 0
		}

		if startAreaOne <= endAreaTwo && endAreaOne >= startAreaTwo {
			count++
		}
	}

	return count
}
