package main
import "fmt"

//programma che controlla se un numero contiene la cifra 3

func main()  {
  var n int

  fmt.Print("Inserire il numero:")
  fmt.Scan(&n)

  for n > 0 {
    if(n % 10 == 3){
      break
    }
    n /= 10
  }
  if(n > 0){
    fmt.Println("Il numero contiene la cifra 3")
  }else{
    fmt.Println("Il numero non contiene la cifra 3")
  }
}
