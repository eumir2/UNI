package main

import (
	"fmt"
	"os"
	"strconv"
)

func checksomma0(seq []int) bool {
	sum := 0
	for _, n := range seq {
		sum += n
	}
	if sum == 0 {
		return true
	}
	return false
}

func main() {
	sequenza := make([]int, len(os.Args[1:]))
	for i := 0; i < len(os.Args[1:]); i++ {
		sequenza[i], _ = strconv.Atoi(os.Args[i+1])
	}
	for i := 0; i < len(sequenza); i++ {
		for j := len(sequenza); j > i; j-- {
			if checksomma0(sequenza[i:j]) {
				fmt.Print(sequenza[i:j], "\n")
			}
		}
	}
}
