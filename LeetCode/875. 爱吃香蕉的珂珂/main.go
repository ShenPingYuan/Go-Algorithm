package main

func main() {

}

func minEatingSpeed(piles []int, h int) int {
	max := 1000000000 + 1
	// for _, v := range piles {
	// 	if v > max {
	// 		max = v
	// 	}
	// }
	left, right := 1, max
	for left <= right {
		mid := left + (right-left)/2
		if time(piles, mid) <= h {
			right = mid - 1
		} else if time(piles, mid) > h {
			left = mid + 1
		}
	}
	return left
}

// 随着吃香蕉的速度x增加，吃完所有香蕉的速度time单调递减
func time(piles []int, x int) int {
	t := 0
	for i := 0; i < len(piles); i++ {
		t += piles[i] / x
		if piles[i]%x > 0 {
			t++
		}
	}
	return t
}
