package main

import "fmt"

func swap(v1 *int, v2 *int) {
	var temp int
	temp = *v2
	*v2 = *v1
	*v1 = temp
}

func main() {

	// class 1: 存取記憶體指標，將變數指標assign給其他變數
	a1 := 123
	point := &a1 // 利用&將變數指標讀出
	fmt.Println(point)

	// class 2: 指標之值，利用指標讀取該指標所儲存的值
	valueOfPoint := *point
	fmt.Println(valueOfPoint)

	// class 3: 更改指標之值，利用指標之值直接更改值
	*point = 987
	fmt.Println(a1)

	// class 4: 指標搭配function，直接改變該指標位置的值
	v1 := 3
	v2 := 9
	fmt.Printf("v1: %d, v2: %d\n", v1, v2)
	swap(&v1, &v2)
	fmt.Printf("v1: %d, v2: %d\n", v1, v2)
}
