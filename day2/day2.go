package main

import (
	"io/ioutil"
	"log"
	"strings"
)

const (
	Rock     int    = 1
	Paper           = 2
	Scissors        = 3
	X        int    = 1
	Y               = 2
	Z               = 3
	WIN             = 6
	LOSE            = 0
	DRAW            = 3
	MustWin  string = "Z"
	MustDraw        = "Y"
	MustLose        = "X"
)

func main() {
	input, err := ioutil.ReadFile("day2/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	partOne(input)
	partTwo(input)
}

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
