// 如何在两个协程间通讯？go提供了一种称为chan的数据类型，叫做信道

package main

import (
	"fmt"
	"time"
)

func Hu(ch chan int) {
	// 使用睡眠模仿一些耗时
	time.Sleep(2 * time.Second)
	fmt.Println("after 2 second hu!!")

	// 执行语句后，通知主协程已经完成操作
	ch <- 1000 // 发送一个整数到信道
}

func main() {
	// 新建一个没有缓冲的信道
	ch := make(chan int) //创建一个能存取int类型的没有缓存的信道，没有缓冲，意味着往里面发送消息，或者接受消息都会堵塞

	// 将信道传入函数，并开启协程
	go Hu(ch)
	fmt.Println("start hu,wait...")

	// 从空缓冲的信道读取int，将会堵塞，直到有消息到来
	v := <-ch // 接收整数可以使用
	fmt.Println("receiver:", v)
}

// 执行协程后，函数里面会睡眠2分钟，所以2分钟之后信道才会收到消息，在没有收到消息之前v := <-ch会堵塞，直到协程go Hu(ch)完成，那么消息收到，程序结束
