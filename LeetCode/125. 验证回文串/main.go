package main

import "unicode"

func main() {

}

func isPalindrome(s string) bool {
	nums := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 65 && c <= 90 {
			nums = append(nums, c)
		}
		if c >= 97 && c <= 122 {
			nums = append(nums, c-32)
		}
		if c >= 48 && c <= 57 {
			nums = append(nums, c)
		}
	}
	slow, fast := 0, len(nums)-1
	for slow < fast {
		if nums[slow] != nums[fast] {
			return false
		}
		slow++
		fast--
	}
	return true
}

func isPalindrome2(s string) bool {
	slow, fast := 0, len(s)-1
	for slow < fast {
		for slow < fast && (!unicode.IsDigit(rune(s[slow])) && !unicode.IsLetter(rune(s[slow]))) {
			slow++
		}
		for slow < fast && (!unicode.IsDigit(rune(s[fast])) && !unicode.IsLetter(rune(s[fast]))) {
			fast--
		}

		if slow < fast && s[slow] != s[fast] && unicode.ToLower(rune(s[slow])) != unicode.ToLower(rune(s[fast])) {
			return false
		}
		slow++
		fast--
	}
	return true
}
