package main

func main() {

}

// https://leetcode.cn/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
func merge(nums1 []int, m int, nums2 []int, n int) {
	var totalLength = m + n

	// 将nums2追加到nums1后面
	nums1 = append(nums1, nums2...)

	// 然后将nums1排序
	for i := 0; i < totalLength; i++ {
		for j := 0; j < totalLength-i; j++ {

		}
	}
}
