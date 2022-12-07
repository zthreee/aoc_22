package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func readInput() string {
	input, err := ioutil.ReadFile("input_test.txt")

	if err != nil {
		log.Fatal(err)
	}

	return string(input)
}

func Test_partOne(t *testing.T) {
	result := partOne(readInput())
	if result != "CMZ" {
		t.Errorf("Expected CMZ, got %s", result)
	}
}

func Test_partTwo(t *testing.T) {
	result := partTwo(readInput())
	if result != "MCD" {
		t.Errorf("Expected MCD, got %s", result)
	}
}
