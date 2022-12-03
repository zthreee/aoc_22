package day3

import (
	"log"
	"strings"
)

func partTwo(input []byte) {
	lines := strings.Split(string(input), "\n")
	sum := 0
	var rucksacks []string
	for count, line := range lines {
		rucksacks = append(rucksacks, line)
		if (count+1)%3 == 0 {
			sum += addPriority(commonCharBetweenRucksacks(rucksacks))
			rucksacks = []string{}
		}
	}

	log.Println("Day 3 [2]: ", sum)
}

func commonCharBetweenRucksacks(rucksacks []string) int32 {
	for _, char := range rucksacks[0] {
		if strings.ContainsRune(rucksacks[1], char) && strings.ContainsRune(rucksacks[2], char) {
			return char
		}
	}

	return 0
}
