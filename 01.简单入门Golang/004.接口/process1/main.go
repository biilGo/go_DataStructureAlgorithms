// 接口类型

package main

import (
	"fmt"
	"reflect"
)

// 我们可以将函数的参数定为interface和变量的定义一样
func print(i interface{}) {
	fmt.Println("i")
}

func main() {
	// 基本使用
	// 声明一个未知类型的a,表明不知道是什么类型
	var a interface{}

	// 给变量赋值一个整数,此时a仍然是未知类型,使用占位符%T可以打印变量的真实类型,占位符%v打印值
	a = 2
	fmt.Printf("%T,%v\n", a, a) // Printf在内部会进行类型判断

	// 传入函数
	print(a)
	print(3)
	print("i love you")
	// 传给print函数的参数可以是任何类型

	// 具体类型判断
	// 我们定义了interface{}但是实际使用时,我们有判断类型的需求,有两种方法可以进行判断
	// 使用断言判断是否是int数据类型
	v, ok := a.(int)
	if ok {
		fmt.Printf("a is int type,value is %sd\n", v)
	}
	// 直接在变量后面使用.(int)有两个返回值v,ok会返回.ok如果时true表明是整数类型,这个整数会被赋予v,然后我们可以拿v愉快地玩耍了,否则ok为false,v为空值,也就是默认值0

	// 使用断言,判断变量类型
	switch a.(type) {
	case int:
		fmt.Println("a is type int")
	case string:
		fmt.Println("a is type string")
	default:
		fmt.Println("a not type found type")
	}
	// switch中断言不再使用.(具体类型)而是a.(type)

	// 使用反射找出变量类型
	t := reflect.TypeOf(a)
	fmt.Printf("a is type:%s", t.Name())
}
