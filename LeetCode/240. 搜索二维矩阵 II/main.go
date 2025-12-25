package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	x, y := 0, n-1
	for x <= m-1 && y >= 0 {
		v := matrix[x][y]
		if v == target {
			return true
		} else if v > target {
			y--
		} else if v < target {
			x++
		}
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		nums := matrix[i]
		if searchNums(nums, n, target) {
			return true
		}
	}
	return false
}

func searchNums(nums []int, n, target int) bool {
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return false
}
