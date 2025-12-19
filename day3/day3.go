package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// content, err := os.ReadFile("./test-input.txt")
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var result int
	var biggestJoltage [2]byte

	text := string(content)
	for line := range strings.SplitSeq(text, "\n") {
		if len(line) == 0 {
			break
		}
		var biggestNumber byte = byte('0')
		var secondBiggestNumber byte = byte('0')

		i := 0

		indexOfBiggestNumber := 0
		for i < len(line)-1 {

			if line[i] > biggestNumber {
				biggestNumber = byte(line[i])
				indexOfBiggestNumber = i
			}
			i++
		}
		biggestJoltage[0] = biggestNumber
		i = indexOfBiggestNumber + 1
		secondBiggestNumber = byte(line[i])
		for i < len(line) {
			if line[i] > secondBiggestNumber {
				secondBiggestNumber = byte(line[i])
			}
			i++
		}
		biggestJoltage[1] = secondBiggestNumber
		combinedString := string(biggestJoltage[:])
		num, err := strconv.Atoi(string(combinedString))
		if err != nil {
			log.Fatal(err)
		}
		result += num
	}
	fmt.Print("result: ", result)
}
