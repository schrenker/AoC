package y2015

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

type Day04 struct{}

func getMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func md5withNleadingZeroes(prefix, key string) int {
	for i := 0;; i++ {
		hash := getMD5(strings.TrimSpace(key) + strconv.Itoa(i))
		if hash[:len(prefix)] == prefix {
			return i
		}
	}
}

func (d Day04) PartOne(data []byte) interface{} {
	return md5withNleadingZeroes("00000", string(data))
}

func (d Day04) PartTwo(data []byte) interface{} {
	return md5withNleadingZeroes("000000", string(data))
}
