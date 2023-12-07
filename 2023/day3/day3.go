package main

import (
	file "aoc_22/pkg"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func process(input []string) int {
	defer file.Timer("Part 2")()
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(input[0], ":")[1], " ", ""))
	recordDistance, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(input[1], ":")[1], " ", ""))
	winners := 0
	for speed := 1; speed < time; speed++ {
		distance := (time - speed) * speed
		if distance > recordDistance {
			winners++
		}
	}
	return winners
}

func part2Optim(input []string) int {
	var result = 1
	startTime := time.Now()

	times, distances := parseLine(input[0]), parseLine(input[1])
	time1 := mergeIntoOneInt(times)
	distance := mergeIntoOneInt(distances)
	fmt.Println("check: ", time1, distance)

	winPresses := optimCalcWinPresses(time1, distance)
	result = winPresses

	fmt.Println("Execution time: ", time.Since(startTime))

	return result
}

func parseLine(line string) []int {
	parts := strings.Fields(line)[1:]
	els := make([]int, len(parts))
	for i, part := range parts {
		els[i], _ = strconv.Atoi(part)
	}
	return els
}

func optimCalcWinPresses(time int, distance int) int {

	firstWin := 0
	lastWin := 0

	for i := 0; i <= time; i++ {
		if firstWin == 0 && i*(time-i) > distance {
			firstWin = i
			lastWin = time - i
			break
		}

	}
	return lastWin - firstWin
}

func mergeIntoOneInt(els []int) int {
	mergedStr := ""
	for _, el := range els {
		mergedStr += strconv.Itoa(el)
	}
	merged, _ := strconv.Atoi(mergedStr)
	return merged
}
