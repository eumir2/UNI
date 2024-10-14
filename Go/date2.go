package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var line string
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		lines = append(lines, line)
	}
	fmt.Println(lines, "\n")
	for _, x := range lines {
		fmt.Println(x)
		check := daInvertire(x)
		if check == false {
			fmt.Println("Prima inversione ", x)
			x = Inverti(x)
			fmt.Println("Dopo inversione ", x)
		}
		x = strings.Replace(x, `/`, `-`, 3)
	}

}
func daInvertire(s string) bool {
	sinRune := []rune(s)
	if sinRune[2] == 45 || sinRune[3] == 45 {
		return true
	}
	return false
}
func Inverti(s string) string {
	var invertito string
	for i := 0; i < 3; i++ {
		for y := len(s) - 1; s[y] == 45; y-- {
			invertito += string(s[y])
		}
	}
	return invertito
}
