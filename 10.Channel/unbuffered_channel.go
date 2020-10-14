package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 建立chan
	c := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		// 使用 <- 賦予channel值
		c <- 1
	}()
	// 使用 <- 將channel的值抓出
	val := <-c
	fmt.Printf("val: %v\n", val)

	go func() {
		time.Sleep(time.Second * 1)
		c <- 2
	}()

	// 會等待channel的值進來
	val = <-c
	fmt.Printf("val: %v\n", val)
	fmt.Println(c)

	fmt.Println("--------------")
	c1 := make(chan int)

	go func() {
		c1 <- 1
		c1 <- 2
		c1 <- 3
		c1 <- 4
		// 使用close關閉channel，避免主程式deadlock
		close(c1)
	}()

	for i := range c1 {
		fmt.Println(i)
	}
}
