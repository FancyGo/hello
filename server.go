package main

import (
	"fmt"
	"time"

	"github.com/fancygo/hello/api"
)

type Hello struct {
	a []int
}

func t(h Hello) {
	h.a[0] = 2
}

func main() {
	fmt.Println("ver:2")
	h := Hello{
		a: make([]int, 0, 0),
	}
	h.a = append(h.a, 1)
	fmt.Println(h.a)
	t(h)
	fmt.Println(h.a)
	api.HelloWorld()
	time.Sleep(60 * time.Second)
}
