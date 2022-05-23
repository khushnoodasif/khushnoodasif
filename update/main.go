package main

import (
	"io"
	"os"
	"time"
)

func makeReadme(filename string) error {

	date := time.Now().Format("2 Jan 2006 3:4:5 pm")
	time_new := "<sub>Last updated: " + date + ".</sub>"

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, time_new)
	if err != nil {
		return err
	}
	return file.Sync()
}

func main() {

	makeReadme("../README.md")

}