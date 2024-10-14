package main

/*Scrivere un programma che, dopo aver letto da standard input un numero intero n , chiede all'utente di
inserire n numeri interi (sempre da standard input).
Il programma deve stampare gli n numeri interi in ordine inverso rispetto a quello di inserimento.
*/
import "fmt"

func main() {
	var s []int
	var n int
	fmt.Println("Quanti numeri vuoi inserire?")
	fmt.Scan(&n)
	s = make([]int, n)
	fmt.Printf("Inserire i %d numeri:\n", n)
	for i := 0; i < n; i++ {
		var elemento int
		fmt.Scan(&elemento)
		s[i] = elemento
	}
	fmt.Print("Numeri iniziali: ")
	for i := 0; i < n; i++ {
		fmt.Print(s[i], " ")
	}
	fmt.Println()
	fmt.Print("Numeri in ordine inverso: ")
	for i := n - 1; i >= 0; i-- {
		fmt.Print(s[i], " ")
	}

}
