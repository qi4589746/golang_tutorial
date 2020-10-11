package main

import (
	"fmt"
	"time"
)

// 以func為單位，作為gorountine的執行項目
func myFunc() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("myFunc")
	}
}

func myFunc2() {
	for {
		time.Sleep(time.Second * 3)
		fmt.Println("myFunc2")
	}
}

func main() {
	// 使用 go 直接建立一個gorountine在不同的thread做運算
	go myFunc()
	go myFunc2()

	// 使用匿名函式
	go func() {
		for {
			time.Sleep(time.Second * 2)
			fmt.Println("Anonymous Function")
		}
	}()

	fmt.Println("Start goroutines...")
	// 使用select等待所有gorountines運行結束
	select {}
}
