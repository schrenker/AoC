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

func Decummulate(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	acc := nums[len(nums)-1]
	for i := len(nums) - 2; i <= 0; i-- {
		acc = acc - nums[i]
	}
	return acc
}

func MultipleEqualInt(nums ...int) bool {
	if len(nums) < 2 {
		return true
	}
	for i := 0; i < len(nums); i++ {
		if nums[0] != nums[i] {
			return false
		}
	}
	return true
}
