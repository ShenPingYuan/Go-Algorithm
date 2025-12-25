package main

import (
	"fmt"
)

func main() {
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println("arr:", arr)
	var slice = arr[1:3]
	fmt.Println("slice:", slice, "len:", len(slice), "cap:", cap(slice))

	fmt.Println("&arr1:", &arr[1], "&slice0:", &slice[0])
	slice[0] = 99
	fmt.Println("slice:", slice, "arr:", arr)
	slice = append(slice, 11, 12, 13, 14, 15, 16, 17, 18)
	fmt.Println("slice:", slice, "arr:", arr)
	fmt.Println("slice:", slice, "len:", len(slice), "cap:", cap(slice))

	slice2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("slice2:", slice2, "len:", len(slice2), "cap:", cap(slice2))
	changeSlice(slice2)
	fmt.Println("changeSlice2:", slice2, "len:", len(slice2), "cap:", cap(slice2))

	var slice3 []int
	fmt.Println("slice3:", slice3, "len:", len(slice3), "cap:", cap(slice3))
	slice4 := append(slice3, 1, 2, 3)
	fmt.Println("slice4:", slice4, "len:", len(slice4), "cap:", cap(slice4))
}

func changeSlice(slice []int) {
	slice = append(slice, 7, 8, 9, 10)
	fmt.Println("changeSlice:", slice, "len:", len(slice), "cap:", cap(slice))
}
