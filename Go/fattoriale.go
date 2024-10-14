package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num int
	num, _ = strconv.Atoi(os.Args[1])
	fattoriali := Fattoriali(num)
	fmt.Printf("%d", fattoriali)
}
func Fattoriali(n int) (f []int) {
	var num int
	num = 1
	for i := 1; i <= n; i++ {
		num *= i
		f = append(f, num)
	}
	return f
}
