package main
import "fmt"


//programma che preso in input un voto in 30esimi, lo converte in 110ecimi

func main()  {
  //varibili
  var voto, votoConvertito int

  fmt.Print("Inserire un voto in 30esimi: ")
  fmt.Scan(&voto)

  //  voto : 30 = votoConvertito : 110
  votoConvertito = (voto * 110) / 30

  fmt.Println("Voto convertito:", votoConvertito)
}
