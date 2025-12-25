package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	res := 0
	window := map[byte]int{}
	left, right := 0, 0
	lens := len(s)
	for right < lens {
		c := s[right]
		right++

		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}

func lengthOfLongestSubstring1(s string) int {
	window := make(map[rune]int)
	repeated := 0

	left, right := 0, 0
	max := 0
	for right < len(s) {
		c := rune(s[right])
		right++
		if window[c] == 1 {
			repeated++
		}
		window[c]++

		for left < right && repeated > 0 {
			d := rune(s[left])
			left++
			if window[d] == 2 {
				repeated--
			}
			window[d]--
		}
		if right-left > max {
			max = right - left
		}
	}
	return max
}
