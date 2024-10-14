package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s []int
	for i := 1; i < len(os.Args); i++ {
		runes := []rune(os.Args[i])
		if len(runes) == 1 || len(runes) == 2 {
			value, _ := strconv.Atoi(os.Args[i])
			s = append(s, value)
		} else {
			s = append(s, 0)
		}
	}
	if i, ok := Valid(s); ok {
		fmt.Println("Sequenza valida")
	} else {
		fmt.Println("Valore in posizione ", i, " non valida")
	}
}

func Valid(s []int) (int, bool) {
	var sum int
	var ok bool
	for i, x := range s {
		if x < 0 {
			return i, false
		}
		if sum == 0 && i%2 == 0 {
			sum += x
		} else if sum != 0 && i%2 == 0 {
			if x >= sum {
				ok = true
				sum += x
			} else if x != 0 {
				ok = false
				return i, ok
			}
		}
	}
	return 0, true
}
