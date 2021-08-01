package main

import (
  "fmt"
  "time"
)

// The operation ch <- int1 means that the variable int1 is sent through the channel ch.

// The operation int2 = <- ch means that variable int2 receives data from the channel ch .


func main() {
  buf:=100
  ch := make(chan string,buf)   // channel can keep maximum 100[buf] items
  go sendData(ch) // calling goroutine to send the data
  go getData(ch)  // calling goroutine to receive the data
  time.Sleep(1e9)
}

func sendData(ch chan string) { // sending data to ch channel
  ch <- "Washington"
  ch <- "Tripoli"
  ch <- "London"
  ch <- "Beijing"
  ch <- "Tokyo"
}

func getData(ch chan string) {
  var input string
  for {
    input = <-ch // receiving data sent to ch channel
    fmt.Printf("%s \n", input)
  }
  close(ch) // closed the channel
}