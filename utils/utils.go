package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type InputLine interface {
	string | int64 | float64
}

func ReadFileLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Couldn't open file '%s' due to error: %s\n", path, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatalf("Couldn't parse contents of file '%s' due to error: %s\n", path, scanner.Err())
	}
	return lines
}

func ReadFileLinesInt64(path string) []int64 {
	lines := ReadFileLines(path)

	linesInt := make([]int64, len(lines))
	for _, line := range lines {
		lineInt, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatalf("Couldn't parse line '%s' due to error: %s\n", line, err)
		}
		linesInt = append(linesInt, lineInt)
	}

	return linesInt
}

func ReadFileLinesFloat64(path string) []float64 {
	lines := ReadFileLines(path)

	linesFloat := make([]float64, len(lines))
	for _, line := range lines {
		lineFloat, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatalf("Couldn't parse line '%s' due to error: %s\n", line, err)
		}
		linesFloat = append(linesFloat, lineFloat)
	}

	return linesFloat
}
