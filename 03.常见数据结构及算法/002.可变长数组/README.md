# 可变长数组
因为数组大小是固定的,当数据元素特别多时,固定的数组无法存储这么多的只,所以可变长数组就出现了,这也是一种数据结构,可变长数组被内置在语言里面:切片slice

slice是对底层数组的抽象和控制,它是一个结构体:
```
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```

1. 指向底层数组的指针.
2. 切片的真正长度,也就是实际元素占用的大小
3. 切片的容量,底层固定数组的长度

每次可以初始化一个固定容量的切片,切片内部维护一个固定大小的数组,当`append`新元素时,固定大小的数组不够时会自动扩容,程序'append1'

我们可以看到go的切片无法原地append,每次添加元素时返回新的引用地址,必须把该引用重新赋予之前的切片变量,并且,当容量不够时,会自动倍数递增扩容,事实上,go在切片长度大于1024后,会以接近于1.25倍进行容量扩容.

## 实现可变长数组
我们来实现一个简单的,存放整数的,可变长的数组版本.

因为go的限制,不允许使用[n]int来创建一个固定大小为n的整数数组,只允许使用常量来创建大小.

所以我们这里会使用切片的部分功能来代替数组,虽然切片本身是可变长数组,但是我们不会用到他的append功能,只把它当成数组用
```
import (
    "sync"
)
// Array 可变长数组
type Array struct {
    array []int      // 固定大小的数组，用满容量和满大小的切片来代替
    len   int        // 真正长度
    cap   int        // 容量
    lock  sync.Mutex // 为了并发安全使用的锁
}
```

### 初始化数组
创建一个len个元素,容量为cap的可变长数组
```
// Make 新建一个可变长数组
func Make(len, cap int) *Array {
    s := new(Array)
    if len > cap {
        panic("len large than cap")
    }
    // 把切片当数组用
    array := make([]int, cap, cap)
    // 元数据
    s.array = array
    s.cap = cap
    s.len = 0
    return s
}
```

主要利用满容量和满大小的切片来充当固定数组,结构体Array里面的字段len和cap来控制值的存取.不允许设置len>cap的可变长数组.

时间复杂度为:O(1),因为分配内存空间和设置几个值是常树时间.

### 添加元素
```
// Append 增加一个元素
func (a *Array) Append(element int) {
    // 并发锁
    a.lock.Lock()
    defer a.lock.Unlock()
    // 大小等于容量，表示没多余位置了
    if a.len == a.cap {
        // 没容量，数组要扩容，扩容到两倍
        newCap := 2 * a.len
        // 如果之前的容量为0，那么新容量为1
        if a.cap == 0 {
            newCap = 1
        }
        newArray := make([]int, newCap, newCap)
        // 把老数组的数据移动到新数组
        for k, v := range a.array {
            newArray[k] = v
        }
        // 替换数组
        a.array = newArray
        a.cap = newCap
    }
    // 把元素放在数组里
    a.array[a.len] = element
    // 真实长度+1
    a.len = a.len + 1
}
```

首先添加一个元素到可变长数组里,会加锁,这样会保证并发安全,然后将值放在数组里:`a.array[a.len]=element`然后`len+1`表示真实大小又多了一个.

当真实大小len=cap时,表明位置都用完了,没有多余的空间放新值,那么会创建一个固定带下2*len的新数组来替换老数组:`a.array=newArray`,当然容量也会变大:`a.cap=newCap`如果一开始设置的容量`cap=0`,那么新的容量会是从1开始.

添加元素中,耗时主要是在老数组中的数据移动到新数组,时间复杂度为:O(n).当然,如果容量够的情况下,时间复杂度会变为:O(1)

如何添加多个元素:
```
// AppendMany 增加多个元素
func (a *Array) AppendMany(element ...int) {
    for _, v := range element {
        a.Append(v)
    }
}
```

只是简单遍历以下,调用append函数,其中`...int`是go语言的特征,表示多个函数变量.

### 获取指定下标元素
```
// Get 获取某个下标的元素
func (a *Array) Get(index int) int {
    // 越界了
    if a.len == 0 || index >= a.len {
        panic("index over len")
    }
    return a.array[index]
}
```

当可变长数组的真实大小为0,或者下标index超出了真实长度len,将会panic越界.因为只获取下标的值,所以时间复杂度为O(1)

### 获取真实长度和容量
```
// Len 返回真实长度
func (a *Array) Len() int {
    return a.len
}
// Cap 返回容量
func (a *Array) Cap() int {
    return a.cap
}
```

时间复杂度为 O(1).

### 示例
程序:example

# 总结
可变长数组在实际开发上,经常会使用到,其实在固定大小数组的基础上,会自动进行容量扩展

因为这一数据结构的使用频率太高了,所以go自动提供了这一数据类型:切片