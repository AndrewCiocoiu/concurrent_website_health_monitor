package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("targets.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s\n", line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("There was an error while scanning the file: %s\n", err)
	}
}
