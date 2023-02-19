# 流程控制语句
计算机编程语言中,流程控制语句很重要,可以让机器知道什么时候做什么事做几次,主要有条件和循环语句

go只有一种循环`for`,只有一种判断`if`,还有一种特殊的`switch`条件选择语句

## 条件语句
举个例子:
```
    // 判断语句
    if a > 0 {
        fmt.Println("a>0")
    } else {
        fmt.Println("a<=0")
    }
```



当`a>0`时打印`a>0`,否则打印`a<=0`其中条件`a>0`不需要加小括号

条件语句表示如果什么,做什么,否则做什么

几种判断形式为:
```
if a > 0{

}
```

只有`if`
```
if a > 0{

}else{

}
```

有`if`以及`else`
```
if a > 0{

}else if a == 0 {

}else{

}
```

中间可混入`else if`

如果中间的条件太多的话,可以使用`switch`条件语句
```
    num := 4
    switch num {
    case 3:
        fmt.Println(3)
    case 4:
        fmt.Println(4)
    case 5:
        fmt.Println(5)
    default:
        fmt.Println("not found")
```

这种语句会从`case`一个个判断,如果找到一个`case`符合条件,那么进入该`case`执行指令,否则进入`default`

## 循环语句
循环语句:
```
    // 循环语句
    for i := 9; i <= 10; i++ {
        fmt.Printf("i=%d\n", i)
    }
```

其中i是局部变量,循环第一次前被赋予了值9,然后判断是否满足`i<=10`条件,如果满足那么进入循环打印,每一次循环后会加1,也就是`i++`,然后继续判断是否满足条件

形式为:`    for 起始状态; 进入循环需满足的条件; 每次循环后执行的指令 {}`

也可以死循环
```
    // 死循环语句
    a = 0
    for {
        if a >= 10 {
            fmt.Println("out")
            // 退出循环
            break
        }
        a = a + 1
        if a > 5 {
            continue
        } 
        fmt.Println(a)
    }
```

死循环直接`for{}`,后面不需要加条件,然后当`a>=10`时跳出循环可以使用`break`,表示跳出`for{}`,对于`a>5`我们不想打印出值,可以使用`continue`跳过后面的语句`fmt.Println(a)`,提前再一次进入循环.

切片和字典都可以使用循环来遍历数据
```
    e := []int64{1, 2, 3}                 // slice
    f := map[string]int64{"a": 3, "b": 4} // map

    // 循环切片
    for k, v := range e {
        fmt.Println(k, v)
    }

    // 循环map
    for k, v := range f {
        fmt.Println(k, v)
    }
```

切片遍历出来的结果为:数据下标,数据,字典遍历出来的结果为:数据的键,数据的值

for循环中,对引用类型的临时变量,不能够修改,容易出现临时变量挂失