package main

import (
	"advent-of-code/2021/base"
	"fmt"
	"strconv"
	"strings"
)

func splitByWidthTrimmed(str string, size int) []string {
	strLength := len(str)
	var splited []string
	var stop int
	for i := 0; i < strLength; i += size {
		stop = i + size
		if stop > strLength {
			stop = strLength
		}
		splited = append(splited, strings.TrimSpace(str[i:stop]))
	}
	return splited
}

type Num struct {
	n      int
	marked bool
}
type Board = [5][5]Num

type Winner struct {
	b            Board
	bIndex       int
	unmarkedCalc int
	calc         int
}

func populateBoards(lines []string) []Board {
	boards := []Board{}

	for i := 2; i < len(lines); i += 6 {
		k := 0
		b := Board{}
		for j := i; j < i+5; j++ {
			cur := splitByWidthTrimmed(lines[j], 3)
			for p, c := range cur {
				n, _ := strconv.Atoi(c)
				b[k][p] = Num{n: n}
			}
			k++
		}
		boards = append(boards, b)
	}
	return boards
}

func checkCurPos(b *Board, posI int, posJ int) bool {
	countI, countJ := 0, 0
	for i := 0; i < 5; i++ {
		if b[posI][i].marked == true {
			countI++
		}
		if b[i][posJ].marked == true {
			countJ++
		}
	}
	if countI == 5 || countJ == 5 {
		return true
	}
	return false
}

func markBoardAndCheck(b *Board, num int) (*Board, bool) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j].n == num {
				b[i][j].marked = true
				if checkCurPos(b, i, j) {
					return b, true
				}
			}
		}
	}
	return b, false
}

func calculateUnmarked(b *Board) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b[i][j].marked {
				sum += b[i][j].n
			}
		}
	}
	return sum
}

func markBoardsAndCalculate(boards []Board, num int) []Winner {
	winners := []Winner{}
	for i := range boards {
		_, win := markBoardAndCheck(&boards[i], num)
		if win {
			unmarkedCalc := calculateUnmarked(&boards[i])
			winners = append(winners, Winner{b: boards[i], bIndex: i, unmarkedCalc: unmarkedCalc, calc: unmarkedCalc * num})
		}

	}
	return winners
}

func puzzle() {

	lines := base.GetLines()

	nums := strings.Split(lines[0], ",")
	_ = nums

	boards := populateBoards(lines)

	winnersStack := []Winner{}

	for _, n := range nums {
		num, _ := strconv.Atoi(n)
		winners := markBoardsAndCalculate(boards, num)

		// if element of winners not in winnersStack, insert to winnersStack
		for _, win := range winners {
			found := false
			for _, winS := range winnersStack {
				if winS.bIndex == win.bIndex {
					found = true
					break
				}
			}
			if !found {
				winnersStack = append(winnersStack, win)
			}
		}
	}

	first := winnersStack[0]
	last := winnersStack[len(winnersStack)-1]

	fmt.Println(first, first.calc)
	fmt.Println(last, last.calc)

}

func main() {
	puzzle()
}
