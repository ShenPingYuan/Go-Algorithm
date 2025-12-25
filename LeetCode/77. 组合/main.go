package main

func main() {

}

func combine(n int, k int) [][]int {
	res := [][]int{}

	trace := make([]int, 0, k)
	var backtrack func(index int)
	backtrack = func(index int) {
		if len(trace) == k {
			t := make([]int, k)
			copy(t, trace)
			res = append(res, t)
			return
		}
		for i := index; i <= n; i++ {
			trace = append(trace, i)
			backtrack(i + 1)
			trace = trace[:len(trace)-1]
		}
	}
	backtrack(1)
	return res
}
