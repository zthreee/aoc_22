package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) Push(st ...string) {
	*s = append(*s, st...)
}

func (s *Stack) Peek() string {
	i := len(*s) - 1
	return (*s)[i]
}

func (s *Stack) Pop() string {
	i := len(*s) - 1
	elem := (*s)[i]
	*s = (*s)[:i]
	return elem
}

func (s *Stack) PopN(amount int) []string {
	i := len(*s) - amount
	elems := (*s)[i:]
	*s = (*s)[:i]
	return elems
}

func (s *Stack) Len() int {
	return len(*s)
}

type Step struct {
	Amount, From, To int
}

func main() {
	input, err := ioutil.ReadFile("day5/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Day 5 [1]: ", partOne(string(input)))
	log.Println("Day 5 [2]: ", partTwo(string(input)))
}

func partOne(input string) string {
	parsedStacks, parsedSteps, found := strings.Cut(input, "\n\n")
	if !found {
		return ""
	}

	var stacks []*Stack
	var steps []Step

	stacksStrings := strings.Split(parsedStacks, "\n")
	// count the number of integers in the final row
	cratesNumbers := strings.Fields(stacksStrings[len(stacksStrings)-1])
	lastCrateNumber, _ := strconv.Atoi(cratesNumbers[len(cratesNumbers)-1])

	log.Println("stacksStrings: ", stacksStrings)
	log.Println(" lastCrateNumber: ", lastCrateNumber)

	stacks = make([]*Stack, lastCrateNumber)
	for i := range stacks {
		stacks[i] = new(Stack)
	}

	for i := len(stacksStrings) - 2; i >= 0; i-- {
		stackString := stacksStrings[i]
		var current string
		for i := 0; i < lastCrateNumber; i++ {
			if len(stackString) < 4 {
				current = stackString
			} else {
				current, stackString = stackString[:4], stackString[4:]
			}
			var crate string
			n, err := fmt.Sscanf(current, "[%1s]", &crate)
			if n != 1 || err != nil {
				continue
			}
			stacks[i].Push(crate)
		}
	}

	stepsStrings := strings.Split(parsedSteps, "\n")
	for _, stepString := range stepsStrings {
		var step Step
		_, err := fmt.Sscanf(stepString, "move %d from %d to %d", &step.Amount, &step.From, &step.To)
		if err != nil {
			return ""
		}
		// to zero index it all
		step.From--
		step.To--
		steps = append(steps, step)
	}

	endStacks := ProcessSteps(stacks, steps)

	var result string
	for _, stack := range endStacks {
		result = result + stack.Peek()
	}

	return result
}

func partTwo(input string) string {
	parsedStacks, parsedSteps, found := strings.Cut(input, "\n\n")
	if !found {
		return ""
	}

	var stacks []*Stack
	var steps []Step

	stacksStrings := strings.Split(parsedStacks, "\n")
	cratesNumbers := strings.Fields(stacksStrings[len(stacksStrings)-1])
	lastCrateNumber, _ := strconv.Atoi(cratesNumbers[len(cratesNumbers)-1])

	log.Println("stacksStrings: ", stacksStrings)
	log.Println(" lastCrateNumber: ", lastCrateNumber)

	stacks = make([]*Stack, lastCrateNumber)
	for i := range stacks {
		stacks[i] = new(Stack)
	}

	for i := len(stacksStrings) - 2; i >= 0; i-- {
		stackString := stacksStrings[i]
		var current string
		for i := 0; i < lastCrateNumber; i++ {
			if len(stackString) < 4 {
				current = stackString
			} else {
				current, stackString = stackString[:4], stackString[4:]
			}
			var crate string
			n, err := fmt.Sscanf(current, "[%1s]", &crate)
			if n != 1 || err != nil {
				continue
			}
			stacks[i].Push(crate)
		}
	}

	stepsStrings := strings.Split(parsedSteps, "\n")
	for _, stepString := range stepsStrings {
		var step Step
		_, err := fmt.Sscanf(stepString, "move %d from %d to %d", &step.Amount, &step.From, &step.To)
		if err != nil {
			return ""
		}

		step.From--
		step.To--
		steps = append(steps, step)
	}

	endStacks := ProcessSteps9001(stacks, steps)

	var result string
	for _, stack := range endStacks {
		result = result + stack.Peek()
	}

	return result
}

func ProcessSteps(stacks []*Stack, steps []Step) []*Stack {
	for _, step := range steps {
		for i := 0; i < step.Amount; i++ {
			from := stacks[step.From]
			to := stacks[step.To]
			crate := from.Pop()
			to.Push(crate)
		}
	}
	return stacks
}

func ProcessSteps9001(stacks []*Stack, steps []Step) []*Stack {
	for _, step := range steps {
		from := stacks[step.From]
		to := stacks[step.To]
		crates := from.PopN(step.Amount)
		to.Push(crates...)
	}
	return stacks
}
