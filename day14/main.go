package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/matgd/advent2022/utils"
)

const (
	AIR  rune = '.'
	SAND      = 'o'
	ROCK      = '#'
)

type Cave struct {
	space      [][]rune
	width      int
	height     int
	restedSand int
}

func NewCave() *Cave {
	width := 1000
	height := 1000
	cave := make([][]rune, 0)
	for i := 0; i < height; i++ {
		airRow := make([]rune, width)
		for j := 0; j < width; j++ {
			airRow[j] = AIR
		}
		cave = append(cave, airRow)
	}
	return &Cave{cave, width, height, 0}
}

func (c *Cave) FillCave(lineStart [2]int, lineEnd [2]int) {
	startX := lineStart[0]
	startY := lineStart[1]
	endX := lineEnd[0]
	endY := lineEnd[1]

	if startX == endX {
		if startY > endY {
			startY, endY = endY, startY
		}
		for y := startY; y <= endY; y++ {
			c.space[y][startX] = ROCK
		}
	} else if startY == endY {
		if startX > endX {
			startX, endX = endX, startX
		}
		for x := startX; x <= endX; x++ {
			c.space[startY][x] = ROCK
		}
	}
}

func (c Cave) PrintFrame() {
	c.PrintExampleArea()
	time.Sleep(50 * time.Millisecond)
}

func recoverIndexError() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in f", r)
	}
}

func (c *Cave) PourSand(startX, startY int) bool {
	defer recoverIndexError()
	x, y := startX, startY
	c.space[y][x] = SAND

	// Go down until we hit the another sand or rock
	for y < c.height-1 && y >= 0 && x < c.width-1 && x >= 0 {
		// c.PrintFrame()

		materialBelow := c.space[y+1][x]
		switch materialBelow {
		case AIR:
			y++
			c.space[y-1][x] = AIR
			c.space[y][x] = SAND
		case ROCK, SAND:
			materialLeftBottom := c.space[y+1][x-1]
			materialRightBottom := c.space[y+1][x+1]

			if materialLeftBottom == AIR {
				y++
				x--
				c.space[y-1][x+1] = AIR
				c.space[y][x] = SAND
				// c.PrintFrame()
			} else if materialRightBottom == AIR {
				y++
				x++
				c.space[y-1][x-1] = AIR
				c.space[y][x] = SAND
				// c.PrintFrame()
			} else {
				c.restedSand++
				c.space[y][x] = SAND
				return true
			}
		}
	}

	return false
}

func (c Cave) TotalMaterials(material rune) uint {
	var total uint = 0
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			if c.space[y][x] == material {
				total++
			}
		}
	}
	return total
}

// Cave as a string
func (c *Cave) String() string {
	var sb strings.Builder
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			sb.WriteString(string(c.space[y][x]))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (c *Cave) PrintExampleArea() {
	var sb strings.Builder
	for y := 0; y < 11; y++ {
		for x := 493; x < 503; x++ {
			sb.WriteString(string(c.space[y][x]))
		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}

func CoordsToInts(coords string) (int, int) {
	x, _ := strconv.Atoi(strings.Split(coords, ",")[0])
	y, _ := strconv.Atoi(strings.Split(coords, ",")[1])
	return x, y
}

func Day14(inputFile string) int {
	lines := utils.ReadFileLines(inputFile)
	cave := NewCave()
	for _, line := range lines {
		split := strings.Split(line, " -> ")
		for i := 0; i < len(split)-1; i++ {
			startX, startY := CoordsToInts(split[i])
			endX, endY := CoordsToInts(split[i+1])
			cave.FillCave([2]int{startX, startY}, [2]int{endX, endY})
		}
	}

	for cave.PourSand(500, 0) {
	}

	// fmt.Println(cave.String())

	return cave.restedSand
}

func main() {
	fmt.Println("Day 14, Part 1:", Day14("input_example.txt"))
	fmt.Println("Day 14, Part 1:", Day14("input.txt"))
}
