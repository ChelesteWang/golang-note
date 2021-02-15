package main

import "fmt"

// 1. 没有初始化的变量，默认值是什么？

func main() {
	// 声明一个变量并初始化
	var sampleInitStr = "wdxtub.com"
	fmt.Println(sampleInitStr)
	// wdxtub.com

	// 字符串
	var sampleStr string
	fmt.Println(sampleStr)
	// ""

	// 整型、浮点型初始化
	var sampleInt int
	fmt.Println(sampleInt)
	//0 

	var sampleFloat float32
	fmt.Println(sampleFloat)
	// 0

	// 布尔值初始化
	var sampleBool bool
	fmt.Println(sampleBool)
	// 0

	var t1 *int
	fmt.Println(t1)
	// <nil>

	var t2 []int
	fmt.Println(t2)
	// []

	var t3 map[string]int
	fmt.Println(t3)
	// map[]

	var t4 chan int
	fmt.Println(t4)
	// <nil> 管道类型

	var t5 func(string) int
	fmt.Println(t5)
	// <nil> 函数类型

	var t6 error // error 是接口
	fmt.Println(t6)
	// <nil> 接口类型 error

}
