package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f := Read()
	r := firstPart(f)
	fmt.Printf("Result : %v\n", r)
	r = secondPart(f)
	fmt.Printf("Result : %v\n", r)
}

func firstPart(f []int) int {
	sort.Ints(f)
	diff := make(map[int]int)
	prev := 0
	for _, v := range f {
		diff[v-prev]++
		prev = v
	}
	fmt.Printf("diff %v\n", diff)
	return diff[1] * (diff[3] + 1)
}

func secondPart(f []int) int {
	f = append([]int{0}, f...)
	sort.Ints(f)
	comp := make([]int, len(f))
	comp[0] = 1
	for i := 1; i < len(f); i++ {
		if f[i]-f[i-1] <= 3 {
			comp[i] += comp[i-1]
		}
		if i > 1 && f[i]-f[i-2] <= 3 {
			comp[i] += comp[i-2]
		}
		if i > 2 && f[i]-f[i-3] <= 3 {
			comp[i] += comp[i-3]
		}
	}
	fmt.Printf("comp %v\n", comp)
	return comp[len(comp)-1]
}

func Read() []int {
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
	return values
}

func toi(in string) int {
	num, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

type ByValue []int64

func (s ByValue) Len() int {
	return len(s)
}

func (s ByValue) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByValue) Less(i, j int) bool {
	return s[i] < s[j]
}
