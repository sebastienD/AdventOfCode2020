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

type operations map[string]func(int, int, int) (int, int)

func NewOerations() operations {
	op := make(map[string]func(int, int, int) (int, int))
	op["nop"] = func(acc int, index int, val int) (int, int) {
		//fmt.Printf("nop -> acc %v, index %v\n", acc, index+1)
		return acc, index + 1
	}
	op["acc"] = func(acc int, index int, val int) (int, int) {
		//fmt.Printf("acc -> acc %v, index %v\n", acc+val, index+1)
		return acc + val, index + 1
	}
	op["jmp"] = func(acc int, index int, val int) (int, int) {
		//fmt.Printf("jmp -> acc %v, index %v\n", acc, index+val)
		return acc, index + val
	}
	return op
}

func main() {
	f := Read()
	r := firstPart(f)
	fmt.Printf("Accumulator : %v\n", r)
	r = secondPart(f)
	fmt.Printf("Accumulator : %v\n", r)
}

func secondPart(ref []string) int {
	instructions := toArray(ref)
	nextIndex := 0
	for {
		instr, newIndex := changeInstruction(instructions, nextIndex)
		if ok, acc := run(instr); ok {
			return acc
		}
		nextIndex = newIndex + 1
	}
	return 0
}

func toArray(ref []string) [][]string {
	new := [][]string{}
	for _, v := range ref {
		args := strings.Split(v, " ")
		new = append(new, args)
	}
	return new
}

func run(instructions [][]string) (bool, int) {
	op := NewOerations()
	acc, index := 0, 0
	view := make(map[int]types.Nil)
	for _, ok := view[index]; !ok; _, ok = view[index] {
		view[index] = types.Nil{}
		args := instructions[index]
		acc, index = op[args[0]](acc, index, toi(args[1]))
		if index >= len(instructions) {
			return true, acc
		}
	}
	return false, acc
}

func changeInstruction(param [][]string, index int) ([][]string, int) {
	ref := [][]string{}
	for _, v := range param {
		n := []string{}
		for _, l := range v {
			n = append(n, l)
		}
		ref = append(ref, n)
	}
	for i := index; i < len(ref); i++ {
		if ref[i][0] == "nop" {
			ref[i][0] = "jmp"
			return ref, i
		}
		if ref[i][0] == "jmp" {
			ref[i][0] = "nop"
			return ref, i
		}
	}
	return ref, index
}

func firstPart(param []string) int {
	op := NewOerations()
	instructions := toArray(param)
	acc, index := 0, 0
	view := make(map[int]types.Nil)
	for _, ok := view[index]; !ok; _, ok = view[index] {
		view[index] = types.Nil{}
		args := instructions[index]
		acc, index = op[args[0]](acc, index, toi(args[1]))
	}
	return acc
}

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

func toi(in string) int {
	num, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
