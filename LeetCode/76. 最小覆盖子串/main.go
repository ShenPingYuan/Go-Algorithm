package main

func main() {
	// 测试用例
	s := "ADOBECODEBANC"
	t := "ABC"
	result := minWindow2(s, t)
	println("最小覆盖子串:", result)
}

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		c := t[i]
		need[c]++
	}

	window := make(map[byte]int)
	have, needCount := 0, len(need)
	res := [2]int{-1, 1<<31 - 1} //0,011111111111...
	left, right := 0, 0
	for right < len(s) {
		newC := s[right]

		if _, found := need[newC]; found {
			window[newC]++
			if window[newC] == need[newC] {
				have++
			}
		}
		right++

		for left < right && have == needCount {
			if right-left < res[1]-res[0] {
				res[0], res[1] = left, right
			}
			deleteC := s[left]
			if _, found := need[deleteC]; found {
				if window[deleteC] == need[deleteC] {
					have--
				}
				window[deleteC]--
			}
			left++
		}
	}
	if res[0] == -1 {
		return ""
	}
	return s[res[0]:res[1]]
}

func stringContains(s, t string) bool {
	t_map := make(map[rune]int)
	for _, v := range t {
		t_map[v]++
	}
	s_map := make(map[rune]int)
	for _, v := range s {
		s_map[v]++
	}
	for k, v := range t_map {
		if s_map[k] < v {
			return false
		}
	}
	return true
}

func minWindow2(s string, t string) string {
	need := map[byte]int{}
	for _, v := range t {
		need[byte(v)]++
	}

	needCount := len(need)
	valid := 0
	window := map[byte]int{}

	resLeft, resRight := -1, 1<<31-1

	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		right++
		if v, f := need[c]; f {
			window[c]++
			if window[c] == v {
				valid++
			}
		}

		for left < right && valid == needCount {
			if right-left < resRight-resLeft {
				resRight, resLeft = right, left
			}
			d := s[left]
			left++
			if v, f := window[d]; f {
				if v == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if resLeft == -1 {
		return ""
	}

	return s[resLeft:resRight]
}
