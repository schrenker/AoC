package tools

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

// Get the highest number in the list
func GetMax(nums ...int) int {
	max := nums[0]
	for _, v := range nums {
		if max < v {
			max = v
		}
	}
	return max
}

func GetMaxMultiple(amount int, nums ...int) []int {
	sort.Ints(nums)
	nums = Reverse(nums)
	if len(nums) <= amount {
		return nums
	}
	return nums[:amount]
}

// Get the lowest number in the list
func GetMin(nums ...int) int {
	sort.Ints(nums)
	return nums[0]
}

// Get X lowest numbers in the list
func GetMinMultiple(amount int, nums ...int) []int {
	sort.Ints(nums)
	if len(nums) <= amount {
		return nums
	}
	return nums[:amount]
}

// Sum all numbers in the list
func SumAll(nums ...int) int {
	acc := 0
	for _, v := range nums {
		acc += v
	}
	return acc
}

// Multiply all numbers in the list by themselves
func MulAll(nums ...int) int {
	acc := 1
	for _, v := range nums {
		acc *= v
	}
	return acc
}

// Remove all numbers in the list from the highest number in the list
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

// checks if all numbers are equal to each other
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

// Positive modulo, returns non negative solution to x % d
func PMod(x, d int) int {
	x = x % d
	if x >= 0 {
		return x
	}
	if d < 0 {
		return x - d
	}
	return x + d
}

func StringMathToInt(s string) int {
	operands := strings.Fields(s)
	left, err := strconv.Atoi(operands[0])
	if err != nil {
		log.Fatalln(err)
	}
	right, err := strconv.Atoi(operands[2])
	if err != nil {
		log.Fatalln(err)
	}
	switch operands[1] {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	case "%":
		return left % right

	}
	log.Fatalf("No answer for: %v\n", s)
	return 0
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
