package main

import "fmt"

// 2. 能否准确给出包含 iota 的代码的执行结果？

// 在 golang 中，一个方便的习惯就是使用 iota 标示符，它简化了常量用于增长数字的定义，给以上相同的值以准确的分类。
// 使用 - 跳过 使用赋值打断 


func main() {
	const (
		a = iota
		b
		c
		d = "wdxtub"
		e
		f
		g = 10086
		h
		i
		j = iota
		k
		l
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l)
	// 0 1 2 wdxtub wdxtub wdxtub 10086 10086 10086 9 10 11

	const (
		z = 1 << iota
		y
		x
		w = 2 << iota
		v
		u
	)
	fmt.Println(z, y, x, w, v, u)
	// 1 2 4 16 32 64 分别对应 1 << 0, 1 << 1, 1 << 2, 2 << 3, 2 << 4, 2 << 5
}
