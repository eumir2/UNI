package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s1, _ := leggi()
	min, max, avg := minimo(s1), massimo(s1), media(s1)
	fmt.Printf("Minimo:%d\nMassimo:%d\nValor medio:%f", min, max, avg)
}

func minimo(s1 []int) int {
	minimo := 10000000000
	for _, x := range s1 {
		if x < minimo {
			minimo = x
		}
	}
	return minimo
}
func massimo(s1 []int) int {
	massimo := -100000000
	for _, x := range s1 {
		if x > massimo {
			massimo = x
		}
	}
	return massimo
}

func media(s1 []int) float64 {
	var somma, count int
	var media float64
	for _, x := range s1 {
		count++
		somma += x
	}
	media = float64(somma) / float64(count)
	return media
}
func leggi() (s1 []int, err error) {
	var num int
	for i := 1; i < len(os.Args); i++ {
		num, err = strconv.Atoi(os.Args[i])
		if err == nil {
			minimo(s1)
			massimo(s1)
			media(s1)
			s1 = append(s1, num)
		}
	}
	return s1, err
}
