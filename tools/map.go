package tools

func CloneMap[T1 comparable, T2 any](m map[T1]T2) map[T1]T2 {
	result := map[T1]T2{}
	for k, v := range m {
		result[k] = v
	}
	return result
}
