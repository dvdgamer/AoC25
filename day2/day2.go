package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checker(val int, stringVal string, len int) int {
	// slice number in half
	// check if halved numbers are the same
	// return number
	return 0
}

func main() {
	content, err := os.ReadFile("./test-input.txt")
	// content, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var text string

	text = string(content)
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
			firstStringVal := strconv.Itoa(firstVal)
			if numberLen%2 == 0 {
				checker(firstVal, firstStringVal, numberLen)
			}
			firstVal++
		}
	}
}
