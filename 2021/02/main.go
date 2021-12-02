package main

import (
	"advent-of-code/2021/base"
	"fmt"
	"strconv"
	"strings"
)

func first() {
	lines := base.GetLines()

	depth := 0
	x := 0

	for _, line := range lines {
		s := strings.Split(line, " ")

		num, _ := strconv.Atoi(s[1])

		switch s[0] {
		case "up":
			depth -= num
			break
		case "down":
			depth += num
			break
		case "forward":
			x += num
			break
		}
	}

	fmt.Println("Depth:", depth, "X:", x, "Result:", depth*x)
}

func second() {
	lines := base.GetLines()

	depth := 0
	x := 0
	aim := 0

	for _, line := range lines {
		s := strings.Split(line, " ")

		num, _ := strconv.Atoi(s[1])

		switch s[0] {
		case "up":
			aim -= num
			break
		case "down":
			aim += num
			break
		case "forward":
			x += num
			depth += aim * num
			break
		}
	}

	fmt.Println("Depth:", depth, "X:", x, "Result:", depth*x)
}

func main() {
	first()
	second()
}
