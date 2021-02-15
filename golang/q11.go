package main

import "fmt"

// 11. 你能说出上面类型转换的最终结果吗？
// 答案是 17 17 17，并没有四舍五入
func main() {
	oneInt := 17
	oneFloat := 17.2

	int2Float := float32(oneInt)
	// 17
	float2Int := int(oneFloat)
	// 17

	fmt.Println(int2Float)
	fmt.Println(float2Int)
	oneFloat = 17.6
	float2Int = int(oneFloat)
	// 17
	fmt.Println(float2Int)
}
