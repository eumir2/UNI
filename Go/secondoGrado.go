package main
import "fmt"
import "math"


//programma che risolve un'equazione di secondo grado

func main()  {
  //variabili
  var a, b, c float64

  fmt.Print("Inserire i coefficenti in ordine (a,b,c): ")
  fmt.Scan( &a, &b, &c)


  //calcolo il delta
  //delta = radice(b^2 - 4ac)
  p1 := (b * b)
  p2 := 4 * (a * c)
  delta := p1 - p2
  fmt.Println("Delta:", delta)

  tmp := math.Sqrt(delta)


  if delta < 0 {

    fmt.Println("Non ci sono soluzioni")

  }else if delta == 0 {
    //calcolo la soluzione
    x1 := ((b * -1) + tmp) / (a * 2)

    //la stampo
    fmt.Println("La soluzione è:", x1)

  }else{
    //calcolo le soluzioni
    x1 := ((b * -1) + tmp) / (a * 2)
    x2 := ((b * -1) - tmp) / (a * 2)

    //le stampo
    fmt.Println("La prima soluzione è:", x1)
    fmt.Println("La seconda soluzione è:", x2)
  }
}
