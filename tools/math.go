package tools

func VarMin(nums ...int) int {
	min := nums[0]
	for _, v := range nums {
		if min > v {
			min = v
		}
	}
	return min
}

func SumAll(nums ...int) int {
	acc := 0
	for _, v := range nums {
		acc += v
	}
	return acc
}
