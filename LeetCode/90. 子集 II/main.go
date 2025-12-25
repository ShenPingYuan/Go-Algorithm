package main

func main() {

}

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}

	trace := []int{}
	sort(nums)
	var dfs func(index int)
	dfs = func(index int) {
		t := make([]int, len(trace))
		copy(t, trace)
		res = append(res, t)
		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			trace = append(trace, nums[i])
			dfs(i + 1)
			trace = trace[:len(trace)-1]
		}
	}
	dfs(0)
	return res
}

func sort(nums []int) {
	for sortedIndex := 1; sortedIndex < len(nums); sortedIndex++ {
		for i := sortedIndex; i > 0; i-- {
			if nums[i] > nums[i-1] {
				break
			}
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
}
