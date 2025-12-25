package main

func main() {

}
func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return true
	}
	window := make(map[int]struct{}, 0)

	left, right := 0, 0
	for right < len(nums) {
		new := nums[right]
		if _, f := window[new]; f {
			return true
		}
		window[new] = struct{}{}
		right++

		if right-left > k {
			d := nums[left]
			delete(window, d)
			left++
		}
	}
	return false
}
