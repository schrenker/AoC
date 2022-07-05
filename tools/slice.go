package tools

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

func StringPermutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
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

func Reverse[T any](s []T) []T {
	// for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
