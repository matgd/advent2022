package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matgd/advent2022/utils"
)

type Section struct{ from, to int }

func getSectionFromDashString(dashString string) Section {
	split := strings.Split(dashString, "-")
	from, to := split[0], split[1]
	fromInt, _ := strconv.Atoi(from)
	toInt, _ := strconv.Atoi(to)
	return Section{from: fromInt, to: toInt}
}

func getSectionsFromFile(filePath string) [][2]Section {
	lines := utils.ReadFileLines(filePath)
	sections := [][2]Section{}

	for _, line := range lines {
		splitLine := strings.Split(line, ",")
		rawSectionA, rawSectionB := splitLine[0], splitLine[1]
		sectionA := getSectionFromDashString(rawSectionA)
		sectionB := getSectionFromDashString(rawSectionB)

		sections = append(sections, [2]Section{sectionA, sectionB})
	}
	return sections
}

func Task1(filePath string) int {
	sections := getSectionsFromFile(filePath)
	totalOverlapping := 0

	for _, section := range sections {
		if section[0].from <= section[1].from && section[0].to >= section[1].to {
			totalOverlapping++
		} else if section[1].from <= section[0].from && section[1].to >= section[0].to {
			totalOverlapping++
		}
	}

	return totalOverlapping
}

func Task2(filePath string) int {
	sections := getSectionsFromFile(filePath)
	totalOverlapping := 0

	for _, section := range sections {
		if section[0].from <= section[1].from && section[0].to >= section[1].from {
			totalOverlapping++
		} else if section[1].from <= section[0].from && section[1].to >= section[0].from {
			totalOverlapping++
		}
	}

	return totalOverlapping
}

func main() {
	fmt.Println("[Part 1]", Task1("input.txt"))
	fmt.Println("[Part 2]", Task2("input.txt"))
}
