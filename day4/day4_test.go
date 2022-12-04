package main

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func readInput() []string {
	input, err := ioutil.ReadFile("input_test.txt")

	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(input), "\n")
}

func Test_partOne(t *testing.T) {
	score := partOne(readInput())
	if score != 2 {
		t.Errorf("Expected 2, got %d", score)
	}
}

func Test_partTwo(t *testing.T) {
	score := partTwo(readInput())
	if score != 4 {
		t.Errorf("Expected 2, got %d", score)
	}
}
