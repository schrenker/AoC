package tools

import "sort"

func GetMin(nums ...int) int {
	sort.Ints(nums)
	return nums[0]
}

func GetMinMultiple(amount int, nums ...int) []int {
	sort.Ints(nums)
	if len(nums) <= amount {
		return nums
	}
	return nums[:amount]
}

func SumAll(nums ...int) int {
	acc := 0
	for _, v := range nums {
		acc += v
	}
	return acc
}

func MulAll(nums ...int) int {
	acc := 1
	for _, v := range nums {
		acc *= v
	}
	return acc
}
