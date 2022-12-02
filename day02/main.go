package main

import (
	"fmt"
	"strings"

	"github.com/matgd/advent2022/utils"
)

type Game struct {
	ActionsScoring map[string]uint8
	OutcomeScoring map[string]uint8
}

func (g *Game) Score(action1, action2 string) (uint8, uint8) {
	outcome1, outcome2 := ActionOutcome(action1, action2)
	return g.ActionsScoring[action1] + g.OutcomeScoring[outcome1], g.ActionsScoring[action2] + g.OutcomeScoring[outcome2]
}

func NewGame() *Game {
	ActionsScoring := map[string]uint8{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	OutcomeScoring := map[string]uint8{
		"lose": 0,
		"draw": 3,
		"win":  6,
	}

	return &Game{
		ActionsScoring: ActionsScoring,
		OutcomeScoring: OutcomeScoring,
	}
}

func ActionToHumanReadeable(action string) string {
	switch action {
	case "A":
		return "rock"
	case "B":
		return "paper"
	case "C":
		return "scissors"
	}
	return ""
}

func ActionOutcome(action1, action2 string) (string, string) {
	if action1 == action2 {
		return "draw", "draw"
	}

	switch action1 {
	case "A":
		if action2 == "B" {
			return "lose", "win"
		}
		return "win", "lose"
	case "B":
		if action2 == "C" {
			return "lose", "win"
		}
		return "win", "lose"
	case "C":
		if action2 == "A" {
			return "lose", "win"
		}
		return "win", "lose"
	}

	return "", ""
}

func Task1(inputFile string) uint32 {
	lines := utils.ReadFileLines(inputFile)
	game := NewGame()

	var myTotalScore uint32 = 0
	for _, line := range lines {
		// Split line into two actions
		splitLine := strings.Fields(line)
		opponentAction, myAction := splitLine[0], splitLine[1]

		// Unify notation
		myAction = map[string]string{
			"X": "A",
			"Y": "B",
			"Z": "C",
		}[myAction]

		// Score the actions
		_, myScore := game.Score(opponentAction, myAction)
		myTotalScore += uint32(myScore)
	}

	return myTotalScore
}

func task2Decode(inputStategy string) string {
	// A - Rock, B - Paper, C - Scissors
	// X - Lose, Y - Draw, Z - Win
	return map[string]string{
		"A X": "A C",
		"A Y": "A A",
		"A Z": "A B",
		"B X": "B A",
		"B Y": "B B",
		"B Z": "B C",
		"C X": "C B",
		"C Y": "C C",
		"C Z": "C A",
	}[inputStategy]
}

func Task2(inputFile string) uint32 {
	lines := utils.ReadFileLines(inputFile)
	game := NewGame()

	var myTotalScore uint32 = 0
	for _, line := range lines {
		line = task2Decode(line)

		// Split line into two actions
		splitLine := strings.Fields(line)
		opponentAction, myAction := splitLine[0], splitLine[1]

		// Score the actions
		_, myScore := game.Score(opponentAction, myAction)
		myTotalScore += uint32(myScore)
	}

	return myTotalScore
}

func main() {
	fmt.Println("[Task 1] Total score is:", Task1("input.txt"))
	fmt.Println("[Task 2] Total score is:", Task2("input.txt"))
}
