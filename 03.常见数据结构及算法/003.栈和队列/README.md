# 栈Stack和队列Queue
我们日常生活中,都需要将物品排列,或者安排事情的先后顺序.更通俗的讲,我们买东西时,人太多的情况下,我们要排队,排队也有先后顺序,有些人早了点来,排完队就离开了,有些人晚一点,才刚刚进入人群排队.

数据是有先后顺序的,从数据`1`到数据`2`,再到数据`3`,和日常生活一样,我们需要放数据,也需要排列数据.

在计算机的世界里,会经常听见两种结构,`栈stack`和`队列queue`.它们是一种收集数据的有序集合,只不过删除和访问数据的顺序不同.

1. 栈:先进后出,先进队的数据最后才出来.在英文的意思里,stack可以作为一叠的意思,这个排列是垂直的,你将一张纸放在另外一张纸上面,先放的纸肯定是最后才会被拿走,因为上面有一张纸挡住了它.
2. 队列:先进先出,先进队的数据先出来.英文的意思里,queue和现实世界的排队意思一样,这个排列是水平的,先排先得.

我们可以用数据结构:`链表`或`数组`来实现`栈`和`队列`

数组实现:能快速随机访问存储的元素,通过下标`index`访问,支持随机访问,查询速度快,但存在元素在数组空间中大量移动的操作,增删效率低.

链表实现:只支持顺序访问,在某些遍历操作中查询的速度慢,但增删元素快.

# 实现数组栈ArrayStack
数组形式的下压栈,后进先出

主要使用可变长数组来实现:
```
// 数组栈,后进先出
type ArrayStack struct {
    array []string // 底层切片
    size  int      // 栈的元素数量
    lock  sync.Mutex // 为了并发安全使用的锁
}
```

## 入栈
```
// 入栈
func (stack *ArrayStack) Push(v string) {
    stack.lock.Lock()
    defer stack.lock.Unlock()
    // 放入切片中，后进的元素放在数组最后面
    stack.array = append(stack.array, v)
    // 栈中元素数量+1
    stack.size = stack.size + 1
}
```

将元素入栈,会先加锁实现并发安全

入栈时直接把元素放在数组的最后面,然后元素数量加1,性能损耗主要花在切片追加元素上,切片如果容量不够会自动扩容,底层损耗的复杂度我们这里不计,所以时间复杂度为`O(1)`

## 出栈
```
func (stack *ArrayStack) Pop() string {
    stack.lock.Lock()
    defer stack.lock.Unlock()
    // 栈中元素已空
    if stack.size == 0 {
        panic("empty")
    }
    // 栈顶元素
    v := stack.array[stack.size-1]
    // 切片收缩，但可能占用空间越来越大
    //stack.array = stack.array[0 : stack.size-1]
    // 创建新的数组，空间占用不会越来越大，但可能移动元素次数过多
    newArray := make([]string, stack.size-1, stack.size-1)
    for i := 0; i < stack.size-1; i++ {
        newArray[i] = stack.array[i]
    }
    stack.array = newArray
    // 栈中元素数量-1
    stack.size = stack.size - 1
    return v
}
```

元素出栈,会先加锁实现并发安全.

如果栈大小为0,那么不允许出栈,否则从数组的最后面拿出元素.

元素取出后:

1. 如果切片偏移量向前移动`stack.array[0:stack.size-1]`表明最后的元素已经不属于该数组了,此时,切片被缩容的部分并不会被回收,仍然占用着空间,所以空间复杂度为:`O(1)`
2. 如果我们创建新的数组`newArray`然后把老数组的元素复制到新数组,就不会占用多余的空间,但移动次数过多,时间复杂度:`O(n)`

最后元素数量减一,并返回值

## 获取栈顶元素
```
// 获取栈顶元素
func (stack *ArrayStack) Peek() string {
    // 栈中元素已空
    if stack.size == 0 {
        panic("empty")
    }
    // 栈顶元素值
    v := stack.array[stack.size-1]
    return v
}
```

获取栈顶元素,但不出栈.和出栈一样,时间复杂度:`O(1)`

## 获取栈大小和判定是否为空
```
// 栈大小
func (stack *ArrayStack) Size() int {
    return stack.size
}
// 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
    return stack.size == 0
}
```

一目了然,时间复杂度:`O(1)`

## 示例:ArrayStack

# 实现链表栈LinkStack
链表形式的下压线,后进先出
```
// 链表栈，后进先出
type LinkStack struct {
    root *LinkNode  // 链表起点
    size int        // 栈的元素数量
    lock sync.Mutex // 为了并发安全使用的锁
}
// 链表节点
type LinkNode struct {
    Next  *LinkNode
    Value string
}
```

## 入栈
```
// 入栈
func (stack *LinkStack) Push(v string) {
    stack.lock.Lock()
    defer stack.lock.Unlock()
    // 如果栈顶为空，那么增加节点
    if stack.root == nil {
        stack.root = new(LinkNode)
        stack.root.Value = v
    } else {
        // 否则新元素插入链表的头部
        // 原来的链表
        preNode := stack.root
        // 新节点
        newNode := new(LinkNode)
        newNode.Value = v
        // 原来的链表链接到新元素后面
        newNode.Next = preNode
        // 将新节点放在头部
        stack.root = newNode
    }
    // 栈中元素数量+1
    stack.size = stack.size + 1
}
```

将元素入栈,会先加锁实现并发安全.

如果栈里面的底层链表为空,表明没有元素,那么新建节点并设置链表起点:`stack.root=new(LinkNode)`

否则取出老的节点:`preNode := stack.root`,新建节点:`newNode := new(LinkNode)`然后将原来的老节点链接在新节点后面:`newNode.Next = preNode`最后将新节点设置为链表起点:`stack.root = newNode`

时间复杂度:`O(1)`

## 出栈
```
// 出栈
func (stack *LinkStack) Pop() string {
    stack.lock.Lock()
    defer stack.lock.Unlock()
    // 栈中元素已空
    if stack.size == 0 {
        panic("empty")
    }
    // 顶部元素要出栈
    topNode := stack.root
    v := topNode.Value
    // 将顶部元素的后继链接链上
    stack.root = topNode.Next
    // 栈中元素数量-1
    stack.size = stack.size - 1
    return v
}
```

元素出栈,如果栈大小为0,那么不允许出栈.

直接将链表的第一个节点`topNode := stack.root`的值取出,然后将表头设置为链表的下一个节点:`stack.root = topNode.Next`,相当于移除了链表的第一个节点.

时间复杂度`O(1)`

## 获取栈顶元素
```
// 获取栈顶元素
func (stack *LinkStack) Peek() string {
    // 栈中元素已空
    if stack.size == 0 {
        panic("empty")
    }
    // 顶部元素值
    v := stack.root.Value
    return v
}
```

获取栈顶元素,但不出栈.和出栈一样,时间复杂度`O(1)`

## 获取栈大小和判定是否为空
```
// 栈大小
func (stack *LinkStack) Size() int {
    return stack.size
}
// 栈是否为空
func (stack *LinkStack) IsEmpty() bool {
    return stack.size == 0
}
```

## 示例:LinkStack

# 实现数组队列ArrayQueue
队列先进先出,和栈操作顺序相反,我们这里只实现入队,和出队操作.其他操作和栈一样.
```
// 数组队列，先进先出
type ArrayQueue struct {
    array []string   // 底层切片
    size  int        // 队列的元素数量
    lock  sync.Mutex // 为了并发安全使用的锁
}
```

## 入队
```
// 入队
func (queue *ArrayQueue) Add(v string) {
    queue.lock.Lock()
    defer queue.lock.Unlock()
    // 放入切片中，后进的元素放在数组最后面
    queue.array = append(queue.array, v)
    // 队中元素数量+1
    queue.size = queue.size + 1
}
```

直接将元素放在数组最后面即可,和栈一样,时间复杂度`O(n)`

## 出队
```
// 出队
func (queue *ArrayQueue) Remove() string {
    queue.lock.Lock()
    defer queue.lock.Unlock()
    // 队中元素已空
    if queue.size == 0 {
        panic("empty")
    }
    // 队列最前面元素
    v := queue.array[0]
    /*    直接原位移动，但缩容后继的空间不会被释放
        for i := 1; i < queue.size; i++ {
            // 从第一位开始进行数据移动
            queue.array[i-1] = queue.array[i]
        }
        // 原数组缩容
        queue.array = queue.array[0 : queue.size-1]
    */
    // 创建新的数组，移动次数过多
    newArray := make([]string, queue.size-1, queue.size-1)
    for i := 1; i < queue.size; i++ {
        // 从老数组的第一位开始进行数据移动
        newArray[i-1] = queue.array[i]
    }
    queue.array = newArray
    // 队中元素数量-1
    queue.size = queue.size - 1
    return v
}
```

出队,把数组的第一个元素的值返回,并对数据进行空间挪位,挪位有2种:

1. 原地挪位,依次补位`queue.array[i-1] = queue.array[i]`,然后数组缩容:`queue.array = queue.array[0:queue.size-1]`,但是这样切片缩容的那部分内存空间不会释放
2. 创建新的数组,将老数组中除第一个元素以外的元素移动到新数组

时间复杂度`O(n)`

# 实现链表队列LinkQueue
队列先进先出,和栈操作顺序相反,我们这里只实现入队,和出队操作,其他操作和栈一样
```
// 链表队列，先进先出
type LinkQueue struct {
    root *LinkNode  // 链表起点
    size int        // 队列的元素数量
    lock sync.Mutex // 为了并发安全使用的锁
}
// 链表节点
type LinkNode struct {
    Next  *LinkNode
    Value string
}
```

## 入队
```
// 入队
func (queue *LinkQueue) Add(v string) {
    queue.lock.Lock()
    defer queue.lock.Unlock()
    // 如果栈顶为空，那么增加节点
    if queue.root == nil {
        queue.root = new(LinkNode)
        queue.root.Value = v
    } else {
        // 否则新元素插入链表的末尾
        // 新节点
        newNode := new(LinkNode)
        newNode.Value = v
        // 一直遍历到链表尾部
        nowNode := queue.root
        for nowNode.Next != nil {
            nowNode = nowNode.Next
        }
        // 新节点放在链表尾部
        nowNode.Next = newNode
    }
    // 队中元素数量+1
    queue.size = queue.size + 1
}
```

将元素放在链表的末尾,所以需要遍历链表,时间复杂度:`O(n)`

## 出队
```
// 出队
func (queue *LinkQueue) Remove() string {
    queue.lock.Lock()
    defer queue.lock.Unlock()
    // 队中元素已空
    if queue.size == 0 {
        panic("empty")
    }
    // 顶部元素要出队
    topNode := queue.root
    v := topNode.Value
    // 将顶部元素的后继链接链上
    queue.root = topNode.Next
    // 队中元素数量-1
    queue.size = queue.size - 1
    return v
}
```