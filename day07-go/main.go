package main

import (
	"bufio"
	"fmt"
	"go/types"
	"log"
	"os"
	"strconv"
	"strings"
)

func Read() []string {
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
	return values
}

func main() {
	f := Read()
	r := firstPart(f)
	fmt.Printf("Nb bag colors %v\n", r)
	r = secondPart(f)
	fmt.Printf("Nb nested bag %v\n", r)
}

func secondPart(content []string) int {
	bags := buildBags(content)
	target := "shiny gold"
	nestedBags := nested(target, bags)
	return nestedBags
}

func nested(target string, bags map[string]map[string]int) int {
	count := 0
	for bag, nb := range bags[target] {
		count += nb + nb*nested(bag, bags)
	}
	return count
}

func firstPart(content []string) int {
	bags := buildBags(content)
	target := "shiny gold"
	uniqueBag := make(map[string]types.Nil)
	search(target, bags, &uniqueBag)
	return len(uniqueBag)
}

func search(target string, bags map[string]map[string]int, uniqueBag *map[string]types.Nil) {
	contains := []string{}
	for bag, v := range bags {
		for innerBag := range v {
			if innerBag == target {
				contains = append(contains, bag)
				break
			}
		}
	}
	for _, bag := range contains {
		(*uniqueBag)[bag] = types.Nil{}
		search(bag, bags, uniqueBag)
	}
}

func Split(r rune) bool {
	return r == '.' || r == ',' || r == ' '
}

func buildBags(content []string) map[string]map[string]int {
	bags := make(map[string]map[string]int)
	for _, line := range content {
		line = strings.ReplaceAll(line, "bags", "")
		line = strings.ReplaceAll(line, "bag", "")
		line = strings.ReplaceAll(line, "contain", "")
		parts := strings.FieldsFunc(line, Split)
		key := parts[0] + " " + parts[1]
		if _, ok := bags[key]; !ok {
			bags[key] = make(map[string]int)
		}
		if parts[2] == "no" {
			continue
		}
		for i := 3; i < len(parts); i += 3 {
			bags[key][parts[i]+" "+parts[i+1]] = toi(parts[i-1])
		}
	}
	return bags
}

func toi(in string) int {
	num, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
