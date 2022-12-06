package y2022

type Day06 struct{}

func scanFirstX(str string, n int) int {
	for i := n - 1; i < len(str); i++ {
		occurrences := make(map[byte]bool)
		for j := i - (n - 1); j <= i; j++ {
			if _, ok := occurrences[str[j]]; ok {
				break
			}
			occurrences[str[j]] = true
		}
		if len(occurrences) == n {
			return i + 1
		}
	}
	return -1
}

func (d Day06) PartOne(data []byte) interface{} {
	return scanFirstX(string(data), 4)
}

func (d Day06) PartTwo(data []byte) interface{} {
	return scanFirstX(string(data), 14)
}
