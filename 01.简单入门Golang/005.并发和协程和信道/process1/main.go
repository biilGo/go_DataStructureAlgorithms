// 协程和go关键字
package main

import (
	"fmt"
	"time"
)

func Hu() {
	// 使用睡眠模仿一些耗时
	time.Sleep(2 * time.Second)
	fmt.Println("after 2 second hu!!")
}

func main() { //main函数本身作为程序的主协程,如果main函数结束的话,其他协程也会死掉,必须使用死循环来避免主协程终止
	// 将会杜塞
	// Hu() //如果直接使用Hu()函数,因为函数内部使用time.Sleep进行睡眠,需要等待2秒,所以程序会堵塞

	// 开启新的协程,不会堵塞
	// go开启一个新的协程,不再堵塞,执行完毕后,马上直接执行后续的语句
	go Hu()

	fmt.Println("start hu, wait...")

	// 必须死循环,不饶主协程退出了,程序就结束了
	for {
		time.Sleep(1 * time.Second)
	}
}
