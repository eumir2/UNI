package main
import "fmt"
import "time"

type Clock struct{
  hour,min,sec int
}

func countDown(c *Clock){
  for(c.hour > 0 || c.min > 0 || c.sec > 0){
    time.Sleep(time.Duration(1) * time.Second)
    c.sec--
    if c.sec < 0{
      c.sec = 59
      changeMin(c)
    }
    fmt.Println(c)
  }
}

func changeMin(c *Clock){
  c.min--
  if c.min < 0{
    c.min = 59
    changeHour(c)
  }
}

func changeHour(c *Clock){
  c.hour--
}

func main(){
  var c Clock

  fmt.Print("time (hh mm ss):")
  fmt.Scan(&c.hour,&c.min,&c.sec)

  countDown(&c)


}
