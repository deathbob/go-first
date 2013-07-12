package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	c := make(chan int)
	n := 100
	var foo [100]string

	for i := 0; i < n; i++ {
		go func(i int) {
			resp, err := http.Get("http://golang.org/")
			if err != nil {
				return
			}
			body, err := ioutil.ReadAll(resp.Body)
			foo[i] = string(body)
			defer resp.Body.Close()
			c <- i
		}(i)
	}

	for i := 0; i < n; i++ {
		fmt.Println(<-c)
	}
//	fmt.Println(foo)

}
