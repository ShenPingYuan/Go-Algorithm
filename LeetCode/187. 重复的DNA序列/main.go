package main

func main() {

}

func findRepeatedDnaSequences(s string) []string {
	m := make(map[string]struct{}, 0)
	r := map[string]struct{}{}

	for i := 0; i+10 <= len(s); i++ {
		sub := s[i : i+10]
		if _, f := m[sub]; f {
			r[sub] = struct{}{}
		} else {
			m[sub] = struct{}{}
		}
	}
	result := []string{}
	for k, _ := range r {
		result = append(result, k)
	}
	return result
}

func findRepeatedDnaSequences2(s string) []string {
	m := make(map[string]int)
	res := []string{}

	for i := 0; i+10 <= len(s); i++ {
		sub := s[i : i+10]
		if m[sub] == 1 {
			res = append(res, sub)
		}
		m[sub]++
	}
	return res
}
