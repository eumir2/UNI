package main

import "fmt"

func main() {
	//TO REMOVE AN ELEMENT
	//CASE 1 AN INT
	s1 := []int{0, 1, 2}
	fmt.Println("BEFORE REMOVE: ", s1)
	fmt.Println("REMOVING: ", s1[1])
	s1 = removeInt(s1, 1)
	fmt.Println("AFTER REMOVE: ", s1)
	//CASE 2 A STRING
	s2 := []string{"bobbo", "babaa", "bebebe"}
	fmt.Println("BEFORE REMOVE: ", s2)
	fmt.Println("Removing:", s2[1])
	s2 = removeString(s2, 1)

	fmt.Println("AFTER REMOVE: ", s2)
}

func removeInt(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func removeString(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
