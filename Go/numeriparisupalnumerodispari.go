package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s1 []int
	for _, x := range os.Args[1:] {
		val, err := strconv.Atoi(x)
		if err != nil {
			s1 = append(s1, val)
		}
	}
	fmt.Println(s1)
	minimo := MinimoDispari(s1)
	for _, x := range s1 {
		if x > minimo && x%2 == 0 {
			fmt.Print(x, "")
		}
	}
	fmt.Println("Minimo valore dipspari: ", minimo)

}
func MinimoDispari(sl []int) int {
	var minDispari int
	minDispari = 0
	for _, x := range sl {
		if x%2 != 0 {
			minDispari = x
			break
		}
	}
	for ; ; i++ {

	}
	return minDispari
}
