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
	floor      int
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
	return &Cave{cave, width, height, 0, 0}
}

func (c *Cave) FillCave(lineStart [2]int, lineEnd [2]int, replaceFloor bool) {
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

	if replaceFloor {
		if endY+2 > c.floor {
			c.floor = endY + 2
		}
		if startY+2 > c.floor {
			c.floor = startY + 2
		}
	}
}

func (c Cave) PrintFrame() {
	c.PrintExampleArea()
	time.Sleep(10 * time.Millisecond)
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

	if c.space[y+1][x] != AIR && c.space[y+1][x-1] != AIR && c.space[y+1][x+1] != AIR {
		c.restedSand++
		return false
	}

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
	for y := 0; y < 14; y++ {
		for x := 483; x < 513; x++ {
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

func Day14(inputFile string, part uint8) int {
	lines := utils.ReadFileLines(inputFile)
	cave := NewCave()
	for _, line := range lines {
		split := strings.Split(line, " -> ")
		for i := 0; i < len(split)-1; i++ {
			startX, startY := CoordsToInts(split[i])
			endX, endY := CoordsToInts(split[i+1])
			cave.FillCave([2]int{startX, startY}, [2]int{endX, endY}, true)
		}
	}
	if part == 2 {
		cave.FillCave([2]int{0, cave.floor}, [2]int{cave.width - 1, cave.floor}, false)
	}

	for cave.PourSand(500, 0) {
	}

	return cave.restedSand
}

func main() {
	// fmt.Println("Day 14, Part 1:", Day14("input_example.txt", 1))
	fmt.Println("Day 14, Part 1:", Day14("input.txt", 1))
	// fmt.Println("Day 14, Part 1:", Day14("input_example.txt", 2))
	fmt.Println("Day 14, Part 1:", Day14("input.txt", 2))
}
