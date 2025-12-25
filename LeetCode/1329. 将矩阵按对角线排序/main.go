package main

import "fmt"

func main() {
	mat := [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}
	result := diagonalSort(mat)
	fmt.Println(result)
}

func diagonalSort1(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i++ {
		nums := []int{}
		x, y := i, 0
		for x < m && y < n {
			nums = append(nums, mat[x][y])
			x++
			y++
		}
		sort(nums)
		x, y = i, 0
		k := 0
		for x < m && y < n {
			mat[x][y] = nums[k]
			x++
			y++
			k++
		}
	}

	for i := 1; i < n; i++ {
		nums := []int{}
		x, y := 0, i
		for x < m && y < n {
			nums = append(nums, mat[x][y])
			x++
			y++
		}
		sort(nums)
		x, y = 0, i
		k := 0
		for x < m && y < n {
			mat[x][y] = nums[k]
			x++
			y++
			k++
		}
	}
	return mat
}

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	h := map[int][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			h[i-j] = append(h[i-j], mat[i][j])
		}
	}

	for _, v := range h {
		sort(v)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = h[i-j][0]
			h[i-j] = h[i-j][1:]
		}
	}
	return mat
}

func sort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j-1 >= 0; j-- {
			if nums[j] >= nums[j-1] {
				break
			}
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
}
