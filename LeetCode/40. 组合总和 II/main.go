package main

func main() {

}

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	trace := []int{}
	sum := 0
	sort(candidates)
	var dfs func(index int)
	dfs = func(index int) {
		if sum > target {
			return
		}
		if sum == target {
			res = append(res, append(make([]int, 0), trace...))
			return
		}
		for i := index; i < len(candidates); i++ {
			if candidates[i] > target {
				continue
			}
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			trace = append(trace, candidates[i])
			sum += candidates[i]
			dfs(i + 1)
			trace = trace[:len(trace)-1]
			sum -= candidates[i]
		}
	}
	dfs(0)
	return res
}

func sort(nums []int) {
	for s := 1; s < len(nums); s++ {
		for i := s; i > 0; i-- {
			if nums[i] > nums[i-1] {
				break
			}
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
}
