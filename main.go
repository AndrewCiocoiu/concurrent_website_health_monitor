package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

func checkUrl(url string, client *http.Client, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	resp, err := client.Get(url)
	if err != nil {
		res := fmt.Sprintf("%s - DOWN\n", url)
		ch <- res
		return
	} else if resp.StatusCode == 200 {
		res := fmt.Sprintf("%s - UP\n", url)
		ch <- res
	} else {
		res := fmt.Sprintf("%s - DOWN\n", url)
		ch <- res
	}

	resp.Body.Close()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
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

	for val := range ch {
		fmt.Printf("%s", val)
	}

	wg.Wait()
}
