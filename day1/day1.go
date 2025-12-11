package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./test-input.txt")
	// file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var number int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := (scanner.Text())
		direction := line[0]
		if line[0] == 'L' {
			number = line.lenlen()
		} else if line[0] == 'R' {

		}

		fmt.Println(direction)
		// fmt.Println(line)
		// fmt.Printf(scanner.Text())
		// fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// input, err := io.ReadAll(file) //gives ASCII values
	// fmt.Print(string(test_input))
}
