package conv

import (
	"log"
	"strconv"
)

func toi(in string) int {
	num, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
