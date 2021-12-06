package main

import (
	"advent-of-code/2021/base"
	"fmt"
	"strconv"
	"strings"
)

func initFishes(input string) []int8 {
	var fishes []int8

	firstFishes := strings.Split(input, ",")

	fishes = []int8{}

	for _, f := range firstFishes {
		n, _ := strconv.Atoi(f)
		fishes = append(fishes, int8(n))
	}

	return fishes
}

// brute force way
func first() {
	input := base.GetLines()

	fishes := initFishes(input[0])
	fmt.Println(fishes)

	day := 0
	for {
		if day == 80 {
			break
		}
		// last index
		nextFishes := []int8{}
		countToInsert := 0
		for i, _ := range fishes {
			if fishes[i] > 0 {
				nextFishes = append(nextFishes, fishes[i]-1)
			} else {
				nextFishes = append(nextFishes, 6)
				countToInsert++
			}
		}

		for i := 0; i < countToInsert; i++ {
			nextFishes = append(nextFishes, 8)
		}

		fishes = nextFishes
		day++
	}

	fmt.Println(len(fishes))
}

// smarter way
func second() {
	input := base.GetLines()

	fishes := initFishes(input[0])
	fmt.Println(fishes)

	counter := make([]int, 9)

	// initial counter
	for _, f := range fishes {
		counter[f] += 1
	}

	for d := 0; d < 256; d++ {
		deadFishes := counter[0]

		// count day for living fishes - shift left
		for i := 0; i < 8; i++ {
			counter[i] = counter[i+1]
		}

		// add new borns
		counter[8] = deadFishes

		// reset
		counter[6] += deadFishes
	}

	sum := 0
	for _, num := range counter {
		sum += num
	}
	fmt.Println(sum)
}

func main() {
	first()
	second()
}
