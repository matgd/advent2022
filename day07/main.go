package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matgd/advent2022/utils"
)

// Dir struct for storing directory metadata.
type Dir struct {
	name        string
	parent      *Dir
	filesInside []*File
	dirsInside  []*Dir
}

// File struct for storing file metadata.
type File struct {
	name   string
	size   uint64
	parent *Dir
}

func noteUnknownDir(name string, parent *Dir) *Dir {
	return &Dir{name: name, parent: parent, filesInside: []*File{}, dirsInside: []*Dir{}}
}

func noteFile(name string, size uint64, parent *Dir) *File {
	return &File{name: name, size: size, parent: parent}
}

func parseCommands(input []string) *Dir {
	var cwdDirObject *Dir = nil
	cwd := make([]*Dir, 0, 10)

	for _, line := range input {
		if strings.HasPrefix(line, "$ cd") {
			dest := strings.TrimPrefix(line, "$ cd ")
			if dest == ".." {
				cwd = cwd[:len(cwd)-1]
				cwdDirObject = cwdDirObject.parent
			} else {
				if len(cwd) == 0 {
					cwdDirObject = noteUnknownDir(dest, nil)
					cwd = append(cwd, cwdDirObject)
					continue
				}

				var destDir *Dir = nil
				for _, dir := range cwdDirObject.dirsInside {
					if dir.name == dest {
						destDir = dir
						break
					}
				}
				cwdDirObject = destDir
				cwd = append(cwd, cwdDirObject)
			}
		} else if strings.HasPrefix(line, "$ ls") {
			continue
		} else { // Only output is reading from `ls` command
			if strings.HasPrefix(line, "dir") {
				dirName := strings.TrimPrefix(line, "dir ")
				cwdDirObject.dirsInside = append(cwdDirObject.dirsInside, noteUnknownDir(dirName, cwdDirObject))
			} else {
				splitLine := strings.Split(line, " ")
				rawSize, name := splitLine[0], splitLine[1]
				size, _ := strconv.ParseUint(rawSize, 10, 64)
				cwdDirObject.filesInside = append(cwdDirObject.filesInside, noteFile(name, size, cwdDirObject))
			}
		}
	}
	return cwd[0]
}

func getDirsSizes(root *Dir, accumulator *[]uint64) uint64 {
	var totalSize uint64 = 0

	for _, dir := range root.dirsInside {
		totalSize += getDirsSizes(dir, accumulator)
	}
	for _, file := range root.filesInside {
		totalSize += file.size
	}

	newAccumulator := append(*accumulator, totalSize)
	*accumulator = newAccumulator
	return totalSize
}

// Day7 - Day 7, Part 1.
func Day7(inputFile string) uint64 {
	lines := utils.ReadFileLines(inputFile)
	root := parseCommands(lines)

	acc := make([]uint64, 0, 10)
	getDirsSizes(root, &acc)
	var totalSize uint64 = 0
	for _, size := range acc {
		if size < 100_000 {
			totalSize += size
		}
	}
	return totalSize
}

// Day7part2 - Day 7, Part 2.
func Day7part2(inputFile string) uint64 {
	lines := utils.ReadFileLines(inputFile)
	root := parseCommands(lines)

	acc := make([]uint64, 0, 10)
	getDirsSizes(root, &acc)

	rootSize := acc[len(acc)-1] // Last one is the root.
	var freeSpace uint64 = 70_000_000 - rootSize
	var neededSpace uint64 = 30_000_000

	bestCandidate := rootSize
	for _, dirSize := range acc {
		if bestCandidate > dirSize && freeSpace+dirSize >= neededSpace {
			bestCandidate = dirSize
		}
	}
	return bestCandidate
}

func main() {
	fmt.Println("Day 7, Part 1: ", Day7("input.txt"))
	fmt.Println("Day 7, Part 2: ", Day7part2("input.txt"))
}
