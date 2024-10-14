package main

import "fmt"

func main() {
	//ROTATE SLICE BY N ELEMENTS
	//APPLY THE REVERSE FUNCITON 3 TIMES:
	//-FIRST TO THE LEADING N ELEMENTS
	//-THE TO THE REMANING ELEMENTS
	//-FINALLY TO THE WHOLE SLICE
	var sl = []int{0, 1, 2, 3, 4, 5}
	//ROATATE S TO 2 POSITION
	rotate(sl[:2])
	rotate(sl[2:])
	rotate(sl)
	fmt.Println(sl)
}

func rotate(sl []int) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}
