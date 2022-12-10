package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/matgd/advent2022/utils"
)

type InstructionType string

const (
	noop InstructionType = "noop"
	addx InstructionType = "addx"
)

type Instruction struct {
	name      InstructionType
	arg       int
	cycleCost int
}

type InstructionQueue struct {
	cycleCounter int
	xVal         int
	historicXVal []int
	instructions []*Instruction
	crt          *CRT
}

type CRT struct {
	litPixel     []bool
	width        int
	height       int
	litPixelMark rune
	dimPixelMark rune
}

func (c CRT) String() string {
	var sb strings.Builder
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			lit := c.litPixel[x+y*c.width]
			if lit {
				sb.WriteRune(c.litPixelMark)
			} else {
				sb.WriteRune(c.dimPixelMark)
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func NewInstructionQueue(maxQueueSize int) *InstructionQueue {
	crt := &CRT{
		width:        40,
		height:       6,
		litPixel:     make([]bool, 40*6),
		litPixelMark: '█',
		dimPixelMark: ' ',
	}

	return &InstructionQueue{
		instructions: make([]*Instruction, maxQueueSize),
		xVal:         1,
		crt:          crt,
	}
}

func (q *InstructionQueue) Enqueue(instruction *Instruction) {
	q.instructions[instruction.cycleCost-1] = instruction
}

func (q *InstructionQueue) NextCycle() {
	q.Draw()
	q.cycleCounter++
	q.historicXVal = append(q.historicXVal, q.xVal)
	q.Execute()
	for i := 0; i < len(q.instructions)-1; i++ {
		q.instructions[i] = q.instructions[i+1]
	}
	q.instructions[len(q.instructions)-1] = nil
}

func (q *InstructionQueue) Draw() {
	row := (q.cycleCounter / q.crt.width)
	cycleOffset := row * q.crt.width
	if math.Abs(float64(q.xVal-(q.cycleCounter-cycleOffset))) >= 2 {
		q.crt.litPixel[q.cycleCounter] = false
		return
	}
	q.crt.litPixel[q.cycleCounter] = true
}

func (q *InstructionQueue) Execute() {
	instruction := q.instructions[0]
	if instruction == nil {
		return
	}
	switch instruction.name {
	case addx:
		q.xVal += instruction.arg
	}
}

func (q *InstructionQueue) Empty() bool {
	for _, instruction := range q.instructions {
		if instruction != nil {
			return false
		}
	}
	return true
}

func ParseInstruction(line string) *Instruction {
	split := strings.Split(line, " ")
	instruction := &Instruction{name: InstructionType(split[0])}
	switch instruction.name {
	case addx:
		arg, _ := strconv.ParseInt(split[1], 10, 64)
		instruction.arg = int(arg)
		instruction.cycleCost = 2
	case noop:
		instruction.cycleCost = 1
	}
	return instruction
}

func Day10(inputFile string) int {
	lines := utils.ReadFileLines(inputFile)

	maxQueueSize := 2
	instructionQueue := NewInstructionQueue(maxQueueSize)

	for _, line := range lines {
		instruction := ParseInstruction(line)
		instructionQueue.Enqueue(instruction)
		for !instructionQueue.Empty() {
			instructionQueue.NextCycle()
		}
	}

	signal_strengths := []int{
		20 * instructionQueue.historicXVal[19],
		60 * instructionQueue.historicXVal[59],
		100 * instructionQueue.historicXVal[99],
		140 * instructionQueue.historicXVal[139],
		180 * instructionQueue.historicXVal[179],
		220 * instructionQueue.historicXVal[219],
	}

	sum := 0
	for _, signal_strength := range signal_strengths {
		sum += signal_strength
	}

	fmt.Println(instructionQueue.crt)

	return sum
}

func main() {
	fmt.Println("Day 10, Part 1:", Day10("input.txt"))
}
