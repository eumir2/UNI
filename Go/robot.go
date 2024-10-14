package main

import (
  ."fmt"
  "bufio"
  "os"
  s"strings"
  "strconv"
  _"unicode"
)

const SU rune = 11014
const DESTRA rune = 10145
const GIU rune = 11015
const SINISTRA rune = 11013
const QUAD rune = 9633

type robot struct{
  x,y,dir int
}

func costruisci(x, y int, dir string)(newR robot,ok bool){
  newR.x = x
  newR.y = y
  switch dir {
  case "su":
    newR.dir = 0
    ok = true
  case "giù":
    newR.dir = 2
    ok = true
  case "destra":
    newR.dir = 1
    ok = true
  case "sinistra":
    newR.dir = 3
    ok = true
  default :
    newR.dir = -1
    ok = false
  }

  return
}

func stato(r robot){
  Printf("%d %d ",r.x,r.y)
  switch r.dir{
  case 0:
    Printf("su \n")
  case 1:
    Printf("destra \n")
  case 2:
    Printf("giù \n")
  case 3:
    Printf("sinistra \n")
  }
  return
}

func avanza(r *robot){
  switch r.dir{
  case 0:
    r.y--
  case 2:
    r.y++
  case 1:
    r.x++
  case 3:
    r.x--
  }
  return
}

func ruota(r *robot, val bool){
  switch val {
  case true: r.dir++
  case false: r.dir--
  }
  return
}

func griglia(rs []robot, xmin,xmax,ymin,ymax int){
  stampo:=0
  for i:=ymin;i<=ymax;i++{
    for j:=xmin;j<=xmax;j++{
      myloop:
      for k:=0;k<len(rs);k++{
        if j == rs[k].x && i == rs[k].y{
          switch rs[k].dir{
          case 0:
            stampo=1
            break myloop
          case 1:
            stampo=2
            break myloop
          case 2:
            stampo=3
            break myloop
          case 3:
            stampo=4
            break myloop
          }
        }else{
          stampo = 0
        }
      }
      switch stampo{
        case 0:
          Print(string(QUAD))
        case 1:
          Print(string(SU))
        case 2:
          Print(string(DESTRA))
        case 3:
          Print(string(GIU))
        case 4:
          Print(string(SINISTRA))
      }
    }
    Println()

  }
  return
}

func vista(rs []robot)(cpp []robot){
  for i:=0;i<len(rs);i++{
    for j:=i+1;j<len(rs);j++{
      if rs[i].x == rs[j].x || rs[i].y==rs[j].y{
        if rs[i].dir - 2 == rs[j].dir || rs[i].dir + 2 == rs[j]. dir{
        cpp=append(cpp,rs[i],rs[j])
        }
      }
    }
  }
  return
}



func main(){

/*  r1, _ := costruisci(5,6,"destra")
  r2, _ := costruisci(3,1,"sinistra")
  r3, _ := costruisci(1,1,"destra")
  r4, _ := costruisci(7,2,"sinistra")
  r5, _ := costruisci(2,2,"destra")
  robots:=[]robot{r1,r2,r3,r4,r5}
  griglia(robots, 0, 10, 0, 10)
  cp:=vista(robots)
  Println(cp)
/*  if ok{
    Println(r2d2)
  }else{
    Printf("il robot %v non è in formato corretto", r2d2)
  }*/

/*  stato(r2d2)
  avanza(&r2d2)
  stato(r2d2)
  ruota(&r2d2, false)
  stato(r2d2)*/
var x,y int
var dir string
var robottini []robot

  myloop2:
  for {
    Println("inserisci un robot")
    new := bufio.NewScanner(os.Stdin)
    new.Scan()
    riga := new.Text()
    input:=s.Fields(riga)
    for i,n := range input{
      switch i{
      case 0:
          x,_=strconv.Atoi(n)
      case 1:
          y,_=strconv.Atoi(n)
      }
      if i==2{
        switch n{
        case "su","giù","destra","sinistra":
          dir = n
        default:
              break myloop2
        }
      }
    }
    var k int
    for k=0;k<len(robottini);k++{
      if robottini[k].x == x && robottini[k].y == y{
        Println("occupato")
        break
      }
    }
    if k<len(robottini){
      continue
    }else{
      r,_:=costruisci(x,y,dir)
      robottini=append(robottini,r)
    }
  }
  Println("robot inseriti: ")
  for _,x:=range robottini{
    stato(x)
  }
}
