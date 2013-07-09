package main

import (
  "fmt"
//  "time"
  "net/http"
//  "io/ioutil"
)

func main() {
  c := make(chan int)
  n := 100

  for i := 0; i < n; i++ {
    go func(i int) {
//      time.Sleep(time.Duration(1) * time.Second)
      resp, err := http.Get("http://golang.org/")
      if err != nil {
        return
      }
      defer resp.Body.Close()
//      body, err := ioutil.ReadAll(resp.Body)
      c <- i
    }(i)
  }

  for i := 0; i < n; i++ {
    fmt.Println(<-c)
  }
}