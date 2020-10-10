package main

import "fmt"

func main() {

	// class 1: if else
	t := true
	condi := &t

	if condi == nil {
		fmt.Println("Value is nil")
	} else if *condi {
		fmt.Println("Value is true")
	} else {
		fmt.Println("Value is false")
	}

	// class 2: for & foreach
	for i := 0; i <= 5; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	values := []int32{01, 3, 11, 74}
	for index, v := range values {
		fmt.Printf("index %d: %d\n", index, v)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "momo"
	mymap["size"] = 25
	for k, v := range mymap {
		fmt.Printf("key is %s: %v\n", k, v)
	}

	// class 3: break continue
	for i := 0; i < 20; i++ {
		if i%5 == 0 {
			fmt.Println(i)
		} else if i%19 == 0 {
			fmt.Println("Break!!")
			break
		} else {
			continue
		}
	}

	// class 4: switch case
	day := "Fri"
	switch day {
	case "Fri":
		fmt.Println("abc")
	case "Mon", "Tue", "Wed":
		fmt.Println("def")
	default:
		fmt.Println("default")
	}

	// 類似if/else的寫法
	switch {
	case day == "Fri":
		fmt.Println("123")
	case day == "Mon" || day == "Tue" || day == "Wed":
		fmt.Println("456")
	default:
		fmt.Println("default")
	}
}
