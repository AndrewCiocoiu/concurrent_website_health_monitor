package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

func checkUrl(url string, client *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("%s - DOWN\n", url)
		return
	} else if resp.StatusCode == 200 {
		fmt.Printf("%s - UP\n", url)
	} else {
		fmt.Printf("%s - DOWN\n", url)
	}

	resp.Body.Close()
}

func main() {
	var wg sync.WaitGroup

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

		wg.Add(1)
		go checkUrl(url, client, &wg)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("There was an error while scanning the file: %s\n", err)
	}

	wg.Wait()
}
