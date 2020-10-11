package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 使用waitGroup來控管gorountine
	wg := &sync.WaitGroup{}

	// 設定gorountine數量
	wg.Add(2)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 2)
			fmt.Println("gorountine1")
		}
		// 將一個waitGroup裡的gorountine數量減一
		wg.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 3)
			fmt.Println("gorountine2")
		}
		wg.Done()
	}()

	// 等待gorountine數量被減為零
	wg.Wait()

	fmt.Println("All gorountines are done.")
}
