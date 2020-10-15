package main

import (
	"fmt"
	"os"
	"strconv"
)

// 使用select
func Select(c chan int, quit chan struct{}) {
	for {
		// v := <-c
		// fmt.Println(v)
		select {
		case v := <-c:
			fmt.Println("Received c chan: " + strconv.Itoa(v))
		case <-quit:
			fmt.Println("Received quit chan")
			os.Exit(1)
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan struct{})

	go Select(c, quit)

	go func() {
		c <- 1
		quit <- struct{}{}
	}()
	for {
	}
}
