package main

func main() {

}

func carPooling(trips [][]int, capacity int) bool {
	max := 0
	for _, v := range trips {
		if v[2] > max {
			max = v[2]
		}
	}

	diff := make([]int, max+1)
	for _, v := range trips {
		n, f, t := v[0], v[1], v[2]
		diff[f] += n
		if t < max {
			diff[t] -= n
		}
	}

	count := 0
	for i := 0; i < max; i++ {
		count += diff[i]
		if count > capacity {
			return false
		}
	}
	return true
}
