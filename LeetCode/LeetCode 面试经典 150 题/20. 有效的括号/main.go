package main

import (
	"container/list"
	"fmt"
)

func main() {
	s := "{(){}}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	hashSet := make(map[rune]rune, 6)
	hashSet[')'] = '('
	hashSet[']'] = '['
	hashSet['}'] = '{'

	stack := list.New()
	for _, ch := range s {
		if stack.Len() != 0 && stack.Front().Value.(rune) == hashSet[ch] {
			stack.Remove(stack.Front())
		} else {
			stack.PushFront(ch)
		}
	}

	return stack.Len() == 0
}
