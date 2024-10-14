package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num int
	var s1 []int
	var min, max int
	var avg float64
	var err error
	for i := 1; i < len(os.Args); i++ {
		num, err = strconv.Atoi(os.Args[i])
		if err == nil {
			fmt.Println(num)
			s1 = append(s1, num)
			min = minimo(s1)
			max = massimo(s1)
			avg = media(s1)
		}
	}
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
