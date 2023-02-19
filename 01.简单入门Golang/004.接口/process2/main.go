// 接口结构

package main

import (
	"fmt"
	"reflect"
)

// 先定义2个接口类型
// A定义一个接口,有一个方法
type A interface {
	Println()
}

// B定义一个接口,有2个方法
type B interface {
	Println()
	Printf() int
}

// 接口A和B是一种抽象的结构,每个接口都有一些方法在里面,只要结构体struct实现了这些方法,那么这些结构体都是这种接口的类型
// A1Instance定义一个结构体
type A1Instance struct {
	Data string
}

// Println结构体实现了println方法,现在它是一个A接口
func (a1 *A1Instance) Println() {
	fmt.Println("a1", a1.Data)
}

// A2Instance定义一个结构体
type A2Instance struct {
	Data string
}

// Println结构体实现了Println()方法,现在它是一个A接口
func (a2 *A2Instance) Println() {
	fmt.Println("a2:", a2.Data)
}

// printf结构体实现了Printf()方法,现在它是一个B接口,它既是A又是B接口
func (a2 *A2Instance) Printf() int {
	fmt.Println("a2:", a2.Data)
	return 0
}

func main() {

	// 我们要求结构体必须实现某些方法,所以可以先定义一个接口类型的变量,然后将结构体赋值给它
	// 定义一个A接口类型的变量
	var a A

	// 将具体的结构体赋予该变量
	a = &A1Instance{Data: "i love you"} //指针类型的对象实现了该接口,而结构体类型的结构没有实现该接口

	// 调用接口的方法
	a.Println()

	// 只有&A1Instance实现了Println接口,而A1Instance没有实现该接口,使用断言和反射判断接口类型是属于哪个实际的结构体
	// 断言类型
	if v, ok := a.(*A1Instance); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not a A1")
	}

	fmt.Println(reflect.TypeOf(a).String())

	// 将具体的结构体赋予该变量
	a = &A2Instance{Data: "i love you"}

	// 调用接口的方法
	a.Println()

	// 断言类型
	if v, ok := a.(*A1Instance); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not a A1")
	}

	fmt.Println(reflect.TypeOf(a).String())

	// 定义一个B接口类型的变量
	var b B

	b = &A2Instance{Data: "i love you"}

	fmt.Println(b.Printf())
}
