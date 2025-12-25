package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	l, r := 0, m*n-1

	for l <= r {
		mid := l + (r-l)/2
		v := getMatrix(matrix, mid, n)
		if v == target {
			return true
		} else if v < target {
			l = mid + 1
		} else if v > target {
			r = mid - 1
		}
	}
	return false
}

func getMatrix(matrix [][]int, index, c int) int {
	row := index / c
	col := index % c
	return matrix[row][col]
}
