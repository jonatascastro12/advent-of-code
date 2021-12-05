package main

import (
	"advent-of-code/2021/base"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type Grid [1000][1000]int

func buildPoint(str string) (int, int) {
	points := strings.Split(str, ",")

	x, _ := strconv.Atoi(points[0])
	y, _ := strconv.Atoi(points[1])

	return x, y
}

func buildLines(input []string) []Line {
	lines := []Line{}

	for _, l := range input {
		points := strings.Split(l, " -> ")

		x1, y1 := buildPoint(points[0])
		x2, y2 := buildPoint(points[1])

		// let's put the points ASC sorted
		if x1*y1 <= x2*y2 {
			lines = append(lines, Line{x1: x1, y1: y1, x2: x2, y2: y2})
		} else {
			lines = append(lines, Line{x1: x2, y1: y2, x2: x1, y2: y1})
		}

	}
	return lines
}

func first() {
	input := base.GetLines()
	lines := buildLines(input)
	grid := Grid{}

	overlaps := map[string]bool{}

	for _, l := range lines {
		if l.x1 == l.x2 || l.y1 == l.y2 {
			for i := l.x1; i <= l.x2; i++ {
				for j := l.y1; j <= l.y2; j++ {
					grid[i][j]++
					if grid[i][j] >= 2 {
						overlaps[strconv.Itoa(i)+strconv.Itoa(j)] = true
					}
				}
			}
		}
	}

	fmt.Println("Result:", len(overlaps))
}

func second() {
	input := base.GetLines()
	lines := buildLines(input)
	grid := Grid{}

	overlapsCounter := map[string]bool{}

	for _, l := range lines {
		dx := int(math.Abs(float64(l.x2 - l.x1)))
		dy := int(math.Abs(float64(l.y2 - l.y1)))

		if dx == dy {
			for k := 0; k <= dx; k++ {
				i, j := 0, 0
				if l.x1 >= l.x2 {
					i = l.x1 - k
				} else {
					i = l.x1 + k
				}

				if l.y1 >= l.y2 {
					j = l.y1 - k
				} else {
					j = l.y1 + k
				}

				grid[i][j]++
				if grid[i][j] >= 2 {
					overlapsCounter[strconv.Itoa(i)+"_"+strconv.Itoa(j)] = true
				}
			}
		} else if dx == 0 || dy == 0 {
			for i := l.x1; i <= l.x2; i++ {
				for j := l.y1; j <= l.y2; j++ {
					grid[i][j]++
					if grid[i][j] >= 2 {
						overlapsCounter[strconv.Itoa(i)+"_"+strconv.Itoa(j)] = true
					}
				}
			}
		}
	}

	fmt.Println("Result:", len(overlapsCounter))

}

func main() {
	first()
	second()
}
