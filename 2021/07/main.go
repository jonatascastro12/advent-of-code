package main

import (
	"advent-of-code/2021/base"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// brute force way
func first() {
	input := base.GetLines()

	max := 0
	min := 99999
	positions := []int{}

	// parse to in and find min and max
	for _, s := range strings.Split(input[0], ",") {
		num, _ := strconv.Atoi(s)
		positions = append(positions, num)
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}

	minFuel := 999999999
	for i := min; i <= max; i++ {
		totalFuel := 0

		for _, pos := range positions {
			totalFuel += int(math.Abs(float64(pos - i)))
		}

		if totalFuel <= minFuel {
			minFuel = totalFuel
		}
	}

	fmt.Println(minFuel)
}

// smarter way
func second() {
	input := base.GetLines()

	max := 0
	min := 99999
	positions := []int{}

	// parse to in and find min and max
	for _, s := range strings.Split(input[0], ",") {
		num, _ := strconv.Atoi(s)
		positions = append(positions, num)
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}

	minFuel := 999999999
	for i := min; i <= max; i++ {
		totalFuel := 0

		for _, pos := range positions {
			dif := int(math.Abs(float64(pos - i)))
			fuel := 0
			for j := 1; j <= dif; j++ {
				fuel += j
			}
			totalFuel += fuel
		}

		if totalFuel <= minFuel {
			minFuel = totalFuel
		}
	}

	fmt.Println(minFuel)
}

func main() {
	first()
	second()
}
