package main
import "fmt"

type Punto struct{
  x,y float64
}

func newPunto(x,y float64)Punto{
  var p Punto
  p.x = x
  p.y = y
  return p
}

func specchiaPunto(p *Punto){
  (*p).x = -(*p).x
}

func funzUno(p1,p2 Punto) Punto {
  var p Punto
  p.x = (p1.x + p2.x) / 2
  p.y = (p1.y + p2.y) / 2

  return p
}

func funzDue(p *Punto, horiz,vert float64)  {
  (*p).x += horiz
  (*p).y += vert
}

func main() {
var p1, p2, p3 Punto

  fmt.Print("x e y di P1: ")
  fmt.Scan(&p1.x , &p1.y)

  fmt.Print("x e y di P2: ")
  fmt.Scan(&p2.x , &p2.y)

  fmt.Println(funzUno(p1,p2))

  funzDue(&p2, 2.0, 3.0)
  fmt.Println(p2)

  p3 = newPunto(3.0 , 4.0)

  specchiaPunto(&p3)

  fmt.Println(p3)

}
