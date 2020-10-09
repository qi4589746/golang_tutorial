package main

import (
	"fmt"
	"strings"
)

func main() {

	// class 1: 兩字串相加
	s1 := "Hello "
	s2 := "world"
	s3 := s1 + s2
	fmt.Println(s3)

	// class 2: 包含查詢
	s4 := "This is class2"
	s5 := "class2"
	fmt.Println(strings.Contains(s4, s5)) // (主要字串，預查詢字串)

	// class 3: 字串替換
	s6 := "a b c d e"
	fmt.Println(strings.ReplaceAll(s6, " ", "_")) // (原始字串，指定字串，預替換成的字串)，此方法不會覆蓋原字串
	fmt.Println(s6)

	// class 4: 依指定字串切割字串
	s7 := strings.Split(s6, " ") // （原始字串， 預使用切割字串）
	for index, a := range s7 {
		fmt.Printf("ibdex %d: %s\n", index, a)
	}

}
