package main

import (
	"bufio"
	"fmt"
	"log"
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
			for number > 99 {
				number -= 100
				res++
			}
			if dial == number {
				res++
			}
			if number > dial {
				if dial == 0 {
					res--
				}
				dial = dial + 100
				res++
			}
			dial -= number
		} else if line[0] == 'R' {
			if number > 99 {
				for number > 99 {
					number -= 100
					res++
				}
			}
			dial += number
			if dial > 99 {
				dial -= 100
				res++
			}
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println("\nres:", res)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
