package main

import (
	"advent-of-code/2021/base"
	"fmt"
	"sort"
	"strings"
)

type W = struct {
	c      string
	points int
}

func first() {
	input := base.GetLines()

	cMap := map[string]W{
		")": {"(", 3},
		"]": {"[", 57},
		"}": {"{", 1197},
		">": {"<", 25137},
	}

	errors := findErrors(input, cMap)

	score := 0
	for _, err := range errors {
		score += cMap[err].points
	}

	fmt.Println("Score", score)
}

func findErrors(input []string, cMap map[string]W) []string {
	errors := []string{}

	for _, line := range input {
		errors, _, _ = findError(line, cMap, errors)
	}
	return errors
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

	var completions [][]string

	for _, line := range input {
		_, stack, err := findError(line, cMap, []string{})

		if !err {
			completions = findCompletions(stack, cMap, completions)
		}
	}

	scores := calculateScores(completions, cMap)
	winner := findTheWinner(scores)

	fmt.Println("Score 2:", winner)
}

func findError(line string, cMap map[string]W, errors []string) ([]string, []string, bool) {
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
	return errors, stack, err
}

func findCompletions(stack []string, cMap map[string]W, completions [][]string) [][]string {
	closing := []string{}
	completion := []string{}
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
				completion = append(completion, cMap[pop].c)
			}
		}

	}
	completions = append(completions, completion)
	return completions
}

func findTheWinner(scores []int) int {
	sort.Ints(scores)
	winner := scores[len(scores)/2]
	return winner
}

func calculateScores(completions [][]string, cMap map[string]W) []int {
	scores := []int{}
	for _, c := range completions {
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
