package main

func main() {

}

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)
	for _, b := range bookings {
		f, l, v := b[0], b[1], b[2]
		diff[f-1] += v
		if l < n {
			diff[l] -= v
		}
	}

	res := make([]int, n)
	res[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		res[i] = res[i-1] + diff[i]
	}
	return res
}
