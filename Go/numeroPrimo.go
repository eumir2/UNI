package main
import "fmt"

//programma che controlla se un numero è primo

func main()  {
  var x,c int

  fmt.Print("Inserire il numero da controllare: ")
  fmt.Scan(&x)

  for i := 2; i < x; i++{
    if x % i == 0{
      c++; //flag
      break
    }
  }

  if c != 0{
    fmt.Println("Non primo")
  }else{
    fmt.Println("Primo")
  }
}
