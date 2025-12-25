package main

func main() {

}

func generateMatrix(n int) [][]int {
	matrix := [][]int{}
	for i := 0; i < n; i++ {
		nums := make([]int, n)
		matrix = append(matrix, nums)
	}
	val := 0
	left, top, right, bottom := 0, 0, n-1, n-1
	for left <= right && top <= bottom {
		if left <= right && top <= bottom {
			for i := left; i <= right; i++ {
				val++
				matrix[top][i] = val
			}
			top++
		}

		if left <= right && top <= bottom {
			for i := top; i <= bottom; i++ {
				val++
				matrix[i][right] = val
			}
			right--
		}

		if left <= right && top <= bottom {
			for i := right; i >= left; i-- {
				val++
				matrix[bottom][i] = val
			}
			bottom--
		}

		if left <= right && top <= bottom {
			for i := bottom; i >= top; i-- {
				val++
				matrix[i][left] = val
			}
			left++
		}
	}
	return matrix
}
