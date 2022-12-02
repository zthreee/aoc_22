package day2

import (
	"log"
	"strings"
)

// A for Rock, B for Paper, and C for Scissors
// X for Rock, Y for Paper, and Z for Scissors
// 1 for Rock, 2 for Paper, and 3 for Scissors
// 0 LOST, 3 DRAW, 6 WIN

const (
	X    int = 1
	Y        = 2
	Z        = 3
	WIN      = 6
	LOSE     = 0
	DRAW     = 3
)

func partOne(input []byte) {
	lines := strings.Split(string(input), "\n")
	score := 0
	for _, line := range lines {
		fields := strings.Split(line, " ")
		score += calculateScore(fields[0], fields[1])
	}

	log.Println("Day 2 [1]: ", score)
}

func calculateScore(first, second string) int {
	switch first {
	case "A":
		if second == "X" {
			return X + DRAW
		}

		if second == "Y" {
			return Y + WIN
		}

		if second == "Z" {
			return Z + LOSE
		}
	case "B":
		if second == "X" {
			return X + LOSE
		}

		if second == "Y" {
			return Y + DRAW
		}

		if second == "Z" {
			return Z + WIN
		}
	case "C":
		if second == "X" {
			return X + WIN
		}

		if second == "Y" {
			return Y + LOSE
		}

		if second == "Z" {
			return Z + DRAW
		}
	}

	return 0
}
