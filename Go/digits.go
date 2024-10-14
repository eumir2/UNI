package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\tInserisci un numero:")
	for scanner.Scan() {
		num = 0
		num, _ = strconv.Atoi(scanner.Text())
		if num == 0 {
			fmt.Println("\tFine")
			break
		}
		ok := ParseNum(num)
		if ok {
			fmt.Println("\tIl numero ha cifre ripetute")
		} else {
			fmt.Println("\tIl numero non ha cifre ripetute")
		}
		fmt.Println("\n\tInserisci un numero")
	}
}

func ParseNum(num int) bool {
	value := strconv.Itoa(num)
	digits := []rune(value)
	for i, dig := range digits {
		for j, dig2 := range digits {
			if dig == dig2 && i != j {
				return true
			}
		}
	}
	return false
}
