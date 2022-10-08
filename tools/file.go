package tools

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

func getInputPath() string {
	p, err := exec.Command("go", "env", "GOMOD").Output()
	if err != nil {
		log.Fatalln(err)
	}
	return strings.TrimSpace(strings.Replace(string(p), "go.mod", "", 1)) + "input/" + os.Args[1] + "/" + os.Args[2] + ".txt"
}

func ReadFileString() string {
	data, err := os.ReadFile(getInputPath())
	if err != nil {
		log.Fatalln(err)
	}
	return string(data)
}

func ReadFileStringSlice() []string {
	data, err := os.ReadFile(getInputPath())
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
	data, err := os.ReadFile(getInputPath())
	if err != nil {
		log.Fatalln(err)
	}
	return bytes.TrimSpace(data)
}
