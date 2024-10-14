package main

import ("fmt"
        "math/rand"
        "time")

func reverse(p *[10]int) {
  j := len(*p)-1
  n := len(*p) / 2
  for i := 0; i < n; i++{
    p[i],p[j] = p[j],p[i]
    j--
  }
}

func switchFirstLast(p *[10]int){
  p[0],p[9] = p[9],p[0]
}

func main()  {
  rand.Seed(time.Now().UnixNano())
  const DIM = 10
  var val[DIM] int

  for i := 0; i < len(val); i++{
    val[i] = rand.Intn(9)+1
  }
  fmt.Println(val)

  reverse(&val)
  fmt.Println(val)

  switchFirstLast(&val)
  fmt.Println(val)


}
