package tools

import "math"

func RemoveElement[T any](s []T, idx int) []T {
	return append(s[:idx], s[idx+1:]...)
}

func RemoveElementValue[T comparable](s []T, val T) []T {
	for i := range s {
		if s[i] == val {
			s = RemoveElement(s, i)
		}
	}
	return s
}

func Permutations[T comparable](arr []T) [][]T {
	var helper func([]T, int)
	res := [][]T{}

	helper = func(arr []T, n int) {
		if n == 1 {
			tmp := make([]T, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func Combinations[T any](source []T) [][]T {
	combinationList := make([][]T, 0)

	for i := 1; i < int(math.Pow(2, float64(len(source)))); i++ {
		list := make([]T, 0)
		for j := 0; j < len(source); j++ {
			if i&int(math.Pow(2, float64(j))) > 0 {
				list = append(list, source[j])
			}
		}
		combinationList = append(combinationList, list)
	}

	return combinationList
}

func CombinationsNArrays[T any](source [][]T) [][]T {
	n := len(source)
	indices := make([]int, n)
	result := make([][]T, 0)

	for {
		tmp := make([]T, 0)
		for i := 0; i < n; i++ {
			tmp = append(tmp, source[i][indices[i]])
		}
		result = append(result, tmp)

		next := n - 1
		for next >= 0 && (indices[next]+1 >= len(source[next])) {
			next--
		}

		if next < 0 {
			return result
		}

		indices[next]++

		for i := next + 1; i < n; i++ {
			indices[i] = 0
		}
	}
}

// Reverse takes in a slice of any type.
// It returns the slice reversed
func Reverse[T any](s []T) []T {
	// for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// Map turns a []T1 to a []T2 using a mapping function.
// This function has two type parameters, T1 and T2.
// This works with slices of any type.
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduce reduces a []T1 to a single value using a reduction function.
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filter filters values from a slice using a filter function.
// It returns a new slice with only the elements of s
// for which f returned true.
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// RemoveDuplicates takes slice s as an input, and returns a slice with duplicate values removed
func RemoveDuplicates[T comparable](s []T) []T {
	enc := make(map[T]bool)
	for _, v := range s {
		enc[v] = true
	}
	result := []T{}
	for k := range enc {
		result = append(result, k)
	}
	return result
}
