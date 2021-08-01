package main
import (
"fmt"
"net/http"
"io/ioutil"
)

var urls = []string{
  "http://www.google.com/",
  "http://golang.org/",
  "http://blog.golang.org/",
}

func main() {
  // Executes an HTTP HEAD request for all URL's
  // and returns the HTTP status string or an error string.
  for _, url := range urls {
    resp, err := http.Get(url)
    if err != nil {
    fmt.Println("Error:", url, err)
  }
  data, err2 := ioutil.ReadAll(resp.Body)
  if err2 != nil {
    fmt.Println("Error:", url, err)
  }
  fmt.Printf("Got: %q\n", string(data))
  fmt.Println("\n\n\n",url, ": ", resp.Status,"\n\n\n")
  }
}