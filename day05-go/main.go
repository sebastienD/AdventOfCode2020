package main

import (
	"bufio"
	"fmt"
	"go/types"
	"log"
	"os"
	"sort"
	"strconv"
)

func toi(in string) int {
	num, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

type ByValue []int

func (s ByValue) Len() int {
	return len(s)
}
func (s ByValue) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByValue) Less(i, j int) bool {
	return s[i] < s[j]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Open file failed")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var values []string
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}
	r := firstPart(values)
	sort.Sort(ByValue(r))
	seat := selectSeat(r)
	fmt.Printf("Seat %v\n", seat)
}

func selectSeat(values []int) int {
	exist := make(map[int]types.Nil)
	for _, v := range values {
		exist[v] = types.Nil{}
	}
	unknown := 1
	for i := 1; i < values[len(values)-1]; i++ {
		if _, contains := exist[i]; !contains {
			unknown = i
		}
	}
	return unknown
}

func max(values []int) int {
	max := 0
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

func firstPart(values []string) []int {
	result := make([]int, len(values))

	for _, value := range values {
		i, _ := compute(value[:7], 'F', 0, 127)
		_, a := compute(value[7:], 'L', 0, 7)
		result = append(result, i*8+a)
	}
	return result
}

func compute(val string, front rune, min int, max int) (int, int) {
	for _, c := range val {
		if c == front {
			max = max - (max-min)/2 - 1
			continue
		}
		min = min + ((max - min) / 2) + 1
	}
	return min, max
}
