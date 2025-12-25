package main

func main() {
	//304. 二维区域和检索 - 矩阵不可变
	// ["NumMatrix","sumRegion","sumRegion","sumRegion"]
	//[[[
	// [3,0,1,4,2],
	// [5,6,3,2,1],
	// [1,2,0,1,5],
	// [4,1,0,1,7],
	// [1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]

	// [[3  3  4 8  10]
	//  [8 14 18 24 27]
	//  [9 17 21 28 36]
	// [13 22 26 34 49]
	// [14 23 30 38 58]]
	// matrix := [][]int{
	// 	{3, 0, 1, 4, 2},
	// 	{5, 6, 3, 2, 1},
	// 	{1, 2, 0, 1, 5},
	// 	{4, 1, 0, 1, 7},
	// 	{1, 0, 3, 0, 5},
	// }
	// obj := Constructor(matrix)
	// param_1 := obj.SumRegion(2, 1, 4, 3)
	// println(param_1)

	// param_2 := obj.SumRegion(1, 1, 2, 2)
	// println(param_2)

	// param_3 := obj.SumRegion(1, 2, 2, 4)
	// println(param_3)
}

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	// 多开一行一列，sums[i+1][j+1] 表示 (0,0) 到 (i,j) 的区域和
	sums := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		sums[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sums[i+1][j+1] = matrix[i][j] + sums[i][j+1] + sums[i+1][j] - sums[i][j]
		}
	}

	return NumMatrix{
		sums: sums,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sums := this.sums
	return sums[row2+1][col2+1] - sums[row1][col2+1] - sums[row2+1][col1] + sums[row1][col1]
}
