package main
import (
"fmt"
"time"
)

// Please look at channels in golang also. This can help to communcate bw two goroutines

func main() {
  fmt.Println("In main()")
  t1:=time.Now()
  go longWait(t1)                   // This go keyword will run this function synchronously
  go shortWait(t1)
//   longWait(t1)                   // you can try this two by uncommenting to see the change
//   shortWait(t1)
  fmt.Println("About to sleep in main()")
  time.Sleep(10 * 1e9) // sleep works with a Duration in nanoseconds (ns) !
  fmt.Println("At the end of main()")
  t2:=time.Now()
  fmt.Println(t2.Sub(t1))
}

func longWait(t1 time.Time) {
  fmt.Println("Beginning longWait()")
  time.Sleep(5 * 1e9) // sleep for 5 seconds
  t2:=time.Now()
  fmt.Println(t2.Sub(t1))
  fmt.Println("End of longWait()")
}

func shortWait(t1 time.Time) {
  fmt.Println("Beginning shortWait()")
  time.Sleep(2 * 1e9) // sleep for 2 seconds
  t2:=time.Now()
  fmt.Println(t2.Sub(t1))
  fmt.Println("End of shortWait()")
}