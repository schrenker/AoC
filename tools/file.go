package tools

import (
	"fmt"
	"os"
	"strings"
)

func ReadFileString(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("No file %v found", path))
	}
	return string(data)
}

func ReadFileStringSlice(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("No file %v found", path))
	}
	spl := strings.Split(string(data), "\n")
	if spl[len(spl)-1] == "" {
		return spl[:len(spl)-1]
	}
	return spl
}

func ReadFileBytes(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("No file %v found", path))
	}
	return data
}
