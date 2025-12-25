package main

func main() {

}

func rotate(matrix [][]int) {
	transpose(matrix)
	reverse(matrix)
}

func transpose(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func reverse(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		left, right := 0, len(matrix[i])-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}
