package main

import (
	"fmt"
	"strconv"

	"github.com/matgd/advent2022/utils"
)

func Task1(inputFile string) uint32 {
	lines := utils.ReadFileLines(inputFile)

	var max uint32 = 0
	var currentTotalCalories uint32 = 0
	for _, line := range lines {
		if line == "" {
			if currentTotalCalories > max {
				max = currentTotalCalories
			}
			currentTotalCalories = 0
		} else {
			calories, _ := strconv.ParseInt(line, 10, 32)
			currentTotalCalories += uint32(calories)
		}
	}

	return max
}

func Task2(inputFile string) uint32 {
	lines := utils.ReadFileLines(inputFile)
	lines = append(lines, "") // Blank end line

	var top3 []uint32 = []uint32{0, 0, 0}

	var currentTotalCalories uint32 = 0
	for _, line := range lines {
		if line == "" {
			if currentTotalCalories > top3[0] {
				top3[2] = top3[1]
				top3[1] = top3[0]
				top3[0] = currentTotalCalories
			} else if currentTotalCalories > top3[1] {
				top3[2] = top3[1]
				top3[1] = currentTotalCalories
			} else if currentTotalCalories > top3[2] {
				top3[2] = currentTotalCalories
			}
			currentTotalCalories = 0
		} else {
			calories, _ := strconv.ParseInt(line, 10, 32)
			currentTotalCalories += uint32(calories)
		}
	}
	return top3[0] + top3[1] + top3[2]
}

func main() {
	fmt.Println("Task 1: ", Task1("input.txt"))
	fmt.Println("Task 2: ", Task2("input.txt"))
}
