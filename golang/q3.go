package main

import "fmt"

// 如何通过 switch 语句来判断变量类型？

func checkType(v interface{}) {
	switch i := v.(type) {
	case nil:
		fmt.Printf("v 的类型 :%T", i)
	case int:
		fmt.Printf("v 是 int 型")
	case float64:
		fmt.Printf("v 是 float64 型")
	case func(int) float64:
		fmt.Printf("v 是 func(int) 型")
	case bool:
		fmt.Printf("v 是 bool 型")
	case string:
		fmt.Printf("v 是 string 型")
	default:
		fmt.Printf("未知型")
	}
	fmt.Println()
}

func main() {
	v1 := 123
	v2 := "123"
	var v3 interface{}

	checkType(v1)
	// int
	checkType(v2)
	// string
	checkType(v3)
	// <nil>
}
