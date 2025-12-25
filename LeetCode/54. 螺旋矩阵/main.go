package main

import "fmt"

func main() {
	//[[1,2,3,4],[5,6,7,8],[9,10,11,12]]
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	left, top, right, bottom := 0, 0, len(matrix[0])-1, len(matrix)-1
	res := []int{}
	for left <= right && top <= bottom {
		if left <= right && top <= bottom {
			for i := left; i <= right; i++ {
				res = append(res, matrix[top][i])
			}
			top++
		}

		if left <= right && top <= bottom {
			for i := top; i <= bottom; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}

		if left <= right && top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right && top <= bottom {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}
