package main

func main() {

}

func characterReplacement(s string, k int) int {
	maxCharCount := 0
	window := make([]int, 26)

	left, right := 0, 0
	res := 0
	for right < len(s) {
		c := s[right]
		window[c-'A']++
		if window[c-'A'] > maxCharCount {
			maxCharCount = window[c-'A']
		}
		right++

		for right-left-maxCharCount > k {
			d := s[left]
			window[d-'A']--
			left++
		}
		// 经过收缩后，此时一定是一个合法的窗口
		if right-left > res {
			res = right - left
		}
	}
	return res
}
