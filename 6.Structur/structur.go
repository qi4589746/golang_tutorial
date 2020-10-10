package main

import (
	"fmt"
	"strconv"
)

// 使用type宣告struct
type Person struct {
	Name   string
	Age    int16
	gender string
}

// 將Print()定義給Person的struct
func (p Person) Print() {
	fmt.Println(p)
}

// 格式： -該方法要賦予給的結構-   -方法名稱及輸入-   -輸出格式(無return可省略)-
func (p Person) eat() {
	fmt.Println(p.Name + " is eating...")
}

// 定義具回傳之方法給結構
func (p Person) GetName() string {
	return p.Name
}

func main() {
	p1 := Person{
		Name:   "Tom",
		Age:    24,
		gender: "male",
	}
	p1.Print()
	p1.eat()
	fmt.Println(p1.GetName())

	p2 := Person{
		Name:   "Tom",
		Age:    24,
		gender: "male",
	}

	// ＊＊＊可利用==直接判斷兩結構實例內容是否相同
	fmt.Println("value is equal: " + strconv.FormatBool(p1 == p2)) // true
	// ＊＊＊利用&直接看記憶體位置，可判斷兩實例(Instance)是否相同
	fmt.Println("Instance is equal: " + strconv.FormatBool(&p1 == &p2)) // false
}
