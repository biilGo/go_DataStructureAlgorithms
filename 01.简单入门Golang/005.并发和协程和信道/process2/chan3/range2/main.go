package main

import "fmt"

func main() {
	buffedChan := make(chan int, 2)
	buffedChan <- 2
	buffedChan <- 3
	close(buffedChan) // 关闭后才能for打印出来，否则死锁

	for i := range buffedChan {
		fmt.Println(i)
	}
}

// 不能多次关闭一个信道，不能往关闭了信道打消息，否则会报错
