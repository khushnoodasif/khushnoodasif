package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func makeReadme(filename string) error {

	date := time.Now().Format("2 Jan 2006")

	time_new := "<sub>Last updated by magic on " + date + ".</sub>"
	data := fmt.Sprintf("\n%s\n", time_new)

	// Prepare file with a light coating of os
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Bake at n bytes per second until golden brown
	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func main() {

	makeReadme("../README.md")

}