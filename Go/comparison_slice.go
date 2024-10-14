package main

import (
	"fmt"
	"reflect"
)

func main() {
	//WE CANT USE FOR SLICE A==B ,UNLIKE ARRAYS,SO ALSO THE DEEP EQAULITY IS PROBLEMATIC
	s1 := []int{0, 1, 2}
	s2 := []int{0, 1, 2}
	s3 := []int{1, 2, 3}
	ok1 := equal(s1, s2)
	ok2 := equal(s2, s3)
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	fmt.Println("s3: ", s3)
	if ok1 {
		fmt.Println("s1 and s2 are equal")
	}
	if ok2 {
		fmt.Println("s2 and s3 are equal")
	} else {
		fmt.Println("s2 and s3 are not equal ")
	}
	fmt.Println("METODO 2")
	check1 := reflect.DeepEqual(s1, s2)
	check2 := reflect.DeepEqual(s2, s3)

	if check1 {
		fmt.Println("s1 and s2 are equal")
	}
	if check2 {
		fmt.Println("s2 and s3 are equal")
	} else {
		fmt.Println("s2 and s3 are not equal ")
	}

	//THE ONLY EQUALITY IS WITH NIL
	s4 := []int(nil)
	if s4 == nil {
		fmt.Println("The slice is empty")
	}
	//THE TEST "IF A SLICE IS NIL", IS IF LEN(S)==0 NOT S==NI
}

func equal(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
