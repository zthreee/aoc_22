package main

import (
	"testing"
)

func Test_partOne(t *testing.T) {
	result := partOne("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	if 7 != result {
		t.Errorf("Expected 7, got %d", result)
	}
	result = partOne("bvwbjplbgvbhsrlpgdmjqwftvncz")
	if 5 != result {
		t.Errorf("Expected 5, got %d", result)
	}
	result = partOne("nppdvjthqldpwncqszvftbrmjlhg")
	if 6 != result {
		t.Errorf("Expected 6, got %d", result)
	}
	result = partOne("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	if 10 != result {
		t.Errorf("Expected 10, got %d", result)
	}
	result = partOne("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
	if 11 != result {
		t.Errorf("Expected 11, got %d", result)
	}
}

func Test_partTwo(t *testing.T) {
	result := partTwo("")
	if 0 != result {
		t.Errorf("Expected 0, got %d", result)
	}
}
