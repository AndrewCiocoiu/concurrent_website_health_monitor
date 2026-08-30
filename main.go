package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("You forgot to enter the filename as a CLI argument! Please do so!")
	}

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %s\n", err)
	}
	defer file.Close()

	client := &http.Client{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url := scanner.Text()

		resp, err := client.Get(url)
		if clean_url := url[8:]; err != nil {
			fmt.Printf("%s - DOWN\n", clean_url)
			continue
		} else if resp.StatusCode == 200 {
			fmt.Printf("%s - UP\n", clean_url)
		} else {
			fmt.Printf("%s - DOWN\n", clean_url)
		}

		resp.Body.Close()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("There was an error while scanning the file: %s\n", err)
	}
}
