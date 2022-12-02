package day1

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

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
