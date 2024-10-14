package main
import "fmt"

/*Scrivere un programma Go maggiore.go che legga due interi, li salvi in due var min e
max in qualsiasi ordine; se non sono in ordine, li sistemi in modo che min contenga il minore e
max il maggiore; infine stampi il contenuto di max.*/

func main()  {
  var min,max int

  fmt.Print("due int: ")
  fmt.Scan(&min, &max)

  if min < max {

  }else{
    min,max = max,min
  }

  fmt.Println(max)
}
