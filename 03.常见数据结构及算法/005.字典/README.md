# 字典
字典是存储键值对的数据结果,把一个健和一个值映射起来,一一映射,健不能重复.在某些教程中,这种结构可能称为符号表,关联数组和映射.

如:
```
键=>值
"cat"=>2
"dog"=>1
"hen"=>3
```

我们拿出键`cat`的值,就是`2`

go提供了这一数据结构:`map`并且要求键的数据类型必须是可比较的,因为如果不可比较,就无法知道键是否存在还是不存在.

go字典的一般操作如下:`map01`

字典的实现有两种方式:哈希表`HashTable`和红黑树`RBTree`.go语言中字典`map`的实现由哈希表实现,具体可参考标准库`runtime`下`map.go`文件.

# 实现不可重复集合set
一般很多编程语言库,会把不可重复集合命名为`Set`,这个`Set`中文直译为集合,在某些上下文条件下,我们大脑自动过滤,集合这词指的是不可重复集合还是指统称的集合.

不可重复集合`Set`存放数据,特点就是没有数据会重复,会去重.你放一个数据进去,再放一个数据进去,如果两个数据一样,那么只会保存一份数据.

集合`Set`可以没有顺序关系,也可以按值排序,算一种特殊的列表

因为我们知道字典的键是不重复的,所以只要我们不考虑字典的值,就可以实现集合,我们来实现存整数的集合`Set`

```
// 集合结构体
type Set struct {
    m            map[int]struct{} // 用字典来实现，因为字段键不能重复
    len          int          // 集合的大小
    sync.RWMutex              // 锁，实现并发安全
}
```

## 初始化一个集合
```
// 新建一个空集合
func NewSet(cap int64) *Set {
    temp := make(map[int]struct{}, cap)
    return &Set{
        m: temp,
    }
}
```

使用一个容量为`cap`的`map`来实现不可重复集合,`map`的值我们不使用,所以值定义为空结构体`struct{}`因为空结构体不占用空间.如:
```
package main
import (
    "fmt"
    "sync"
)
func main() {
    // 为什么使用空结构体
    a := struct{}{}
    b := struct{}{}
    if a == b {
        fmt.Printf("right:%p\n", &a)
    }
    fmt.Println(unsafe.Sizeof(a))
}
```

```
right:0x1198a98
0
```

空结构体的内存地址都一样,并且不占用内存空间.

## 添加一个元素
```
// 增加一个元素
func (s *Set) Add(item int) {
    s.Lock()
    defer s.Unlock()
    s.m[item] = struct{}{} // 实际往字典添加这个键
    s.len = len(s.m)       // 重新计算元素数量
}
```

首先,加并发锁,实现线程安全,然后往结构体`s *Set`里面的内置`map`添加元素:`item`,元素作为字典的键,会自动去重.同时,集合大小重新生成.

时间复杂度等于字典设置键值对的复杂度,哈希表不冲突的时间复杂度为:`O(1)`,否则为`O(n)`

## 删除一个元素
```
// 查看是否存在元素
func (s *Set) Has(item int) bool {
    s.RLock()
    defer s.RUnlock()
    _, ok := s.m[item]
    return ok
}
```

时间复杂度等于字典获取键值对的复杂度,哈希不冲突的时间复杂度为:`O(1)`,否则为`O(n)`.

## 查看集合大小
```
// 查看集合大小
func (s *Set) Len() int {
    return s.len
}
```

时间复杂度:`O(1)`

## 查看集合是否为空
```
// 集合是够为空
func (s *Set) IsEmpty() bool {
    if s.Len() == 0 {
        return true
    }
    return false
}
```

时间复杂度:`O(1)`

## 清楚集合所有元素
```
// 清除集合所有元素
func (s *Set) Clear() {
    s.Lock()
    defer s.Unlock()
    s.m = map[int]struct{}{} // 字典重新赋值
    s.len = 0                // 大小归零
}
```

将原先的`map`释放掉,并且重新赋一个空的`map`

时间复杂度:`O(1)`

## 将集合转化为列表
```
func (s *Set) List() []int {
    s.RLock()
    defer s.RUnlock()
    list := make([]int, 0, s.len)
    for item := range s.m {
        list = append(list, item)
    }
    return list
}
```

时间复杂度:`O(n)`

## 完整例子