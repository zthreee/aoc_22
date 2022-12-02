package day2

import (
	"log"
	"strings"
)

const (
	Rock     int = 1
	Paper        = 2
	Scissors     = 3
)

const (
	MustWin  string = "Z"
	MustDraw        = "Y"
	MustLose        = "X"
)

func partTwo(input []byte) {
	lines := strings.Split(string(input), "\n")
	score := 0
	for _, line := range lines {
		fields := strings.Split(line, " ")
		score += calculateScore_2(fields[0], fields[1])
	}

	log.Println("Day 2 [2]: ", score)
}

func calculateScore_2(first, second string) int {
	switch first {
	// X LOSE, Y DRAW, Z WIN
	case "A":
		if second == MustWin {
			return Paper + WIN
		}

		if second == MustDraw {
			return Rock + DRAW
		}

		if second == MustLose {
			return Scissors + LOSE
		}
	case "B":
		if second == MustWin {
			return Scissors + WIN
		}

		if second == MustDraw {
			return Paper + DRAW
		}

		if second == MustLose {
			return Rock + LOSE
		}
	case "C":
		if second == MustWin {
			return Rock + WIN
		}

		if second == MustDraw {
			return Scissors + DRAW
		}

		if second == MustLose {
			return Paper + LOSE
		}
	}

	return 0
}
