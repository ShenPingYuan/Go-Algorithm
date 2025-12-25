package main

import (
	"fmt"
	"strings"
)

func main() {
	var pz = [][]int{{1, 2, 3}, {5, 4, 0}}
	fmt.Print(slidingPuzzle(pz))
}

func slidingPuzzle(board [][]int) int {
	start := ""
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			start += fmt.Sprint(board[i][j])
		}
	}
	end := "123450"
	q := []string{start}
	slidMap := [][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{1, 3, 5},
		{2, 4},
	}
	visited := map[string]bool{start: true}
	step := 0
	for len(q) != 0 {
		lenQ := len(q)
		for lenQ > 0 {
			el := q[0]
			q = q[1:]
			if el == end {
				return step
			}
			var zeroIdx = strings.Index(el, "0")
			for _, v := range slidMap[zeroIdx] {
				elRune := []rune(el)
				elRune[zeroIdx], elRune[v] = elRune[v], elRune[zeroIdx]
				newEl := string(elRune)
				if visited[newEl] {
					continue
				}
				q = append(q, newEl)
				visited[newEl] = true
			}
			lenQ--
		}
		step++
	}
	return -1
}
