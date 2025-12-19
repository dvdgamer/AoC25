package main

import (
	"log"
	"os"
)

func main() {
	// content, err := os.ReadFile("../input.txt")
	content, err := os.ReadFile("../test-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
}
