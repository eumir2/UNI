package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num int
	var somma int
	var s1 []int
	for i := 0; i < len(os.Args); i++ {
		num, _ = strconv.Atoi(os.Args[i])
		s1 = append(s1, num)
	}
	somma = Calcola(s1)
	fmt.Println(somma)
}
func Calcola(s1 []int) (sum int) {
	for k, x := range s1[1:] {
		for _, y := range s1[k+1:] {
			if x != y && (x*y)%2 == 0 {
				sum += (x * y)
				fmt.Printf("Somma: %d di %d*%d\n", sum, x, y)
			}
		}
	}
	return sum
}
