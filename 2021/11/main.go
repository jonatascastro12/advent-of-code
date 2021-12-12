package main

import (
	"advent-of-code/2021/base"
	"fmt"
	"strconv"
	"strings"
)

func buildMatrix(input []string) [][]int {
	var matrix [][]int
	// row
	for i, r := range input {
		row := strings.Split(r, "")

		matrix = append(matrix, []int{})
		// col
		for j, c := range row {
			_ = j
			num, _ := strconv.Atoi(c)
			matrix[i] = append(matrix[i], num)
		}
	}
	return matrix
}

// C coordinate
type C = [2]int

func isOutOfIndex(i int, j int, matrix [][]int) bool {
	if i < 0 ||
		j < 0 ||
		j > len(matrix[0])-1 ||
		i > len(matrix)-1 {
		return true
	}
	return false
}

func getAdjacents(i, j int) []C {
	adjacents := []C{
		C{i - 1, j},     // T
		C{i - 1, j + 1}, // TR
		C{i, j + 1},     // R
		C{i + 1, j + 1}, // BR
		C{i + 1, j},     // B
		C{i + 1, j - 1}, // BL
		C{i, j - 1},     // L
		C{i - 1, j - 1}, // TL
	}
	return adjacents
}

func first() {
	input := base.GetLines()
	matrix := buildMatrix(input)

	count := 0

	for step := 0; step < 100; step++ {
		var flashes []C
		for i, _ := range matrix {
			for j, _ := range matrix[i] {
				matrix[i][j]++

				if matrix[i][j] > 9 && notIn(C{i, j}, flashes) {
					flashes = append(flashes, C{i, j})
					flashes = flash(C{i, j}, matrix, flashes)
				}

			}
		}
		for _, f := range flashes {
			matrix[f[0]][f[1]] = 0
		}
		printMatrix(matrix)
		count += len(flashes)
	}
	fmt.Println(count)
}

func notIn(t C, list []C) bool {
	for _, c := range list {
		if c[0] == t[0] && c[1] == t[1] {
			return false
		}
	}
	return true
}

func second() {
	input := base.GetLines()
	matrix := buildMatrix(input)

	count := 0

	for step := 0; step < 1000; step++ {
		var flashes []C
		for i, _ := range matrix {
			for j, _ := range matrix[i] {
				matrix[i][j]++

				if matrix[i][j] > 9 && notIn(C{i, j}, flashes) {
					flashes = append(flashes, C{i, j})
					flashes = flash(C{i, j}, matrix, flashes)
				}

			}
		}
		for _, f := range flashes {
			matrix[f[0]][f[1]] = 0
		}
		printMatrix(matrix)
		count += len(flashes)

		if checkAllZero(matrix) {
			printMatrix(matrix)
			fmt.Println("Step: ", step+1)
			break
		}
	}
	//fmt.Println(count)
}

func flash(c C, matrix [][]int, flashes []C) []C {
	i := c[0]
	j := c[1]
	_, _ = i, j

	ads := getAdjacents(i, j)

	for _, a := range ads {
		ai, aj := a[0], a[1]
		if isOutOfIndex(ai, aj, matrix) {
			continue
		}
		matrix[ai][aj]++
		if matrix[ai][aj] > 9 && notIn(C{ai, aj}, flashes) {
			flashes = append(flashes, C{ai, aj})
			flashes = flash(C{ai, aj}, matrix, flashes)
			matrix[ai][aj] = 0
		}
	}

	return flashes
}

func checkAllZero(matrix [][]int) bool {
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			if matrix[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func printMatrix(matrix [][]int) {
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			fmt.Printf("%d", matrix[i][j])
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n")
}

func main() {
	//first()
	second()
}
