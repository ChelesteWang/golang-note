package main

import "fmt"

// 如果 switch 后没有表达式，具体执行哪个 case？
// golang 中 switch case 语句自带 break 使用 fallthrough 强制执行下面语句  
// fallthrough 不能用在 switch 的最后一个分支。
func main() {
	switch {
	case false:
		fmt.Println("1 - false")
		fallthrough
	case true:
		fmt.Println("2 - true")
		fallthrough
	case false:
		fmt.Println("3 - false")
		fallthrough
	case true:
		fmt.Println("4 - true")
	case false:
		fmt.Println("5 - false")
		fallthrough
	default:
		fmt.Println("6 - default")
	}

}
