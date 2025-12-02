package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), ",")

	numbersToCheck := []int{}
	for _, line := range lines {
		rangeArray := strings.Split(line, "-")
		if len(rangeArray) != 2 {
			// Handle invalid line format, e.g., skip or log error
			fmt.Printf("Skipping invalid line: %s\n", line)
			continue
		}
		lowerRange, err := strconv.Atoi(rangeArray[0])
		if err != nil {
			panic(err)
		}
		upperRange, err := strconv.Atoi(rangeArray[1])
		if err != nil {
			panic(err)
		}
		for i := lowerRange; i <= upperRange; i++ {
			strDigit := strconv.Itoa(i)
			if len(strDigit)%2 != 0 {
				continue
			}
			fullLength := len(strDigit)
			halfLength := fullLength / 2
			if strDigit[:halfLength] == strDigit[halfLength:] {
				numbersToCheck = append(numbersToCheck, i)
				fmt.Println(i)
			}
		}
	}

	fmt.Println("Total numbers obtained:", len(numbersToCheck))
	sum := 0
	for _, number := range numbersToCheck {
		sum += number
	}
	fmt.Println("Final sum:", sum)
}
