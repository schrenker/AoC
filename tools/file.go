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
	return strings.Split(string(data), "\n")
}

func ReadFileBytes(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("No file %v found", path))
	}
	return data
}
