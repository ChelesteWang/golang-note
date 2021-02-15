package main

import (
	"fmt"
)

// 7. 传引用和引用类型一样吗？你能说出下面代码的执行结果吗？

// 第一种交换方式
func swap1(x, y int) {
	x, y = y, x
	fmt.Println("swap1 函数交换", x, y)
}

// 第二种交换方式
func swap2(x, y *int) {
	*x, *y = *y, *x
	fmt.Println("swap1 函数交换", *x, *y)
}

// 修改 map
func modifyMap(item map[string]int) {
	item["a"] = 314
	fmt.Println("modifyMap 函数修改", item)
}

// 修改 切片
func modifySlice(item []int) {
	item[0] = 314
	fmt.Println("modifySlice 函数修改", item)
}

// 修改 数组
func modifyArray(item [3]int) {
	item[0] = 314
	fmt.Println("modifyArray 函数修改", item)
}

// 修改数组指针
func modifyArrayPtr(item *[3]int) {
	item[0] = 314
	fmt.Println("modifyArrayPtr 函数修改", *item)
}

// 切片与字典为指针
func main() {
	x := 314
	y := 220
	fmt.Println("x y 交换前", x, y)
	// 314 220
	swap1(x, y)
	fmt.Println("x y swap1 交换后", x, y)
	// 314 220
	swap2(&x, &y)
	fmt.Println("x y swap2 交换后", x, y)
	// 220 314 

	arraySample := [3]int{1, 2, 3}
	fmt.Println("arraySample 修改前", arraySample)
	// 1 2 3
	modifyArray(arraySample)
	fmt.Println("arraySample 修改后", arraySample)
	// 1 2 3
	modifyArrayPtr(&arraySample)
	fmt.Println("modifyArrayPtr 修改后", arraySample)
	// 314 2 3
	sliceSample := []int{1, 2, 3}
	fmt.Println("sliceSample 修改前", sliceSample)
	// 1 2 3
	modifySlice(sliceSample)
	fmt.Println("sliceSample 修改后", sliceSample)
	// 314 2 3
	mapSample := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("mapSample 修改前", mapSample)
	// "a":1 "b":2 "c":3
	modifyMap(mapSample)
	fmt.Println("mapSample 修改后", mapSample)
	// "a":314 "b":2 "c":3
}
