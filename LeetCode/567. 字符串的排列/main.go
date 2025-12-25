package main

func main() {
	//测试
	s1 := "hello"
	s2 := "ooolleoooleh"
	println(checkInclusion1(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}

	need := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	window := make(map[byte]int)
	valid := 0
	left, right := 0, 0

	for right < len(s2) {
		newC := s2[right]
		right++
		if _, has := need[newC]; has {
			window[newC]++
			if window[newC] == need[newC] {
				valid++
				if valid == len(need) {
					return true
				}
			}
		}

		if right-left >= len(s1) {
			deleteC := s2[left]
			if _, has := need[deleteC]; has {
				if window[deleteC] == need[deleteC] {
					valid--
				}
				window[deleteC]--
			}
			left++
		}
	}
	return false
}

func checkInclusion1(s1 string, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)

	need := make(map[byte]int)
	window := make(map[byte]int)

	for _, v := range s1 {
		need[byte(v)]++
	}

	valid := 0
	left, right := 0, 0
	for right < l2 {
		c := s2[right]
		right++
		if v, f := need[c]; f {
			window[c]++
			if window[c] == v {
				valid++
			}
		}

		if l1 == right-left {
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			if v, f := need[d]; f {
				if window[d] == v {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
