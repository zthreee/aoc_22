package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("day1/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	partOne(input)
	partTwo(input)
}

func partOne(input []byte) {
	lines := strings.Split(string(input), "\n")
	elfCalories := 0
	max := 0
	for _, line := range lines {
		if line == "" {
			if elfCalories > max {
				max = elfCalories
			}
			elfCalories = 0
		}

		calories, _ := strconv.Atoi(line)
		elfCalories += calories
	}

	log.Println("Day 1 [1]: ", max)
}

func partTwo(input []byte) {
	lines := strings.Split(string(input), "\n")
	var elfsCalories []int
	elfCalories := 0
	for _, line := range lines {
		if line == "" {
			elfsCalories = append(elfsCalories, elfCalories)
			elfCalories = 0
		}

		calories, _ := strconv.Atoi(line)
		elfCalories += calories
	}

	sort.Ints(elfsCalories)
	top3Calories := 0

	for _, calories := range elfsCalories[len(elfsCalories)-3:] {
		top3Calories += calories
	}

	log.Println("Day 1 [2]: ", top3Calories)
}
