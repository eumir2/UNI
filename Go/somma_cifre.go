package main
import "fmt"

/*Scrivere un programma somma_cifre.go che calcola la somma delle cifre
di un numero intero positivo fornito da standard input.*/

func main()  {
  var n int
  var somma int

  fmt.Scan(&n)

	for n != 0 {
	  tmp := n % 10
    somma += tmp
		n /= 10
	}

  fmt.Println(somma)
}
