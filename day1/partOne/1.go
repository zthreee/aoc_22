package partOne

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Resolve() {
	input, err := ioutil.ReadFile("day1/partOne/input.txt")

	if err != nil {
		log.Fatal(err)
	}

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
