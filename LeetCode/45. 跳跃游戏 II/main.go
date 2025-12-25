package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	// 长度为2，只需要跳1步，长度为1，需要跳0步。
	if len(nums) <= 2 {
		return len(nums) - 1
	}
	// 如果第一步就能跳到结尾，返回1步
	if nums[0] >= len(nums)-1 {
		return 1
	}
	ans := 0                // 需要跳的步骤
	farthest := 0           // 当前能跳的最远距离
	start := 1              // 下一步的开始
	end := nums[0]          // 下一步结束
	for end < len(nums)-1 { // 如果end>=len(nums)-1,说明只需要最后一步就能到结尾
		for i := start; i <= end; i++ {
			if i >= len(nums)-1 {
				return ans + 1
			}
			if i+nums[i] > farthest {
				farthest = i + nums[i]
			}
		}
		start = end + 1
		end = farthest
		ans++
	}
	return ans + 1
}

func jump2(nums []int) int {
	steps := 0    // 步数
	end := 0      // 当前这一步能跳到的最远位置
	farthest := 0 // 下一步能跳到的最远位置

	for i := 0; i < len(nums)-1; i++ {
		// 更新下一步能跳的最远位置
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		// 到达边界，必须跳一次
		if i == end {
			steps++
			end = farthest // 更新边界
		}
	}

	return steps
}
