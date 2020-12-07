package main

import (
	"bufio"
	"fmt"
	"go/types"
	"log"
	"os"
)

type group struct {
	answers []string
}

func (g *group) add(answer string) {
	g.answers = append(g.answers, answer)
}

func (g group) nbYes() int {
	answersYes := make(map[rune]types.Nil)
	for _, answer := range g.answers {
		for _, c := range answer {
			answersYes[c] = types.Nil{}
		}
	}
	return len(answersYes)
}

func (g group) sameYes() int {
	ref := []rune(g.answers[0])
	same := []rune{}
	for i := 1; i < len(g.answers); i++ {
		for _, a := range g.answers[i] {
			for _, r := range ref {
				if a == r {
					same = append(same, r)
					break
				}
			}
		}
		ref = same
		same = []rune{}
	}
	return len(ref)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Open file failed")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	people := []group{}
	p := &group{answers: []string{}}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			people = append(people, *p)
			p = &group{answers: []string{}}
			continue
		}
		p.add(line)
	}
	people = append(people, *p)
	nbYes := firstPart(people)
	fmt.Printf("NbYes %v\n", nbYes)
	intersectYes := secondPart(people)
	fmt.Printf("IntersectYes %v\n", intersectYes)

}

func firstPart(people []group) int {
	sum := 0
	for _, p := range people {
		sum += p.nbYes()
	}
	return sum
}

func secondPart(people []group) int {
	sum := 0
	for _, g := range people {
		v := g.sameYes()
		sum += v
	}
	return sum
}
