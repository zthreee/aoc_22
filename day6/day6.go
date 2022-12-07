package main

import (
	"github.com/mdwhatcott/go-collections/set"
	"io/ioutil"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("day6/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Day 6 [1]: ", partOne(string(input)))
	log.Println("Day 6 [2]: ", partTwo(string(input)))
}

func partOne(input string) int {
	for x := 4; x < len(input); x++ {
		if set.From([]rune(input[x-4:x])...).Len() == 4 {
			return x
		}
	}

	panic("No solution found.")
}

func partTwo(input string) int {
	for i := 0; i < len(input); i++ {
		fourLetters := input[i : i+14]
		lettersMap := make(map[rune]bool)
		for _, letter := range fourLetters {
			lettersMap[letter] = true
		}

		if len(lettersMap) == 14 {
			return i + 14
		}
	}

	return 0
}
