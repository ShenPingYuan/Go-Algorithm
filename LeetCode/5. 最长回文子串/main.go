package main

func main() {
	// 测试
	s := "b"
	result := longestPalindrome(s)
	println(result) // 输出: "bab" 或 "aba"
}

// func longestPalindrome(s string) string {
// 	res := ""
// 	for i := 0; i < len(s); i++ {
// 		s1 := palindrome(s, i, i)
// 		if len(s1) > len(res) {
// 			res = s1
// 		}

// 		s2 := palindrome(s, i, i+1)
// 		if len(s2) > len(res) {
// 			res = s2
// 		}
// 	}
// 	return res
// }

// func palindrome(s string, i, j int) string {
// 	for i >= 0 && j < len(s) && s[i] == s[j] {
// 		i--
// 		j++
// 	}
// 	return s[i+1 : j]
// }

func longestPalindrome(s string) string {
	maxLeft, maxRight := 0, 0
	pos := 0
	for pos < len(s) {
		left, right := findPalindrome(s, pos, pos)
		if right-left > maxRight-maxLeft {
			maxLeft, maxRight = left, right
		}
		left, right = findPalindrome(s, pos, pos+1)
		if right-left > maxRight-maxLeft {
			maxLeft, maxRight = left, right
		}
		pos++
	}
	return s[maxLeft : maxRight+1]
}

func findPalindrome(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
