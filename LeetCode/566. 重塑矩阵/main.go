package main

func main() {

}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat) == 0 {
		return mat
	}

	m := len(mat)
	n := len(mat[0])
	if r*c != m*n {
		return mat
	}

	if m == r {
		return mat
	}

	shaped := make([][]int, r)
	for i, _ := range shaped {
		shaped[i] = make([]int, c)
	}

	x, y := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			shaped[x][y] = mat[i][j]
			y++
			if y == c {
				x++
				y = 0
			}
		}
	}
	return shaped
}
