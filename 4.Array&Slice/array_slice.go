package main

import (
	"fmt"
	"reflect"
)

// Array：固定長度, 不可append元素
// Slice：可動態調整長度，可append元素

func main() {
	// class 1: 宣告陣列(Array)
	myArray := [4]int32{1, 2, 3, 4}
	fmt.Println(myArray)

	// class 2: 更動陣列元素
	myArray[0] = 0
	fmt.Println(myArray)

	// class 3: 依序讀出陣列元素
	for index, v := range myArray {
		fmt.Printf("ibdex %d: %d\n", index, v)
	}

	// class 4: 宣告切片(Slice)
	intArr := []int32{1, 2, 3, 4, 5}
	fmt.Println(intArr)

	// class 5: 切片加入元件
	strArr := []string{"Hi", "my", "name", "is"}
	// 不可append兩切片 (error)
	// s2 := []string{"Tom", "!"}
	// strArr = append(strArr, s2)
	strArr = append(strArr, "Tom", "!")
	fmt.Println(strArr)

	// class 6: 兩陣列比較
	a1 := [4]int32{1, 2, 3, 4}
	a2 := [4]int32{1, 2, 3, 4}
	fmt.Println(a1 == a2) // true

	// class 7: 兩切片比較
	a3 := []int32{1, 2, 3}
	a4 := []int32{1, 2, 3}
	// 切片無法直接進行 ==，切片只允許於nil做判斷
	// fmt.Println(a3 == a4)
	fmt.Println(a3 == nil) // false

	// class 8: 若要比較兩切片是否為相同物件，需比較指標
	fmt.Println(reflect.ValueOf(a3).Pointer() == reflect.ValueOf(a4).Pointer()) // false
	a5 := a4
	fmt.Println(reflect.ValueOf(a4).Pointer() == reflect.ValueOf(a5).Pointer()) // true

	// class 9: 建立特定大小切片，注意是切片不是陣列，只是初始化特定長度的值
	a6 := make([]int, 5)
	fmt.Println(a6)
	fmt.Println(append(a6, 123))

}
