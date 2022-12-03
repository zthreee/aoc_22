package day3

import (
	"io/ioutil"
	"log"
)

func Day3() {
	input, err := ioutil.ReadFile("day3/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	partOne(input)
	partTwo(input)
}
