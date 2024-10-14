package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num int
	var s1 []int
	var somma int
	for i := 1; i < len(os.Args); i++ {
		num, _ = strconv.Atoi(os.Args[i])
		s1 = append(s1, num)
	}
	fmt.Println(s1)
	if (len(s1))%2 == 0 {
		somma = calcola(s1)
	} else {
		os.Exit(0)
	}
	fmt.Println("LA SOMMA E': ", somma)
}
func calcola(s1 []int) int {
	var somma int
	for i := 0; i < len(s1); i += 2 {
		somma += (s1[i] * s1[i+1])
	}
	return somma
}
