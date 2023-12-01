package input

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func truncateDay(day string) string {
	if day[0] == '0' {
		return day[1:]
	}
	return day
}

func getData(year, day string) []byte {
	client := &http.Client{}

	cookie := &http.Cookie{
		Name:  "session",
		Value: os.Getenv("session"),
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, truncateDay(day)), nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.AddCookie(cookie)
	req.Header.Set("User-Agent", "Go http client by sebastian@zawadzki.tech - github.com/schrenker/AoC")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return body
}

func checkIfInputAlreadyExists(year, day string) bool {
	_, err := os.Stat(fmt.Sprintf("./input/%v/%v.txt", year, day))
	if os.IsNotExist(err) {
		err := os.MkdirAll(fmt.Sprintf("./input/%v/", year), 0755)
		if err != nil {
			log.Fatalln(err)
		}
		return false
	}
	return true
}

func GetInput(year, day string) {
	if !checkIfInputAlreadyExists(year, day) {
		err := os.WriteFile(fmt.Sprintf("./input/%v/%v.txt", year, day), getData(year, day), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func GetDefaultInputPath() string {
	p, err := exec.Command("go", "env", "GOMOD").Output()
	if err != nil {
		log.Fatalln(err)
	}
	return strings.TrimSpace(strings.Replace(string(p), "go.mod", "", 1)) + "input/" + os.Args[1] + "/" + os.Args[2] + ".txt"
}

func byteToStringSlice(data []byte) []string {
	spl := strings.Split(string(data), "\n")
	if spl[len(spl)-1] == "" {
		return spl[:len(spl)-1]
	}
	return spl
}

func byteToIntGrid(data []byte) [][]int {
	result := make([][]int, 0)
	tmp := make([]int, 0)
	for i := range data {
		if data[i] == '\n' {
			result = append(result, tmp)
			tmp = make([]int, 0)
			continue
		}
		num, _ := strconv.Atoi(string(data[i]))
		tmp = append(tmp, num)
	}
	return result
}

func ReadFileBytes(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
