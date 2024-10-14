package main
import ("fmt"
        "math")

/*: Scrivere un programma Go distanza_punti.go che calcola la distanza tra due punti
nel piano cartesiano.*/

func  main()  {
  var x1, x2, y1, y2 float64

  fmt.Print("x e y del primo punto:")
  fmt.Scan(&x1, &y1)

  fmt.Print("x e y del secondo punto:")
  fmt.Scan(&x2, &y2)


  //La formula della distanza tra due punti P1 e P2 è:
  //radiceQuadrata( (x2 − x1)^2 + (y2 − y1)^2 )

  distanza := math.Sqrt((x2 - x1) * (x2 - x1) + (y2 - y1) * (y2 - y1))

  fmt.Println("La distanza tra i due punti è:",distanza)

}
