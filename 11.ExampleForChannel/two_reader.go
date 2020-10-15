package main

import (
	"fmt"
	"time"
)

func reader1(c <-chan int) {
	for {
		v := <-c
		fmt.Printf("reader  1: %d\n", v)
	}
}

func reader2(c <-chan int) {
	for {
		v := <-c
		fmt.Printf("reader  2: %d\n", v)
	}
}

func reader3(c <-chan int) {
	for {
		v := <-c
		fmt.Printf("reader  3: %d\n", v)
	}
}

// 一次channel的給值只會觸發一個個gorountine，幾乎為宣告順序依序執行
func main() {
	c := make(chan int)

	go reader1(c)
	go reader3(c)
	go reader2(c)

	for {
		c <- 1
		time.Sleep(time.Millisecond * 10)
	}
}
