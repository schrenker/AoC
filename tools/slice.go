package tools

// func RemoveElement(slice interface{}, index int) interface{} {
// 	if v := reflect.ValueOf(slice); v.Len() > index {
// 		return reflect.AppendSlice(v.Slice(0, index), v.Slice(index+1, v.Len())).Interface()
// 	}
// 	return slice
// }

func RemoveElementString(s []string, idx int) []string {
	return append(s[:idx], s[idx+1:]...)
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
