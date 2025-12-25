package main

func main() {
	//hours = [9,9,6,0,6,6,9]
	longestWPI([]int{9, 9, 6, 0, 6, 6, 9})
}

func longestWPI(hours []int) int {
	preSum := make([]int, len(hours)+1)
	for i := 0; i < len(hours); i++ {
		v := 0
		if hours[i] > 8 {
			v = 1
		} else {
			v = -1
		}
		preSum[i+1] = preSum[i] + v
	}

	max := 0
	for i := 0; i < len(preSum) && i < len(preSum)-max; i++ {
		for j := i + max + 1; j < len(preSum); j++ {
			sub := preSum[j] - preSum[i]
			if sub > 0 && j-i > max {
				max = sub
			}
		}
	}
	return max
}

func longestWPI1(hours []int) int {
	n := len(hours)
	prefix := make([]int, n+1) // 前缀和
	for i := 0; i < n; i++ {
		if hours[i] > 8 {
			prefix[i+1] = prefix[i] + 1
		} else {
			prefix[i+1] = prefix[i] - 1
		}
	}

	// 单调递减栈：存储前缀和的下标
	stack := []int{}
	for i := 0; i < len(prefix); i++ {
		if len(stack) == 0 || prefix[i] < prefix[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}

	res := 0
	// 倒序扫描（从右往左找最大 j - i）
	for j := len(prefix) - 1; j >= 0; j-- {
		for len(stack) > 0 && prefix[j] > prefix[stack[len(stack)-1]] {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if j-i > res {
				res = j - i
			}
		}
	}

	return res
}

func longestWPI2(hours []int) int {
	n := len(hours)
	prefix := make([]int, n+1) // 前缀和
	for i := 0; i < n; i++ {
		if hours[i] > 8 {
			prefix[i+1] = prefix[i] + 1
		} else {
			prefix[i+1] = prefix[i] - 1
		}
	}

	res := 0
	m := map[int]int{}
	for i := 0; i < len(prefix); i++ {
		if _, f := m[prefix[i]]; !f {
			m[prefix[i]] = i
		}
	}

	for i := 0; i < len(prefix); i++ {
		if prefix[i] > 0 {
			res = i
		} else {
			if v, f := m[prefix[i]-1]; f && v < i && i-v > res {
				res = i - v
			}
		}
	}

	return res
}

func longestWPI0(hours []int) int {
	n := len(hours)
	prefix := make([]int, n+1) // 前缀和
	for i := 0; i < n; i++ {
		if hours[i] > 8 {
			prefix[i+1] = prefix[i] + 1
		} else {
			prefix[i+1] = prefix[i] - 1
		}
	}

	res := 0
	stack := []int{0}
	for i := 1; i < len(prefix); i++ {
		if prefix[i] < prefix[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}

	for i := len(prefix) - 1; i >= 0; i-- {
		for len(stack) > 0 && prefix[i] > prefix[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if i-j > res {
				res = i - j
			}
		}
	}

	return res
}
