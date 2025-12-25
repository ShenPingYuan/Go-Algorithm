package main

import "fmt"

func main() {
	a, b := 3, 6
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; ; i++ {
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[j][:i]
			}
		}
	}
}
