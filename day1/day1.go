package main

import (
	"bufio"
	"fmt"
	"log"
	// "math"
	"os"
	"strconv"
)

func main() {
	// file, err := os.Open("./test-input.txt")
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := 0
	dial := 50
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		if line[0] == 'L' {
			number = number % 100
			if number > dial {
				dial = dial + 100
			}
			dial -= number
		} else if line[0] == 'R' {
			dial += number
		} else {
			log.Fatal(err)
		}
		if dial > 99 {
			dial = dial % 100
		}
		if dial == 0 {
			res++
		}
		fmt.Println(dial)
	}
	fmt.Println("res:", res)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
