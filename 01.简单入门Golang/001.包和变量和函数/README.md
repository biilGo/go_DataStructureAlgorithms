# 包、变量和函数
golang语言只有小括号和大括号,不需要使用逗号来分割代码,只有一种循环`for`

## 工程管理:包机制
每一个大型的软件工程项目,都需要进行工程管理.工程管理的一个环节就是代码层次的管理.

包,也称为库,如代码的一个包,代码的一个库,英文:Library或者Package.比如:我们常常听到某程序员说:嘿.X哥,我知道github上有一个更好用的数据加密库,几千颗星

在高级语言层次,也就是代码本身,各种语言发明了包机制来更好的管理代码,将代码按功能分类归属于不同的包.

golang语言目前的包管理新机制叫`go mod`,我们的项目结构是:
```
├── diy
│   └── diy.go
└── main.go
```

每一个`*.go`源码文件,必须属于一个包,假设包名叫`diy`,在代码最顶端必须有`package diy`,在此之前不能有其他代码片段,如`diy/diy.go`文件中:
```
// 包名
package diy
// 结构体
type Diy struct {
    A int64   // 大写导出成员
    b float64 // 小写不可以导出
}
```

作为执行入口的源码,则强制包名必须为`main`,入口函数为`func main()`,如`main.go`文件中:
```
// Golang程序入口的包名必须为 main
package main // import "golang"
// 导入其他地方的包，包通过 go mod 机制寻找
import (
    "fmt"
    "golang/diy"
)
```

在入口文件`main.go`文件夹下执行以下命令:`go mod init`

该命令会解析`main.go`文件的第一行`package main // import "golang`,注意注释`//`后面的`import "golang"`会生成`go.mod`文件:
```
module golang
go 1.13
```

golang编译器会将这个项目认为是包golang,这是整个项目最上层的包,而底下的文件夹diy作为package diy,包名全路径是`golang/diy`

接着,main.go为了导入包,使用import():
```
// 导入其他地方的包，包通过 go mod 机制寻找
import (
    "fmt"
    "golang/diy"
)
```

可以看到导入了官方的包和我们自己定义的包,官方的包会自动寻找到,不需要任何额外的处理,而自己的包会在当前项目往下找.

在包diy中,我们定义了一个结构体和函数:
```
// 结构体
type Diy struct {
    A int64   // 大写导出成员
    b float64 // 小写不可以导出
}
// 小写函数，不能导出，只能在同一包下使用
func sum(a, b int64) int64 {
    return a + b
}
```

对于包中大小写函数或者结构体中小写的字段,不能导出,其他包不能使用它,golang用它实现了私有或公有控制,毕竟有些包的内容我们不想在其他包中被使用,类似Java的private关键字

最后,golang的程序入口统一在包的main中的main函数,执行程序时是从这里开始的
```
package main
import "fmt"
// init函数在main函数之前执行
func init() {
    // 声明并初始化三个值
    var i, j, k = 1, 2, 3
    // 使用格式化包打印
    fmt.Println("init hello world")
    fmt.Println(i, j, k)
}
// 程序入口必须为 main 函数
func main() {
}
```

有个必须注意的是函数`init()`会在每个包被导入之前执行,如果导入了多个包,那么会根据包导入的顺序先后执行init(),再回到执行函数main()

## 变量
Golang语言可以先声明变量,再赋值,也可以直接创建一个带值的变量.
```
// 声明并初始化三个值
var i, j, k = 1, 2, 3
// 声明后再赋值
var i int64
i = 3
// 直接赋值，创建一个新的变量
j := 5
```

可以看到`var i int64`数据类型是在变量的后面而不是前面,这是golang语言与其他语言最大的区别之一.

同时,作为一门静态语言,golang在编译前还会检查哪些变量和包未被引用,强制禁止游离的变量和包,从而避免某些人类低级错误.如:
```
package main
func main(){
    a := 2
}
```

变量定义后,如果没有赋值,那么存在默认值,我们也可以定义常量,只需加关键字`const`,如:`const s = 2`

常量一旦定义就不能修改