package day3

import (
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("./test-input.txt")
	// content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// result := 0
	text := string(content)
	fmt.Print(text)
}
