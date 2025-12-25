package main

func main() {
	//s = "axc"  t = "ahbgdc"
	s := "axc"
	t := "ahbgdc"
	isSubsequence(s, t)
}

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}

	if t == "" {
		return false
	}
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
