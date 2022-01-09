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
