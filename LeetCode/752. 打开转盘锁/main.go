package main

import "fmt"

func main() {
	deadends := []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	target := "8888"
	fmt.Println(openLock(deadends, target))
}

func openLock(deadends []string, target string) int {
	start := "0000"
	q := []string{start}
	deadMap := make(map[string]struct{})
	for _, v := range deadends {
		deadMap[v] = struct{}{}
	}
	visited := map[string]struct{}{start: struct{}{}}
	step := 0
	for len(q) != 0 {
		lenQ := len(q)
		for lenQ > 0 {
			e := q[0]
			q = q[1:]
			if e == target {
				return step
			}
			subs := children(e)
			for _, v := range subs {
				if _, f := visited[v]; f {
					continue
				}
				if _, f := deadMap[e]; f {
					continue
				}
				q = append(q, v)
				visited[v] = struct{}{}
			}
			lenQ--
		}
		step++
	}
	return -1
}

func children(numStr string) []string {
	res := make([]string, 0, 8)
	for i := 0; i < 4; i++ {
		res = append(res, plusOne(numStr, i))
		res = append(res, minusOne(numStr, i))
	}
	return res
}

func plusOne(str string, i int) string {
	bytes := []byte(str)
	if bytes[i] == '9' {
		bytes[i] = '0'
	} else {
		bytes[i] += 1
	}
	return string(bytes)
}

func minusOne(str string, i int) string {
	bytes := []byte(str)
	if bytes[i] == '0' {
		bytes[i] = '9'
	} else {
		bytes[i] -= 1
	}
	return string(bytes)
}
