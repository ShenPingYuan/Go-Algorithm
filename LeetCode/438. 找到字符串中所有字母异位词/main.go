package main

func main() {

}

func findAnagrams(s string, p string) []int {
	if len(s) == 0 || len(p) == 0 {
		return []int{}
	}

	res := []int{}
	need := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	window := make(map[byte]int)
	valid := 0
	len_p := len(p)
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		if right-left == len_p {
			if valid == len(need) {
				res = append(res, left)
			}
			d := s[left]
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
			left++
		}
	}
	return res
}

func findAnagrams1(s string, p string) []int {
	l_p := len(p)
	l_s := len(s)

	need, window := make(map[byte]int), make(map[byte]int)

	for _, v := range p {
		need[byte(v)]++
	}
	res := []int{}
	valid := 0
	left, right := 0, 0
	for right < l_s {
		c := s[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		if right-left == l_p {
			if valid == len(need) {
				res = append(res, left)
			}
			d := s[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}
