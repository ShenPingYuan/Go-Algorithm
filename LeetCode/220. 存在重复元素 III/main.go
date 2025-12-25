package main

func main() {

}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if indexDiff <= 0 || valueDiff < 0 {
		return false
	}

	bucket := make(map[int]int)
	w := valueDiff + 1 // 桶的宽度
	left, right := 0, 0
	for right < len(nums) {
		new := nums[right]
		k := getBucketID(new, w)
		if _, f := bucket[k]; f {
			return true
		} else if v, f := bucket[k-1]; f && abs(new-v) <= valueDiff {
			return true
		} else if v, f := bucket[k+1]; f && abs(new-v) <= valueDiff {
			return true
		}
		bucket[k] = new
		right++

		if right-left > indexDiff {
			d := nums[left]
			dK := getBucketID(d, w)
			delete(bucket, dK)
			left++
		}
	}
	return false
}

func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	if k <= 0 || t < 0 {
		return false
	}

	bucket := make(map[int]int) // bucketID -> value
	w := t + 1                  // 桶的宽度

	for i, num := range nums {
		id := getBucketID(num, w)

		// 同一个桶
		if _, ok := bucket[id]; ok {
			return true
		}

		// 相邻桶
		if val, ok := bucket[id-1]; ok && abs(num-val) < w {
			return true
		}
		if val, ok := bucket[id+1]; ok && abs(num-val) < w {
			return true
		}

		// 插入当前桶
		bucket[id] = num

		// 超出窗口大小 k，删除最早的元素
		if i >= k {
			delete(bucket, getBucketID(nums[i-k], w))
		}
	}

	return false
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func getBucketID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1 // 负数桶编号特殊处理
}

func getBucket(value, valueDiff int) int {
	r := valueDiff + 1
	if value >= 0 {
		return value / r
	}
	return (value+1)/r - 1
}
