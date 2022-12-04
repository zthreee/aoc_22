package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("day3/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	partOne(input)
	partTwo(input)
}

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

func addPriority(c int32) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	}

	return int(c - 'A' + 27)
}

func commonCharBetweenRucksacks(rucksacks []string) int32 {
	for _, char := range rucksacks[0] {
		if strings.ContainsRune(rucksacks[1], char) && strings.ContainsRune(rucksacks[2], char) {
			return char
		}
	}

	return 0
}
