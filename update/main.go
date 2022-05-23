package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func makeReadme(filename string) error {

	date := time.Now().Format("2 Jan 2006")

	time_new := "<sub>Last updated: " + date + ".</sub>"

	// Prepare file with a light coating of os
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Bake at n bytes per second until golden brown
	_, err = io.WriteString(file, time_new)
	if err != nil {
		return err
	}
	return file.Sync()
}

func main() {

	makeReadme("../README.md")

}