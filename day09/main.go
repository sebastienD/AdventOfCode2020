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
	window := 25
	r := firstPart(f, window)
	fmt.Printf("Bad value %v for window %v\n", r, window)
	r = secondPart(f, r)
	fmt.Printf("Weakness : %v\n", r)
}

func firstPart(ref []int64, window int) int64 {
	for index := 0 + window; index < len(ref); index++ {
		win := nextWindow(ref, index-window, window)
		if !check(ref[index], win) {
			return ref[index]
		}
	}
	return 0
}

func nextWindow(ref []int64, index int, window int) []int64 {
	win := []int64{}
	for i := index; i < window+index; i++ {
		win = append(win, ref[i])
	}
	return win
}

func check(v int64, win []int64) bool {
	for c := 0; c < len(win)-1; c++ {
		for i := c + 1; i < len(win); i++ {
			if win[c]+win[i] == v {
				return true
			}
		}
	}
	return false
}

func secondPart(ref []int64, v int64) int64 {
	for c := 0; c < len(ref)-1; c++ {
		t := ref[c]
		save := []int64{ref[c]}
		for i := c + 1; i < len(ref); i++ {
			t += ref[i]
			save = append(save, ref[i])
			if t == v {
				return extractValue(save)
			}
			if t > v {
				break
			}
		}
	}
	return 0
}

func extractValue(t []int64) int64 {
	sort.Sort(ByValue(t))
	return t[0] + t[len(t)-1]
}

func Read() []int64 {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Open file failed")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var values []int64
	for scanner.Scan() {
		values = append(values, toi64(scanner.Text()))
	}
	return values
}

func toi64(s string) int64 {
	num, err := strconv.ParseInt(s, 10, 64)
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
