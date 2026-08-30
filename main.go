package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type WebsiteStatus struct {
	url    string
	status string
}

func checkUrl(url string, client *http.Client, wg *sync.WaitGroup, ch chan WebsiteStatus) {
	defer wg.Done()
	resp, err := client.Get(url)

	new_website := WebsiteStatus{url: url}

	if err != nil {
		new_website.status = "DOWN"
		ch <- new_website
		return
	} else if resp.StatusCode == 200 {
		new_website.status = "UP"
		ch <- new_website
	} else {
		new_website.status = "DOWN"
		ch <- new_website
	}

	resp.Body.Close()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan WebsiteStatus)
	defer close(ch)

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
		go checkUrl(url, client, &wg, ch)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("There was an error while scanning the file: %s\n", err)
	}

	for websiteStatus := range ch {
		fmt.Printf("%s - %s\n", websiteStatus.url, websiteStatus.status)
	}

	wg.Wait()
}
