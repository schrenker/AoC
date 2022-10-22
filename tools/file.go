package tools

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

func GetDefaultInputPath() string {
	p, err := exec.Command("go", "env", "GOMOD").Output()
	if err != nil {
		log.Fatalln(err)
	}
	return strings.TrimSpace(strings.Replace(string(p), "go.mod", "", 1)) + "input/" + os.Args[1] + "/" + os.Args[2] + ".txt"
}

func ReadFileString(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	return string(data)
}

func ReadFileStringSlice(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
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
		log.Fatalln(err)
	}
	return bytes.TrimSpace(data)
}
