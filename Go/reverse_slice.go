package main

import (
	"fmt"
	"reflect"
)

func main() {
	//REVERSE SLICE
	var sl = []int{1, 2, 3, 4}
	reverse(sl)
	fmt.Println(sl)
	fmt.Println("METODO 2")
	reverseAny(sl)
	fmt.Println(sl)
}

func reverse(sl []int) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}

func reverseAny(sl interface{}) {
	n := reflect.ValueOf(sl).Len()
	swap := reflect.Swapper(sl)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
