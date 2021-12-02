package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func puzzle1() {
	file, err := os.Open("./2021/01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	last := -1
	countInc := 0
	countDec := 0

	// read each line
	for scanner.Scan() {

		current, _ := strconv.Atoi(scanner.Text())

		if last == -1 {
			last = current
			continue
		}

		if current > last {
			countInc++
		} else {
			countDec++
		}

		last = current
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inc", countInc)
	// Inc 1387
	fmt.Println("Dec", countDec)
	// Dec 612
}

func moveWindow(array [3]int, new int) [3]int {
	array[0] = array[1]
	array[1] = array[2]
	array[2] = new
	return array
}

func sumWindow(array [3]int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func puzzle2() {
	file, err := os.Open("./2021/01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var window [3]int

	countInc := 0
	countDec := 0
	lastSum := 0

	i := 0
	// read each line
	for scanner.Scan() {

		current, _ := strconv.Atoi(scanner.Text())

		// start summing after 3rd
		// 1 A
		// 2 A
		// 3 A
		if i >= 2 {
			currentSum := sumWindow(window)
			// start counting after the 4th item
			// 1 A
			// 2 A B
			// 3 A B
			// 4   B
			if i >= 3 {
				if currentSum > lastSum {
					countInc++
				} else {
					countDec++
				}
			}

			lastSum = currentSum
		}

		window = moveWindow(window, current)
		i++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inc", countInc)
	// Inc 1362
	fmt.Println("Dec", countDec)
	// Dec 635
}

func main() {
	puzzle1()
	puzzle2()
}
