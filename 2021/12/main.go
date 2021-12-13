package main

import (
	"advent-of-code/2021/base"
	"fmt"
	"strings"
	"unicode"
)

type Node struct {
	code string
	rels map[string]*Node
}

func (n *Node) addRel(rel *Node) {
	_, prst := n.rels[rel.code]
	if !prst {
		n.rels[rel.code] = rel
	}
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func buildNodes(input []string) map[string]*Node {
	nodes := map[string]*Node{}

	for _, line := range input {
		parts := strings.Split(line, "-")

		// create nodes if don't exist
		for _, p := range parts {
			_, prst := nodes[p]
			if !prst {
				nodes[p] = &Node{
					code: p,
					rels: map[string]*Node{},
				}
			}
		}

		// add connection
		nodes[parts[0]].addRel(nodes[parts[1]])
		nodes[parts[1]].addRel(nodes[parts[0]])
	}

	return nodes
}

func copyMap(oldMap map[string]bool) map[string]bool {
	newMap := map[string]bool{}
	for k, v := range oldMap {
		newMap[k] = v
	}
	return newMap
}

func copyMapInt(oldMap map[string]int) map[string]int {
	newMap := map[string]int{}
	for k, v := range oldMap {
		newMap[k] = v
	}
	return newMap
}

func walk(root *Node, path []string, visited map[string]bool, paths *[][]string) {
	_, prst := visited[root.code]
	if prst {
		return
	}

	if IsLower(root.code) {
		visited[root.code] = true
	}
	path = append(path, root.code)

	if root.code == "end" {
		*paths = append(*paths, path)
		return
	} else {
		for _, rel := range root.rels {
			walk(rel, path, copyMap(visited), paths)
		}
	}

	return
}

func isThereATwo(visited map[string]int) bool {
	for _, v := range visited {
		if v == 2 {
			return true
		}
	}
	return false
}

func allowedVisitRule(code string, visited map[string]int) bool {
	if IsLower(code) {
		_, prst := visited[code]

		if prst {
			if code == "start" || code == "end" {
				return false
			}

			if visited[code] == 2 {
				return false
			}

			if isThereATwo(visited) && visited[code] == 1 {
				return false
			}
		}
	}
	return true
}

func walk2(root *Node, path []string, visited map[string]int, paths *[][]string) {
	if !allowedVisitRule(root.code, visited) {
		return
	}

	if IsLower(root.code) {
		visited[root.code]++
	}
	path = append(path, root.code)

	if root.code == "end" {
		*paths = append(*paths, path)
		return
	} else {
		for _, rel := range root.rels {
			walk2(rel, path, copyMapInt(visited), paths)
		}
	}

	return
}

func first() {
	input := base.GetLines()
	nodes := buildNodes(input)

	paths := [][]string{}
	walk(nodes["start"], []string{}, map[string]bool{}, &paths)

	fmt.Println(len(paths))
}

func second() {
	input := base.GetLines()
	nodes := buildNodes(input)

	paths := [][]string{}
	walk2(nodes["start"], []string{}, map[string]int{}, &paths)

	fmt.Println(len(paths))
}

func main() {
	first()
	second()
}
