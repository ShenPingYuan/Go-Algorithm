package main

import "fmt"

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	minEatingSpeed(piles, h)
}

func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := 1
	for i := 0; i < len(piles); i++ {
		if piles[i] > right {
			right = piles[i]
		}
	}
	// fmt.Println(left, right)
	for left <= right {
		mid := left + (right-left)/2
		var t = time(piles, mid)
		fmt.Println(mid, t)
		if t == h {
			right = mid - 1
		} else if t > h {
			left = mid + 1
		} else if t < h {
			right = mid - 1
		}
	}
	return left
}

// 随着吃香蕉的速度x增加，吃完所有香蕉的速度time单调递减
func time(piles []int, x int) int {
	t := 0
	for _, pile := range piles {
		// fmt.Println(pile)
		t += pile / x
		if pile%x != 0 {
			t++
		}
	}
	return t
}
