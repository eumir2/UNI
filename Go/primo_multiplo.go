package main
import ("fmt"
        "bufio"
        "os"
        "strconv")

/*Scrivi un programma primo_multiplo.go che riceve in input una sequenza di 5 numeri interi >=0 e stampa il primo numero della sequenza che è multiplo di 3, oppure "no" se la sequenza non ne contiene. Caricalo poi su upload
Nota che, se la sequenza contiene un multiplo di 3, non occorre leggere ed elaborare i numeri nella sequenza che vengono dopo tale numero*/

func main()  {
  const DIV = 3
  var c int
  c = 0

  
  i := 1
  for scanner.Scan() && i < 5{
    i++
    l := scanner.Text()
    n,_ := strconv.Atoi(l)
    if n % DIV == 0{
      fmt.Print(n)
      c = 1
      break
    }

  }
  if c == 0{
    fmt.Println("no")
  }

}
