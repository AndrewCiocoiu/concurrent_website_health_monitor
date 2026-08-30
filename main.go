package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("targets.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s\n", err)
	}
	defer file.Close()
}
