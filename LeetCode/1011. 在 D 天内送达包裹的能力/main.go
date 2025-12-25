package main

func main() {
	// weights = [1,2,3,4,5,6,7,8,9,10] days = 5
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	days := 5
	shipWithinDays(weights, days)
}

func shipWithinDays(weights []int, days int) int {
	left, right := getMinMax(weights)
	for left <= right {
		mid := left + (right-left)/2
		shipDays := shipDays(weights, mid)
		if shipDays == days {
			right = mid - 1
		} else if shipDays < days {
			right = mid - 1
		} else if shipDays > days {
			left = mid + 1
		}
	}
	return left
}

func shipDays(weights []int, cap int) int {
	leftCap := 0
	days := 0
	for _, v := range weights {
		if leftCap < v {
			days++
			leftCap = cap
		}
		leftCap -= v
	}
	return days
}

func getMinMax(weights []int) (int, int) {
	sum := 0
	max := 0
	for _, v := range weights {
		sum += v
		if v > max {
			max = v
		}
	}
	return max, sum
}
