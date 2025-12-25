package main

func main() {

}
func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	nums := make([]int, m*n)
	d := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			nums[d] = grid[i][j]
			d++
		}
	}
	k = k % (m * n)
	nums = append(nums[m*n-k:], nums[:m*n-k]...)
	d = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = nums[d]
			d++
		}
	}
	return grid
}
