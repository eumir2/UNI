package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	N, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[2])
	s := GeneraNumeri(N, n)
	s2 := FiltraNumeri(s)
	for _, x := range s2 {
		fmt.Println(x)
	}
}

func GeneraNumeri(N, k int) []int {
	n2 := strconv.Itoa(N)
	var s []int
	for i := 0; i < len(n2); i++ {
		for j := len(n2); j > i; j-- {
			if len(n2[i:j]) == k {
				value, _ := strconv.Atoi(n2[i:j])
				s = append(s, value)
			}

		}
	}
	return s
}

func FiltraNumeri(sl []int) []int {
	var s []int
	for _, x := range sl {
		if x%2 == 0 {
			s = append(s, x)
		}
	}
	return s
}
