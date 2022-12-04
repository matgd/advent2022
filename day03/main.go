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

// Prevent creating multiple times and create one object instead.
var priorities map[string]int = getCharPriorities()

func splitRacksack(input string) (string, string) {
	first := input[:len(input)/2]
	second := input[len(input)/2:]

	return first, second
}

func getDuplicatePriorities(firstRacksack, secondRacksack string) int {
	occurences := map[string][2]bool{}
	totalPriority := 0

	for _, ch := range firstRacksack {
		occurences[string(ch)] = [2]bool{true, false}
	}
	for _, ch := range secondRacksack {
		occurences[string(ch)] = [2]bool{occurences[string(ch)][0], true}
	}

	for ch, occurence := range occurences {
		if occurence[0] && occurence[1] {
			totalPriority += priorities[ch]
		}
	}
	return totalPriority
}

func getDuplicatePriorities3(firstRacksack, secondRacksack, thirdRacksack string) int {
	occurences := map[string][3]bool{}
	totalPriority := 0

	for _, ch := range firstRacksack {
		occurences[string(ch)] = [3]bool{true, false, false}
	}
	for _, ch := range secondRacksack {
		occurences[string(ch)] = [3]bool{occurences[string(ch)][0], true, false}
	}
	for _, ch := range thirdRacksack {
		occurences[string(ch)] = [3]bool{occurences[string(ch)][0], occurences[string(ch)][1], true}
	}

	for ch, occurence := range occurences {
		if occurence[0] && occurence[1] && occurence[2] {
			totalPriority += priorities[ch]
		}
	}
	return totalPriority
}

func Task1(inputFile string) int {
	lines := utils.ReadFileLines(inputFile)

	totalPriority := 0
	for _, line := range lines {
		first, second := splitRacksack(line)
		totalPriority += getDuplicatePriorities(first, second)
	}

	return totalPriority
}

func Task2(inputFile string) int {
	lines := utils.ReadFileLines(inputFile)

	totalPriority := 0
	// for _, line := range lines {
	// totalPriority += getDuplicatePriorities(first, second)
	// }
	for i := 0; i < len(lines)/3; i++ {
		chunk := lines[i*3 : i*3+3]
		totalPriority += getDuplicatePriorities3(chunk[0], chunk[1], chunk[2])
	}

	return totalPriority
}

func main() {
	fmt.Println("[Task 1]", Task1("input.txt"))
	fmt.Println("[Task 2]", Task2("input.txt"))
}
