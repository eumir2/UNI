package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s []int
	for i := 1; i < len(os.Args); i++ {
		value, _ := strconv.Atoi(os.Args[i])
		s = append(s, value)
	}
	countEven, countOdd := Count(s)
	fmt.Println("\tdispari: ", countOdd, ", pari: ", countEven)
}

func Count(s []int) (even int, odd int) {
	for _, elem := range s {
		if elem%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return even, odd
}
