package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func repeatedPattern(stringVal string, len int) bool {
	i := 1
	for i <= len/2 {
		slice := stringVal[0:i]
		restOfString := stringVal[i:]
		remainingChars := strings.ReplaceAll(restOfString, slice, "")
		if remainingChars == "" {
			return true
		}
		i++
	}
	return false
}

func main() {
	// content, err := os.ReadFile("../test-input.txt")
	content, err := os.ReadFile("../input.txt")
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
			if repeatedPattern(firstStringVal, numberLen) == true {
				result += firstVal
			}
			firstVal++
		}
	}
	fmt.Print("result : ", result)
}
