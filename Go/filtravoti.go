package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s1, suff, insuff []int
	s1 = LeggiNumeri()
	suff, insuff = filtra(s1)
	fmt.Println("VOTI SUFF:", suff, "VOTI INSUFF", insuff)
}
func LeggiNumeri() []int {
	var s1 []int
	for i := 0; i < len(os.Args); i++ {
		num, _ := strconv.Atoi(os.Args[i])
		s1 = append(s1, num)
	}
	return s1
}

func filtra(s1 []int) ([]int, []int) {
	var suff, inf []int
	for _, x := range s1 {
		if x >= 60 {
			suff = append(suff, x)
		} else {
			inf = append(inf, x)
		}
	}
	return suff, inf
}
