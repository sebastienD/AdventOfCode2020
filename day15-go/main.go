package main

import "fmt"

func main() {
	input := []int{7, 12, 1, 0, 16, 2}
	loop := 2020
	r := firstPart(input, loop)
	fmt.Printf("%vth %v\n", loop, r)
	input = []int{7, 12, 1, 0, 16, 2}
	loop = 30000000
	r = firstPart(input, loop)
	fmt.Printf("%vth %v\n", loop, r)
}

func firstPart(input []int, loop int) int {
	lastNumber := 0
	memory := make(map[int][]int)
	// init
	fmt.Println("Init ...")
	for i, v := range input {
		turn := i + 1
		lastNumber = v
		memory[lastNumber] = append(memory[lastNumber], turn)
		//fmt.Printf("turn %v -> %v : %v\n", turn, memory, lastNumber)
	}
	// run
	fmt.Println("Run ...")
	for turn := len(input) + 1; turn < loop+1; turn++ {
		if len(memory[lastNumber]) < 2 {
			lastNumber = 0
			memory[lastNumber] = switchMem(memory, lastNumber, turn)
		} else {
			lastNumber = memory[lastNumber][1] - memory[lastNumber][0]
			memory[lastNumber] = switchMem(memory, lastNumber, turn)
		}
		//fmt.Printf("turn %v -> %v : %v\n", turn, memory, lastNumber)
	}
	return lastNumber
}

func switchMem(mem map[int][]int, index int, turn int) []int {
	if len(mem[index]) < 2 {
		return append(mem[index], turn)
	}
	mem[index][1], mem[index][0] = turn, mem[index][1]
	return mem[index]
}
