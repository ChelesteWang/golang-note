package main

import (
	"fmt"
	"time"
)

// 5. 你能准确给出包含 select 语句的代码执行结果吗？

// Sender 用来发送信号
// ch <- v   将 v 发送至信道 ch
// v := <-ch 从 ch 接收值并赋予 v
func Sender(ch chan int, stopCh chan bool) {
	i := 220
	for j := 0; j < 10; j++ {
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- true
}


// 对应代码为 5.go，从结果中可以看到，在所有满足条件的 case 中会随机选一个执行
// select 语句主要用于控制通讯，要么是发送要么是接收，如果没有条件满足，则会一直阻塞（特定情况可以利用这样的特性，但一定要小心），具体语法为：
// 每个 case 都必须是一个通信
// 所有 channel 表达式都会被求值
// 所有被发送的表达式都会被求值
// 如果任意一个通信可以进行，就会执行并忽略其他
// 如果多个 case 都可以运行，select 会随机选出一个执行

func main() {
	// 创建一个发送单个 int 的 channel
	ch := make(chan int)
	var c int
	// 创建一个发送单个 bool 的 channel，用于停止程序
	stopCh := make(chan bool)

	go Sender(ch, stopCh)

	for {
		select {
		case c = <-ch: // 从 ch 中读取
			fmt.Println("Receiver 1", c)
		case s := <-ch: // 从 ch 中读取
			fmt.Println("Receiver 2", s)
		case t := <-ch: // 从 ch 中读取
			fmt.Println("Receiver 3", t)
		case _ = <-stopCh:
			// 代码域
			goto end
		}
	}
end:
}
