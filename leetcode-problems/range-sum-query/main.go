package main

import "fmt"

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	return NumArray{nums: prefix}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left < 0 || right < left || right > len(this.nums) {
		panic(1)
	}

	return this.nums[right+1] - this.nums[left]
}

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(obj.SumRange(2, 5))
}
