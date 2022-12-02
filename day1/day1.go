package day1

import (
	"io/ioutil"
	"log"
)

func Day1() {
	input, err := ioutil.ReadFile("day1/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	partOne(input)
	partTwo(input)
}
