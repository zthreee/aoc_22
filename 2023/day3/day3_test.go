package main

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	dat, err := os.ReadFile("../input.txt")
	if err != nil {
		panic("Unable to load file: " + err.Error())
	}
	value := part2Optim(strings.Split(string(dat), "\n"))

	expect := 36749103
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}

	value = part2Optim(strings.Split(string(dat), "\n"))
	log.Println("The answer is", value)
}
