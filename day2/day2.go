package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checker(stringVal string, len int) bool {
	firstHalf := stringVal[0 : len/2]
	secondHalf := stringVal[len/2 : len]
	if firstHalf == secondHalf {
		return true
	}
	return false
}

func main() {
	// content, err := os.ReadFile("./test-input.txt")
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var text string
	var result int

	result = 0
	text = string(content)
	text = strings.ReplaceAll(text, "\n", "")
	array := strings.Split(text, ",")

	for _, ranges := range array {
		values := strings.Split(ranges, "-")
		firstVal, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatal(err)
		}
		secondVal, err := strconv.Atoi(values[1])
		if err != nil {
			log.Fatal(err)
		}

		for firstVal <= secondVal {
			firstStringVal := strconv.Itoa(firstVal)
			numberLen := len(strconv.Itoa(firstVal))
			if numberLen%2 == 0 {
				if checker(firstStringVal, numberLen) == true {
					fmt.Println(firstStringVal)
					result += firstVal
				}
			}
			firstVal++
		}
	}
	fmt.Print("result : ", result)
}
