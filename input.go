package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func getInput(year, day string) {
	if !checkIfInputAlreadyExists(year, day) {
		err := ioutil.WriteFile(fmt.Sprintf("./input/%v/%v.txt", year, day), getData(year, day), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
