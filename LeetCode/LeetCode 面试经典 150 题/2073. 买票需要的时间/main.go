package main

import (
	"container/list"
	"fmt"
)

func main() {
	tickets := []int{2, 3, 2}
	fmt.Println(timeRequiredToBuy2(tickets, 2))
}

type Customer struct {
	Pos   int
	Count int
}

func timeRequiredToBuy(tickets []int, k int) int {

	if len(tickets) == 0 || k < 0 || k > len(tickets)-1 {
		return -1
	}

	var queue = list.New()
	for i, v := range tickets {
		queue.PushBack(Customer{
			Pos:   i,
			Count: v,
		})
	}

	var time int
	for queue.Len() > 0 {
		// 队列第一个取票
		var front Customer = queue.Front().Value.(Customer)
		time++
		// 买一张票，需要的票数减 1
		front.Count -= 1

		// 从队列中删除
		queue.Remove(queue.Front())

		if front.Count > 0 {
			// 如果还有需要购买的票，则继续排队
			queue.PushBack(front)
		} else {
			// 如果已经完成买票，判断是否是位置为k的那个人
			if front.Pos == k {
				break
			}
		}
	}
	return time
}

func timeRequiredToBuy2(tickets []int, k int) int {
	var queue []int = make([]int, len(tickets))
	for i := 0; i < len(tickets); i++ {
		queue[i] = i
	}

	var time int
	for len(queue) > 0 {
		frontIndex := queue[0]
		queue = queue[1:]
		time++

		tickets[frontIndex]--
		front := tickets[frontIndex]
		if frontIndex == k && front == 0 {
			return time
		}

		if front == 0 {
			continue
		}

		queue = append(queue, frontIndex)
	}
	return time
}
