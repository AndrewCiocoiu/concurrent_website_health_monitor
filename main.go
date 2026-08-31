package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type WebsiteStatus struct {
	url    string
	status string
}

func worker(jobs chan string, client *http.Client, wg *sync.WaitGroup, ch chan WebsiteStatus) {
	defer wg.Done()

	for url := range jobs {
		resp, err := client.Get(url)

		new_website := WebsiteStatus{url: url}

		if err != nil {
			new_website.status = "DOWN"
			ch <- new_website
			continue
		} else if resp.StatusCode == 200 {
			new_website.status = "UP"
			ch <- new_website
		} else {
			new_website.status = "DOWN"
			ch <- new_website
		}

		resp.Body.Close()
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan WebsiteStatus)
	jobs := make(chan string)

	if len(os.Args) < 2 {
		log.Fatalf("You forgot to enter the filename as a CLI argument! Please do so!")
	}

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %s\n", err)
	}
	defer file.Close()

	client := &http.Client{Timeout: 5 * time.Second}

	//Launch 15 workers to prevent socket exhaustion
	for i := 0; i < 15; i++ {
		go worker(jobs, client, &wg, ch)
		wg.Add(1)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url := scanner.Text()
		jobs <- url
	}
	//Close the jobs channel as soon as I put all the jobs in so that my workers dont become deadlocked when they finnish all jobs
	close(jobs)

	if err := scanner.Err(); err != nil {
		log.Fatalf("There was an error while scanning the file: %s\n", err)
	}

	//Prevents deadlock
	go func() {
		wg.Wait()

		close(ch)
	}()

	output_file, err := os.Create("results.txt")
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	defer output_file.Close()

	for websiteStatus := range ch {
		_, err = output_file.WriteString(fmt.Sprintf("%s - %s\n", websiteStatus.url, websiteStatus.status))
		if err != nil {
			log.Fatalf("%s\n", err)
		}
	}

	fmt.Printf("Done! Results written in: results.txt!\n")

}
