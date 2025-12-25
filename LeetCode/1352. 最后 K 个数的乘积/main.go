package main

func main() {
	// [3,0,2,5,4] [2],[3],[4],[8],[2]
	obj := Constructor()
	obj.Add(3)
	obj.Add(0)
	obj.Add(2)
	obj.Add(5)
	obj.Add(4)
	obj.GetProduct(2)
	obj.GetProduct(3)
	obj.GetProduct(4)
	obj.GetProduct(8)
	obj.GetProduct(2)
}

type ProductOfNumbers struct {
	preNums []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		preNums: []int{1},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.preNums = []int{1}
		return
	}
	this.preNums = append(this.preNums, this.preNums[len(this.preNums)-1]*num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	l := len(this.preNums)
	if k >= l {
		return 0
	}
	return this.preNums[l-1] / this.preNums[l-1-k]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
