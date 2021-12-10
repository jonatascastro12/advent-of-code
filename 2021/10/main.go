package main

import (
	"advent-of-code/2021/base"
	"fmt"
	"sort"
	"strings"
)

func first() {
	input := base.GetLines()

	chunkMap := map[string]struct {
		open   string
		points int
	}{
		")": {"(", 3},
		"]": {"[", 57},
		"}": {"{", 1197},
		">": {"<", 25137},
	}

	errors := []string{}

	for _, line := range input {
		stack := []string{}

		for _, r := range line {
			s := string(r)

			if strings.ContainsAny(s, "([{<") {
				stack = append(stack, s)
			} else {
				pop := stack[len(stack)-1]

				if pop == chunkMap[s].open {
					stack = stack[:len(stack)-1]
				} else {
					errors = append(errors, s)
					break
				}
			}
		}
	}

	score := 0
	for _, err := range errors {
		score += chunkMap[err].points
	}

	fmt.Println("Score", score)
}

func second() {
	input := base.GetLines()

	cMap := map[string]struct {
		c      string
		points int
	}{
		")": {"(", 1},
		"]": {"[", 2},
		"}": {"{", 3},
		">": {"<", 4},

		"(": {")", 1},
		"[": {"]", 2},
		"{": {"}", 3},
		"<": {">", 4},
	}

	errors := []string{}
	completitions := [][]string{}

	for _, line := range input {
		stack := []string{}
		err := false

		for _, r := range line {
			s := string(r)

			if strings.ContainsAny(s, "([{<") {
				stack = append(stack, s)
			} else {
				pop := stack[len(stack)-1]

				if pop == cMap[s].c {
					stack = stack[:len(stack)-1]
				} else {
					err = true
					errors = append(errors, s)
					break
				}
			}
		}

		if !err {
			closing := []string{}
			completition := []string{}
			for {
				if len(stack) == 0 {
					break
				}

				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if strings.ContainsAny(pop, ")]}>") {
					closing = append(closing, pop)
				} else {
					if len(closing) > 1 {
						last_closing := closing[len(closing)-1]
						if cMap[pop].c != last_closing {
							closing = closing[:len(closing)-1]
						}
					} else {
						completition = append(completition, cMap[pop].c)
					}
				}

			}
			completitions = append(completitions, completition)
		}
	}

	// calculate scores
	scores := calculateScores(completitions, cMap)

	// look for the winner
	winner := findTheWinner(scores)

	fmt.Println("Score 2:", winner)
}

func findTheWinner(scores []int) int {
	sort.Ints(scores)
	winner := scores[len(scores)/2]
	return winner
}

func calculateScores(completitions [][]string, cMap map[string]struct {
	c      string
	points int
}) []int {
	scores := []int{}
	for _, c := range completitions {
		score := 0
		for _, s := range c {
			score = score*5 + cMap[s].points
		}
		scores = append(scores, score)
	}
	return scores
}

func main() {
	first()
	second()
}
