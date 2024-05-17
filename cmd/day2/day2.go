package main

import (
	"log"
	"os"
	"strings"

	"github.com/eshom/aoc2022-go/pkg/assert"
)

type Game struct {
	opp string
	you string
}

type OutcomeMap map[Game]int
type ScoreMap map[string]int

const (
	OUTCOME_LOSE = 0
	OUTCOME_DRAW = 3
	OUTCOME_WIN  = 6

	SCORE_ROCK     = 1
	SCORE_PAPER    = 2
	SCORE_SCISSORS = 3
)

func main() {
	// part 1
	outcomes := makeOutcomeMap()
	scores := makeScoreMap()

	var input string = readInput("cmd/day2/data/input.txt")
	var games []Game = parseInput(input)
	var results []int = playGames(games, outcomes, scores)
	var answer int = sum(results)
	log.Println(answer)

	// part 2
	results = playGames2(games)
	answer = sum(results)
	log.Println(answer)
}

func readInput(filename string) string {
	contents, err := os.ReadFile(filename)
	assert.NoError(err)
	return string(contents)
}

func parseInput(input string) []Game {
	var out []Game
	for _, line := range strings.Split(input, "\n") {
		str := strings.Split(line, " ")
		if line == "" {
			break
		}
		game := Game{opp: str[0], you: str[1]}
		out = append(out, game)
	}

	return out
}

func makeOutcomeMap() OutcomeMap {
	var outcome = make(OutcomeMap)
	outcome[Game{opp: "A", you: "X"}] = OUTCOME_DRAW
	outcome[Game{opp: "B", you: "Y"}] = OUTCOME_DRAW
	outcome[Game{opp: "C", you: "Z"}] = OUTCOME_DRAW

	outcome[Game{opp: "A", you: "Y"}] = OUTCOME_WIN
	outcome[Game{opp: "B", you: "Z"}] = OUTCOME_WIN
	outcome[Game{opp: "C", you: "X"}] = OUTCOME_WIN

	outcome[Game{opp: "A", you: "Z"}] = OUTCOME_LOSE
	outcome[Game{opp: "B", you: "X"}] = OUTCOME_LOSE
	outcome[Game{opp: "C", you: "Y"}] = OUTCOME_LOSE

	return outcome
}

func makeScoreMap() ScoreMap {
	var score = make(ScoreMap, 3)
	score["X"] = SCORE_ROCK
	score["Y"] = SCORE_PAPER
	score["Z"] = SCORE_SCISSORS
	return score
}

func playGames(games []Game, outcomes OutcomeMap, scores ScoreMap) []int {
	out := make([]int, 0, 2500)
	for _, game := range games {
		outcome, ok := outcomes[game]
		assert.Assert(ok, "map value not found")
		score, ok := scores[game.you]
		assert.Assert(ok, "map value not found")
		out = append(out, outcome+score)
	}

	return out
}

func playGames2(games []Game) []int {
	out := make([]int, 0, 2500)
	for _, game := range games {
		switch game.you {
		case "X":
			switch game.opp {
			case "A":
				out = append(out, OUTCOME_LOSE+SCORE_SCISSORS)
			case "B":
				out = append(out, OUTCOME_LOSE+SCORE_ROCK)
			case "C":
				out = append(out, OUTCOME_LOSE+SCORE_PAPER)
			}
		case "Y":
			switch game.opp {
			case "A":
				out = append(out, OUTCOME_DRAW+SCORE_ROCK)
			case "B":
				out = append(out, OUTCOME_DRAW+SCORE_PAPER)
			case "C":
				out = append(out, OUTCOME_DRAW+SCORE_SCISSORS)
			}
		case "Z":
			switch game.opp {
			case "A":
				out = append(out, OUTCOME_WIN+SCORE_PAPER)
			case "B":
				out = append(out, OUTCOME_WIN+SCORE_SCISSORS)
			case "C":
				out = append(out, OUTCOME_WIN+SCORE_ROCK)
			}
		}
	}

	return out
}

func sum(vals []int) int {
	var out int = 0
	for _, num := range vals {
		out += num
	}
	return out
}
