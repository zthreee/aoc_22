package day2

import (
	"io/ioutil"
	"log"
)

func Day2() {
	input, err := ioutil.ReadFile("day2/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	partOne(input)
	partTwo(input)
}
