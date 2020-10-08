package main

import "fmt"

func main() {

	/* 單一變數情況 */

	// class 1: 基本寫法，使用var宣告變數，以預留記憶體空間
	var v1 string
	v1 = "this is v1"
	fmt.Println(v1)

	// class 2: 使用var告變數，並直接給值（此寫法編譯器可自行判斷出變數型態）
	var v2 = "this is v2"
	fmt.Println(v2)

	// class 3(error): 不使用var預留記憶體，則會出錯（無法預留記憶體，編譯器也無法知道變數型態）
	// v3 = "this is v3"
	// fmt.Println(v3)

	// class 4: 不使用var宣告變數，使用:=讓編譯器可自行判斷變數型態
	v4 := "this is v4"
	fmt.Println(v4)

	/* 多變數賦值情況 */

	// class 5: 基本寫法，兩者都先以var宣告並指定型態
	var v5_1 int32
	var v5_2 int32
	v5_1 = 1
	v5_2 = 2
	fmt.Printf("v5_1: %d, v5_2: %d\n", v5_1, v5_2)

	// class 6: 集合寫法，將多個變數以一個var做宣告，並指定各個型態
	var (
		v5_3 int32
		v5_4 int32
	)
	v5_3 = 3
	v5_4 = 4
	fmt.Printf("v5_3: %d, v5_4: %d\n", v5_3, v5_4)

	// class 7: 集合寫法, 將多個變數以一個var做宣告，並直接給值
	var (
		v5_5 = 5
		v5_6 = 6
	)
	fmt.Printf("v5_5: %d, v5_6: %d\n", v5_5, v5_6)

	// class 8: 偷懶寫法，若並非獨立存值，依賴其他變數的話，可透過:=來判斷變數型態
	var v5_7 = 7
	v5_8 := v5_7 + 1
	fmt.Printf("v5_7: %d, v5_8: %d\n", v5_7, v5_8)
}
