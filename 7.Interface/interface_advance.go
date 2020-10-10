package main

import "fmt"

// 可使用interface{}作為輸入的型態，可接受所有型態資料，"類似"其他語言之泛型
func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	// 輸入各種型態資料至Anything
	Anything(true)
	Anything(5.68)
	Anything("Golang")
	Anything(struct{}{})

	// 利用interface，達到map可儲存各種型態資料的技巧
	mymap := make(map[string]interface{})
	mymap["name"] = "Jeff"
	mymap["age"] = 21
	fmt.Println(mymap)

	// mymap2 := make(map[string]string)
	// mymap2["age"] = 22 // 因為定義為string型態，會無法給予整數
}
