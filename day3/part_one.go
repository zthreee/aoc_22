package day3

import (
	"log"
	"strings"
)

func partOne(input []byte) {
	lines := strings.Split(string(input), "\n")
	sum := 0

	for _, line := range lines {

		firstRucksack := line[0 : len(line)/2]
		secondRucksack := line[len(line)/2:]
		commonChars := make(map[rune]interface{})

		for _, char := range firstRucksack {
			if strings.ContainsRune(secondRucksack, char) {
				if _, ok := commonChars[char]; !ok {
					commonChars[char] = nil
					sum += addPriority(char)
				}
			}
		}
	}

	log.Println("Day 3 [1]: ", sum)
}

func addPriority(c int32) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	}

	return int(c - 'A' + 27)
}
