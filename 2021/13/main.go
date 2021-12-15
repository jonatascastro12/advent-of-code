package main

import (
	"advent-of-code/2021/base"
	"fmt"
	"strconv"
	"strings"
)

type Mark struct {
	x, y int
}

type Inst struct {
	c string // x or y
	n int
}

// i => y   ROW
// j => x   COL

func first() {
	input := base.GetLines()

	maxX, maxY := 0, 0
	var instructions []Inst
	var marks []Mark

	// get min, max, instructions
	marks, instructions, maxX, maxY = extractInfo(input, instructions, marks, maxY, maxX)

	// build matrix with min, max and markds
	matrix := buildMatrix(maxX, maxY, marks)

	// start folding

	for step, inst := range instructions {
		if step > 0 {
			break
		}
		if inst.c == "x" {
			for i := 0; i <= maxY; i++ {
				for j := 0; j < inst.n; j++ {
					matrix[i][inst.n-1-j] += matrix[i][inst.n+1+j]
				}
			}
			maxX = inst.n - 1
		} else {
			for i := 0; i < inst.n; i++ {
				for j := 0; j <= maxX; j++ {
					matrix[inst.n-1-i][j] += matrix[inst.n+1+i][j]
				}
			}
			maxY = inst.n - 1
		}
	}

	count := 0
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			if matrix[i][j] > 0 {
				count++
			}
		}
	}

	printMatrix(matrix, maxX, maxY)

	fmt.Println(count)
}

//  00 10 20
//  01 11 21
//  02 12 22

func buildMatrix(maxX int, maxY int, marks []Mark) [][]int {
	matrix := [][]int{}
	for i := 0; i <= maxY; i++ {
		matrix = append(matrix, []int{})
		for j := 0; j <= maxX; j++ {
			matrix[i] = append(matrix[i], 0)
		}
	}

	// set 1s
	for _, mark := range marks {
		matrix[mark.y][mark.x] = 1
	}

	return matrix
}

func printMatrix(matrix [][]int, maxX int, maxY int) {
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			if matrix[i][j] > 0 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n")
}

func extractInfo(input []string, instructions []Inst, marks []Mark, maxX int, maxY int) ([]Mark, []Inst, int, int) {
	startedInstructions := false
	for _, line := range input {
		if line == "" {
			startedInstructions = true
			continue
		} else if startedInstructions {
			line = strings.ReplaceAll(line, "fold along ", "")
			parts := strings.Split(line, "=")
			n, _ := strconv.Atoi(parts[1])
			c := parts[0]
			instructions = append(instructions, Inst{c, n})

		} else {
			parts := strings.Split(line, ",")

			x, _ := strconv.Atoi(parts[0]) // j
			y, _ := strconv.Atoi(parts[1]) // i

			marks = append(marks, Mark{x, y})

			if x > maxX {
				maxX = x
			}

			if y > maxY {
				maxY = y
			}

		}
	}
	return marks, instructions, maxX, maxY
}

func second() {
	input := base.GetLines()

	maxX, maxY := 0, 0
	var instructions []Inst
	var marks []Mark

	// get min, max, instructions
	marks, instructions, maxX, maxY = extractInfo(input, instructions, marks, maxY, maxX)

	// build matrix with min, max and markds
	matrix := buildMatrix(maxX, maxY, marks)

	// start folding

	for _, inst := range instructions {
		if inst.c == "x" {
			for i := 0; i <= maxY; i++ {
				for j := 0; j < inst.n; j++ {
					matrix[i][inst.n-1-j] += matrix[i][inst.n+1+j]
				}
			}
			maxX = inst.n - 1
		} else {
			for i := 0; i < inst.n; i++ {
				for j := 0; j <= maxX; j++ {
					matrix[inst.n-1-i][j] += matrix[inst.n+1+i][j]
				}
			}
			maxY = inst.n - 1
		}
	}

	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			if matrix[i][j] > 0 {
				matrix[i][j] = 1
			}
		}
	}

	printMatrix(matrix, maxX, maxY)

}

func main() {
	//first()
	second()
}
