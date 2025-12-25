package main

import (
	"fmt"
)

func main() {
	// 初始化 map
	myMap := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
	}

	// 定义遍历 map 的函数
	printMapKeys := func(m map[string]int) {
		for key := range m {
			fmt.Print(key, " ")
		}
		fmt.Println()
	}

	// 多次遍历 map，观察键的顺序
	printMapKeys(myMap)
	printMapKeys(myMap)
	printMapKeys(myMap)
	printMapKeys(myMap)

	var a = 5
	var b = 6

	a, b = b, a

	fmt.Println(a, b)
}
