package main
import "fmt"


//programma che stabilisce quale frazione è maggiore

func main()  {
  //variabili
  var n1, d1, n2, d2 int

  fmt.Print("Inserire il numeratore e denominatore della prima frazione: ")
  fmt.Scan(&n1, &d1)

  fmt.Print("Inserire il numeratore e denominatore della seconda frazione: ")
  fmt.Scan(&n2, &d2)

  //se i denominatori sono uguali confronto i numeratori
  if d1 == d2{
    if n1 > n2 {
      fmt.Println("La prima frazione è maggiore")
    }else{
      fmt.Println("La seconda frazione è maggiore")
    }
  }else{
    if d1 < d2 {
      fmt.Println("La prima frazione è maggiore")
    }else{
        fmt.Println("La seconda frazione è maggiore")
    }
  }

}
