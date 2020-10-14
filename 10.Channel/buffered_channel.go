package main

import (
	"fmt"
)

type Car struct {
	Model string
}

func main() {

	c1 := make(chan bool)
	go func() {
		fmt.Println("unbuffered: READ channel out of gorountine.")
		c1 <- true
	}()
	fmt.Println("unbuffered: WRITE channel in gorountine.")
	<-c1
	close(c1)

	fmt.Println("-----------------")

	c2 := make(chan bool)
	go func() {
		fmt.Println("unbuffered: READ channel in gorountine.")
		<-c2
	}()
	fmt.Println("unbuffered: WRITE channel out of gorountine.")
	c2 <- true
	close(c2)

	fmt.Println("-----------------")

	c3 := make(chan bool, 1)
	go func() {
		fmt.Println("buffered: READ channel out of gorountine.")
		c3 <- true
	}()
	fmt.Println("buffered: WRITE channel in gorountine.")
	<-c3

	fmt.Println("-----------------")

	// 使用buffer會使channel不必被讀取出，所以在gorountine裡讀取的channel不一定會執行到
	c4 := make(chan bool, 1)
	go func() {
		fmt.Println("buffered: READ channel in gorountine.")
		<-c4
	}()
	fmt.Println("buffered: WRITE channel out of gorountine.")
	c4 <- true
}
