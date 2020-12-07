package file

import (
	"bufio"
	"log"
	"os"
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
