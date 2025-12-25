package main

import "fmt"

func main() {
	var nums []int = []int{2, 2, 1}
	fmt.Println(singleNumber(nums))
	// var hashSet map[int]int = map[int]int{}
	Print("5 & 3 =", 5&3)
	Print("5 | 3 =", 5|3)
	Print("5 ^ 3 =", 5^3)
	Print("^3 =", ^3)
}

func Print(a ...any) {
	fmt.Println(a...)
}

// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

func singleNumber1(nums []int) int {

	var hashSet map[int]int = map[int]int{}

	for _, v := range nums {
		hashSet[v]++
	}

	for k, v := range hashSet {
		if v == 1 {
			return k
		}
	}
	return -1
}

func singleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
	}
	return result
}

