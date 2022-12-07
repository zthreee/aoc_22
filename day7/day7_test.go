package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func ReadInput() string {
	input, err := ioutil.ReadFile("input_test.txt")

	if err != nil {
		log.Fatal(err)
	}

	return string(input)
}

func Test_partOne(t *testing.T) {
	result := partOne(ReadInput())
	if result != 95437 {
		t.Errorf("Expected 95437, got %d", result)
	}
}

func Test_partTwo(t *testing.T) {
	result := partTwo(ReadInput())
	if result != 24933642 {
		t.Errorf("Expected 24933642, got %d", result)
	}
}
