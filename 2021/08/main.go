package main

import (
	"advent-of-code/2021/base"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func first() {
	input := base.GetLines()

	count := 0

	for _, l := range input {
		parts := strings.Split(l, " | ")

		outParts := strings.Split(parts[1], " ")

		for _, p := range outParts {
			switch len(p) {
			case 2, 4, 3, 7: // 1 digit
				count++
			}
		}
	}

	fmt.Println(count)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i int, j int) bool { return r[i] < r[j] })
	return string(r)
}

func subStringFromAny(p string, s string) string {
	str := p
	for _, s := range s {
		str = strings.ReplaceAll(str, string(s), "")
	}
	return str
}

func containsAll(str string, sub string) bool {
	for _, s := range sub {
		if !strings.Contains(str, string(s)) {
			return false
		}
	}
	return true
}

func second() {
	input := base.GetLines()

	sum := 0

	for _, l := range input {
		parts := strings.Split(l, " | ")

		sigParts := strings.Split(parts[0], " ")

		sort.Slice(sigParts, func(i, j int) bool {
			return len(sigParts[i]) < len(sigParts[j])
		})

		m := map[string]string{}
		patterns := map[string]string{}

		//Len: 2      3       4            5                6                7
		//
		//          _                _   _    _       _     _     _          _
		//    |      |     |_|       _|  _|  |_      |_    |_|   | |        |_|
		//    |      |       |      |_   _|   _|     |_|    _|   |_|        |_|
		//
		//
		// Some inferences
		//   _               _      ||
		//  |_|  - |_|  =           ||      |_|  -   |   =    |_
		//  |_|      |      |_      ||        |      |
		//

		for _, p := range sigParts {
			p = sortString(p)

			switch len(p) {
			case 2:
				m[p] = "1"
				patterns["1"] = p
			case 3:
				m[p] = "7"
				patterns["7"] = p
			case 4:
				m[p] = "4"
				patterns["4"] = p
				// (4-1)
				patterns["4-1"] = subStringFromAny(p, patterns["1"])
			case 7:
				m[p] = "8"
				patterns["8-4"] = subStringFromAny(p, patterns["4"])
			}
		}

		for _, p := range sigParts {
			p = sortString(p)

			switch len(p) {
			case 5:
				if containsAll(p, patterns["1"]) && patterns["6"] == "" {
					// (1 in p)
					m[p] = "3"
					patterns["3"] = p
				} else if containsAll(p, patterns["8-4"]) && patterns["2"] == "" {
					// (8-4 in p)
					m[p] = "2"
					patterns["2"] = p
				} else if patterns["5"] == "" {
					// (4-1 in p)
					m[p] = "5"
					patterns["5"] = p
				}
			case 6:
				if containsAll(p, patterns["4"]) && patterns["9"] == "" {
					// (7 in p)
					m[p] = "9"
					patterns["9"] = p
				} else if containsAll(p, patterns["7"]) && patterns["0"] == "" {
					m[p] = "0"
					patterns["0"] = p
				} else if patterns["6"] == "" {
					m[p] = "6"
					patterns["6"] = p
				}
			}
		}

		if len(m) != 10 {
			fmt.Println("PROBLEM")
			break
		}

		outParts := strings.Split(parts[1], " ")

		value := ""
		for _, p := range outParts {
			p = sortString(p)
			value += m[p]
		}

		num, _ := strconv.Atoi(value)

		//fmt.Println(outParts, value, num)

		sum += num
	}

	fmt.Println(sum)

}

func main() {
	first()
	second()
}
