package main
import "fmt"

/*Scrivere un programma quadrati.go che legge da standard input
un intero n positivo e calcola, utilizzando solo variabili di tipo int,
il massimo quadrato (k^2) <= n.*/

func main()  {
    var n int

    fmt.Print("n: ")
    fmt.Scan(&n)
    c := 0

    for i := 0; i * i <= n; i++{
      c = i * i
    }
    fmt.Println(c)

}
