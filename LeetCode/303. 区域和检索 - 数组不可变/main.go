package main

func main() {

}

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sum := 0
	preSums := make([]int, len(nums)+1)
	for i := range nums {
		sum += nums[i]
		preSums[i+1] = sum
	}
	return NumArray{
		sums: preSums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
