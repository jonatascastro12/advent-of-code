package main

import (
	"advent-of-code/2021/base"
	"fmt"
	"strconv"
)

func first() {
	lines := base.GetLines()

	bitsCounter := [12]int{}
	linesCounter := 0
	for _, line := range lines {
		for i, bit := range line {
			if bit == '1' {
				bitsCounter[i]++
			}
		}
		linesCounter++
	}
	onesStr := ""
	zerosStr := ""
	for _, c := range bitsCounter {
		if c > linesCounter/2 {
			onesStr += "1"
			zerosStr += "0"
		} else {
			onesStr += "0"
			zerosStr += "1"

		}
	}

	ones, _ := strconv.ParseInt(onesStr, 2, 64)
	zeros, _ := strconv.ParseInt(zerosStr, 2, 64)

	fmt.Println("Result: ", ones*zeros)
}

func filterLines(lines []string, bit uint8, pos int) []string {
	result := []string{}
	for _, line := range lines {
		if line[pos] == bit {
			result = append(result, line)
		}
	}
	return result
}

func second() {
	countRates := func(rateType string, lines []string) int64 {
		for i := 0; i < 12; i++ {
			cOnes := 0
			cZeros := 0

			for _, line := range lines {
				if line[i] == '1' {
					cOnes++
				} else {
					cZeros++
				}
			}

			if cOnes >= cZeros && rateType == "o2" || (cOnes < cZeros && rateType == "co2") {
				lines = filterLines(lines, '1', i)
			} else if cOnes < cZeros && rateType == "o2" || (cOnes >= cZeros && rateType == "co2") {
				lines = filterLines(lines, '0', i)
			}

			if len(lines) == 1 {
				break
			}
		}

		num, _ := strconv.ParseInt(lines[0], 2, 64)
		return num
	}

	linesO2 := base.GetLines()
	linesCO2 := base.GetLines()

	o2 := countRates("o2", linesO2)
	co2 := countRates("co2", linesCO2)

	fmt.Println("Result o2xco2", o2*co2)
	//6940518
}

func main() {
	first()
	second()
}
