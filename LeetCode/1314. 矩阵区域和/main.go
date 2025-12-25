package main

func main() {
	// [1,2,3],
	// [4,5,6],
	// [7,8,9]
}

func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	sums := make([][]int, m+1)
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := range sums {
		sums[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sums[i+1][j+1] = mat[i][j] + sums[i][j+1] + sums[i+1][j] - sums[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			r1, c1, r2, c2 := i-k, j-k, i+k, j+k
			if r1 < 0 {
				r1 = 0
			}
			if c1 < 0 {
				c1 = 0
			}
			if r2 >= m {
				r2 = m - 1
			}
			if c2 >= n {
				c2 = n - 1
			}
			res[i][j] = SumRegion(sums, r1, c1, r2, c2)
		}
	}
	return res
}

func SumRegion(sums [][]int, r1, c1, r2, c2 int) int {
	return sums[r2+1][c2+1] - sums[r2+1][c1] - sums[r1][c2+1] + sums[r1][c1]
}
