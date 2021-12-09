package main

import (
	"advent-of-code/2021/base"
	"fmt"
	"sort"
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

// B basin = group of coordinates
type B = []struct {
	c C
	v int
}

func first() {
	input := base.GetLines()
	matrix := buildMatrix(input)

	sum := 0

	for i, _ := range matrix {
		for j, _ := range matrix[i] {

			adjacents := []C{
				C{i - 1, j},
				C{i, j + 1},
				C{i + 1, j},
				C{i, j - 1},
			}

			isLower := true
			for _, a := range adjacents {
				if isOutOfIndex(a[0], a[1], matrix) {
					continue
				}

				if matrix[i][j] >= matrix[a[0]][a[1]] {
					isLower = false
					break
				}
			}

			if isLower {
				fmt.Println("Found Lower", matrix[i][j])
				sum += 1 + matrix[i][j]
			}

		}
	}

	fmt.Println(sum)
}

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
		C{i - 1, j}, // T
		C{i, j + 1}, // R
		C{i + 1, j}, // B
		C{i, j - 1}, // L
	}
	return adjacents
}

func second() {
	input := base.GetLines()
	matrix := buildMatrix(input)

	sum := 0

	var basins = []B{}
	for i, _ := range matrix {
		for j, _ := range matrix[i] {

			adjacents := getAdjacents(i, j)

			isLower := true
			for _, a := range adjacents {
				if isOutOfIndex(a[0], a[1], matrix) {
					continue
				}

				if matrix[i][j] >= matrix[a[0]][a[1]] {
					isLower = false
					break
				}
			}

			if isLower {
				fmt.Println("Found Lower", matrix[i][j])
				sum += 1 + matrix[i][j]

				newBasin := B{}

				newBasin = append(newBasin, B{{C{i, j}, matrix[i][j]}}...)

				newBasinMap := lookAhead(C{i, j}, matrix)
				for k, v := range newBasinMap {
					newBasin = append(newBasin, B{{k, v}}...)
				}

				fmt.Println("Basin", newBasin)
				basins = append(basins, newBasin)
			}

		}
	}

	lens := []int{}

	for _, b := range basins {
		lens = append(lens, len(b))
	}

	sort.Ints(lens)

	result := lens[len(lens)-3] * lens[len(lens)-2] * lens[len(lens)-1]

	fmt.Println(result)
}

func lookAhead(c C, matrix [][]int) map[C]int {
	i := c[0]
	j := c[1]

	basin := map[C]int{}

	adjacents := getAdjacents(i, j)

	for _, a := range adjacents {
		if isOutOfIndex(a[0], a[1], matrix) {
			continue
		}

		if matrix[i][j] < matrix[a[0]][a[1]] && matrix[a[0]][a[1]] < 9 && matrix[i][j] < 9 {
			n := C{a[0], a[1]}
			basin[n] = matrix[n[0]][n[1]]
			newBasin := lookAhead(a, matrix)

			for k, e := range newBasin {
				basin[k] = e
			}
		}
	}
	return basin
}

func main() {
	//first()
	second()
}
