package utils

import (
	"bufio"
	"log"
	"os"
)

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
