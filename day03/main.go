package main

import (
	"fmt"

	"github.com/matgd/advent2022/utils"
)

func getCharPriorities() map[string]int {
	priorities := map[string]int{}

	currentPriority := 1
	for ch := 'a'; ch <= 'z'; ch++ {
		priorities[string(ch)] = currentPriority
		currentPriority++
	}

	for ch := 'A'; ch <= 'Z'; ch++ {
		priorities[string(ch)] = currentPriority
		currentPriority++
	}

	return priorities
}

func splitRacksack(input string) (string, string) {
	first := input[:len(input)/2]
	second := input[len(input)/2:]

	return first, second
}

func getDuplicatePriority(firstRacksack, secondRacksack string) int {
	occurences := map[string]int{}

}

func Task1(inputFile string) int {
	lines := utils.ReadFileLines(inputFile)

	for _, line := range lines {
		first, second := splitRacksack(line)
	}

	return 0
}

func main() {
	fmt.Println("[Task 1]", Task1("input.txt"))
}
