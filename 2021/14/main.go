package main

import (
	"advent-of-code/2021/base"
	"fmt"
	"strings"
)

func first() {
	input := base.GetLines()

	template := ""

	rules := map[string]string{}

	for i, line := range input {
		if i == 0 {
			template = line
		}

		if i > 1 {
			parts := strings.Split(line, " -> ")
			rules[parts[0]] = parts[1] + parts[0][1:2]
		}
	}

	for step := 0; step < 10; step++ {
		temp := string(template[0])
		for i := 1; i < len(template); i++ {
			current := string(template[i])
			pair := string(temp[len(temp)-1]) + current

			_, prst := rules[pair]
			if prst {
				temp += rules[pair]
			} else {
				temp += current
			}
		}
		template = temp
	}

	counter := map[string]int{}
	for _, r := range template {
		counter[string(r)]++
	}

	var max, min = 0, 1000000000

	for _, v := range counter {
		if v >= max {
			max = v
		}

		if v <= min {
			min = v
		}
	}

	fmt.Println(max - min)
}

func second() {
	input := base.GetLines()

	template := ""

	rules := map[string]string{}

	for i, line := range input {
		if i == 0 {
			template = line
		}

		if i > 1 {
			parts := strings.Split(line, " -> ")
			rules[parts[0]] = parts[1]
		}
	}
	// init counter with pairs of the template
	counter := map[string]int{}
	for i, _ := range template {
		if i > len(template)-2 {
			break
		}
		counter[template[i:i+2]]++
	}

	for step := 0; step < 40; step++ {
		cnt := map[string]int{}

		for k, _ := range counter {
			// given NN and NN -> C

			// count NC
			cnt[string(k[0])+rules[k]] += counter[k]

			// count CN
			cnt[rules[k]+string(k[1])] += counter[k]
		}

		counter = cnt
	}

	// count letters
	letters := map[string]int{}
	for k, _ := range counter {
		letters[string(k[0])] += counter[k]
	}
	letters[string(template[len(template)-1])]++

	// get min and max
	var max, min = 0, 100000000000000000
	for _, v := range letters {
		if v >= max {
			max = v
		}

		if v <= min {
			min = v
		}
	}

	fmt.Println(max - min)

}

func main() {
	//first()
	second()
}
