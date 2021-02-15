package main

import "fmt"

// 13. 给你一个数组，你能通过 channel 用 2 个 goroutine 完成求和计算吗？
// 我们可以利用两个 goroutine 分别计算前半和后半部分，并最终加起来
func splitSum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("split sum", sum)
	c <- sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)
	gap := len(numbers) / 2

	go splitSum(numbers[gap:], c)
	go splitSum(numbers[:gap], c)
	part1, part2 := <-c, <-c // 从同道中接收

	fmt.Println("Total Sum", part1+part2)

}
