package main

import (
	"fmt"
	"sync"

	"github.com/matgd/advent2022/utils"
)

type Tree struct {
	height      int8
	visible     bool
	scenicScore uint64
}

type Forest struct {
	trees [][]*Tree
}

func ForestFrom(lines []string) Forest {
	forest := Forest{}
	for _, line := range lines {
		var trees []*Tree
		for _, char := range line {
			trees = append(trees, &Tree{height: int8(char - '0'), visible: false})
		}
		forest.trees = append(forest.trees, trees)
	}
	return forest
}

type Direction uint8

const (
	UpDown Direction = iota
	DownUp
	LeftRight
	RightLeft
)

func (f Forest) VerticalLength() int {
	return len(f.trees)
}

func (f Forest) HorizontalLength() int {
	return len(f.trees[0])
}

func (f *Forest) MarkVisible(direction Direction, wg *sync.WaitGroup) {
	defer wg.Done()
	maxHeight := int8(9)
	highestHeight := int8(-1)

	if direction == UpDown {
		for i := 0; i < f.HorizontalLength(); i++ {
			maxHeight = int8(9)
			highestHeight = int8(-1)
			for j := 0; j < f.VerticalLength(); j++ {
				currentTree := f.trees[j][i]
				if currentTree.height > highestHeight {
					currentTree.visible = true
					highestHeight = currentTree.height
					if highestHeight == maxHeight {
						break
					}
				}
			}
		}
	} else if direction == DownUp {
		for i := 0; i < f.HorizontalLength(); i++ {
			maxHeight = int8(9)
			highestHeight = int8(-1)
			for j := f.VerticalLength() - 1; j >= 0; j-- {
				currentTree := f.trees[j][i]
				if currentTree.height > highestHeight {
					currentTree.visible = true
					highestHeight = currentTree.height
					if highestHeight == maxHeight {
						break
					}
				}
			}
		}
	} else if direction == LeftRight {
		for j := 0; j < f.VerticalLength(); j++ {
			maxHeight = int8(9)
			highestHeight = int8(-1)
			for i := 0; i < f.HorizontalLength(); i++ {
				currentTree := f.trees[j][i]
				if currentTree.height > highestHeight {
					currentTree.visible = true
					highestHeight = currentTree.height
					if highestHeight == maxHeight {
						break
					}
				}
			}
		}
	} else if direction == RightLeft {
		for j := 0; j < f.VerticalLength(); j++ {
			maxHeight = int8(9)
			highestHeight = int8(-1)
			for i := f.HorizontalLength() - 1; i >= 0; i-- {
				currentTree := f.trees[j][i]
				if currentTree.height > highestHeight {
					currentTree.visible = true
					highestHeight = currentTree.height
					if highestHeight == maxHeight {
						break
					}
				}
			}
		}
	}
}

func (f Forest) CountVisible() uint64 {
	count := uint64(0)
	for _, trees := range f.trees {
		for _, tree := range trees {
			if tree.visible {
				count++
			}
		}
	}
	return count
}

func (f Forest) PrintVisibleMatrix() {
	for _, trees := range f.trees {
		for _, tree := range trees {
			if tree.visible {
				fmt.Print("#")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}

func Day8(inputFile string) uint64 {
	lines := utils.ReadFileLines(inputFile)
	forest := ForestFrom(lines)
	directions := []Direction{UpDown, DownUp, LeftRight, RightLeft}

	wg := new(sync.WaitGroup)
	wg.Add(len(directions))

	for _, direction := range directions {
		forest.MarkVisible(direction, wg)
	}

	return forest.CountVisible()
}

func (f Forest) ScenicScore(row, col int) uint64 {
	seenTreesFromAngles := [4]uint64{0, 0, 0, 0}
	treeHeight := f.trees[row][col].height

	// Check up
	for i := row - 1; i >= 0; i-- {
		checkedHeight := f.trees[i][col].height
		if checkedHeight <= treeHeight {
			seenTreesFromAngles[0]++
			if checkedHeight == treeHeight {
				break
			}
		}
		if checkedHeight > treeHeight {
			seenTreesFromAngles[0]++
			break
		}
	}

	// Check down
	for i := row + 1; i < len(f.trees); i++ {
		checkedHeight := f.trees[i][col].height
		if checkedHeight <= treeHeight {
			seenTreesFromAngles[1]++
			if checkedHeight == treeHeight {
				break
			}
		}
		if checkedHeight > treeHeight {
			seenTreesFromAngles[1]++
			break
		}
	}

	// Check left
	for i := col - 1; i >= 0; i-- {
		checkedHeight := f.trees[row][i].height
		if checkedHeight <= treeHeight {
			seenTreesFromAngles[2]++
			if checkedHeight == treeHeight {
				break
			}
		}
		if checkedHeight > treeHeight {
			seenTreesFromAngles[2]++
			break
		}
	}

	// Check right
	for i := col + 1; i < len(f.trees[0]); i++ {
		checkedHeight := f.trees[row][i].height
		if checkedHeight <= treeHeight {
			seenTreesFromAngles[3]++
			if checkedHeight == treeHeight {
				break
			}
		}
		if checkedHeight > treeHeight {
			seenTreesFromAngles[3]++
			break
		}
	}

	totalScore := uint64(1)
	for _, score := range seenTreesFromAngles {
		totalScore *= score
	}
	if totalScore > 4_200_000 {
		fmt.Println("row:", row, "col:", col, "score:", totalScore)
		fmt.Println("seenTreesFromAngles =>", seenTreesFromAngles)
		fmt.Println("totalScore =>", totalScore)

	}
	return totalScore
}

func (f Forest) BestScenicScore() uint64 {
	bestScenicScore := uint64(0)

	for row, trees := range f.trees {
		for col := range trees {
			scenicScore := f.ScenicScore(row, col)
			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}
		}
	}

	return bestScenicScore
}

func Day8_2(inputFile string) uint64 {
	lines := utils.ReadFileLines(inputFile)
	forest := ForestFrom(lines)

	return forest.BestScenicScore()
}

func main() {
	fmt.Println("Day 8, Part 1:", Day8("input.txt"))
	fmt.Println("Day 8, Part 2:", Day8_2("input.txt"))
}
