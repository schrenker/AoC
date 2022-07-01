package tools

import (
	"bytes"
	"log"
	"os"
	"strings"
)

func ReadFileString() string {
	data, err := os.ReadFile("./input/" + os.Args[1] + "/" + os.Args[2] + ".txt")
	if err != nil {
		log.Fatalln(err)
	}
	return string(data)
}

func ReadFileStringSlice() []string {
	data, err := os.ReadFile("./input/" + os.Args[1] + "/" + os.Args[2] + ".txt")
	if err != nil {
		log.Fatalln(err)
	}
	spl := strings.Split(string(data), "\n")
	if spl[len(spl)-1] == "" {
		return spl[:len(spl)-1]
	}
	return spl
}

func ReadFileBytes() []byte {
	data, err := os.ReadFile("./input/" + os.Args[1] + "/" + os.Args[2] + ".txt")
	if err != nil {
		log.Fatalln(err)
	}
	return bytes.TrimSpace(data)
}
