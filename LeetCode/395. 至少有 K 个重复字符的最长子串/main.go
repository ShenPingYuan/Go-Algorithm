package main

func main() {
	//s = "bbaaacbd" k = 3
	s := "bbaaacbd"
	k := 3
	longestSubstring(s, k)
}

func longestSubstring(s string, k int) int {
	if len(s) == 0 || k > len(s) {
		return 0
	}
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i, v := range s {
		if m[byte(v)] < k {
			left := longestSubstring(s[:i], k)
			right := longestSubstring(s[i+1:], k)
			return max(left, right)
		}
	}
	return len(s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestSubstring1(s string, k int) int {
	res := 0
	for i := 1; i <= 26; i++ {
		if i > len(s) {
			break
		}
		count := nCountlongestSubstring(s, k, i)
		if count > res {
			res = count
		}
	}
	return res
}

func nCountlongestSubstring(s string, k, n int) int {
	windowCharCount := 0
	window := make(map[byte]int)
	validCount := 0
	left, right := 0, 0
	res := 0
	for right < len(s) {
		new := s[right]
		if _, f := window[new]; !f {
			windowCharCount++
		}
		window[new]++
		if window[new] == k {
			validCount++
		}
		right++
		for left < right && windowCharCount > n {
			d := s[left]
			if window[d] == k {
				validCount--
			}
			window[d]--
			if window[d] == 0 {
				delete(window, d)
				windowCharCount--
			}
			left++
		}
		if validCount == n && right-left > res {
			res = right - left
		}
	}
	return res
}
