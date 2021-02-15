package main

import (
	"fmt"
)
// 6. 你能准确给出包含 select 语句的代码执行结果吗？(powered by jiasen)
func main() {
	// 创建一个发送单个 int 的 channel
	// ch <- v   将 v 发送至信道 ch
	// v := <-ch 从 ch 接收值并赋予 v
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
