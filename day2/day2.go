package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checker(val int, len int) int {
	return 1
}

func main() {

	content, err := os.ReadFile("./test-input.txt")
	// content, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	text = strings.ReplaceAll(text, "\n", "")
	array := strings.Split(text, ",")
	fmt.Print(array)
	for _, ranges := range array {
		// fmt.Println(ranges)
		values := strings.Split(ranges, "-")
		firstVal, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatal(err)
		}
		secondVal, err := strconv.Atoi(values[1])
		if err != nil {
			log.Fatal(err)
		}

		numberLen := len(strconv.Itoa(firstVal))
		for firstVal < secondVal {
			if numberLen%2 == 0 {
				checker(firstVal, numberLen)
			}
			firstVal++
		}
	}
}
