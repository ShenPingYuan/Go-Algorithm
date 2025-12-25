package main

func main() {

}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	trace := []int{}
	sum := 0
	var dfs func(start int)
	dfs = func(start int) {
		if sum == target {
			res = append(res, append([]int{}, trace...))
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			sum += candidates[i]
			trace = append(trace, candidates[i])
			dfs(i)
			trace = trace[:len(trace)-1]
			sum -= candidates[i]
		}
	}
	dfs(0)
	return res
}
