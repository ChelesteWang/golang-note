package main

import (
	"fmt"
	// "time"
)
// 12. 你能给出下面并发代码的执行结果吗？
// 答案是什么都不输出，因为需主进程在 goroutine 执行之前就结束了，所以我们一般会通过 channel 来控制并发！
func gossip(s string) {
	fmt.Println("[Gossip]", s)
}

func main() {
	go gossip("wdxtub")

	// 如果想要看到结果，需要让主进程等待
	// time.Sleep(100 * time.Millisecond)
}
