// 設定當前package名稱，不需與資料夾名稱相同
package main

// 載入所需package
import "fmt"

// 程式的進入執行點，與c/c++/java一樣，使用main
func main() {
	fmt.Println("Hello World")
}

// 使用 go build 對本檔案進行編譯，在直接執行，以linux為例：
// $ go build hello_world.go
// $ ./hello_world
