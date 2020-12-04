package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func toi(in string) int {
	num, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Open file failed")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var values []int
	for scanner.Scan() {
		values = append(values, toi(scanner.Text()))
	}
	r := firstPart(values)
	fmt.Printf("Total %v\n", r)
	r = secondPart(values)
	fmt.Printf("Total %v\n", r)
}

func firstPart(values []int) int {
	for index, value := range values {
		for i := index + 1; i < len(values)-1; i++ {
			if value+values[i] == 2020 {
				return value * values[i]
			}
		}
	}
	fmt.Printf("Not found\n")
	return 0
}

func secondPart(values []int) int {
	for index, value := range values {
		for i := index + 1; i < len(values)-1; i++ {
			for j := index + 2; j < len(values)-1; j++ {
				if value+values[i]+values[j] == 2020 {
					return value * values[i] * values[j]
				}
			}
		}
	}
	fmt.Printf("Not found\n")
	return 0
}
