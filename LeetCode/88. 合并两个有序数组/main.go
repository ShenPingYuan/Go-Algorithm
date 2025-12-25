package main

func main() {

}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	k := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, m+n)
	k := 0
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
		k++
	}

	for i < m {
		nums[k] = nums1[i]
		i++
		k++
	}
	for j < n {
		nums[k] = nums2[j]
		j++
		k++
	}

	for i := 0; i < m+n; i++ {
		nums1[i] = nums[i]
	}
}
