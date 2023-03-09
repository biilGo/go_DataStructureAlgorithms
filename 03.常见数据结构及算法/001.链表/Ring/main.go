package main

// 循环链表
type Ring struct {
	next, prev *Ring       // 前驱和后驱节点
	Value      interface{} // 数据
}

// 1.1 初始化空的循环链表,前驱和后驱都指向自己,因为是循环的
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

func main() {
	r := new(Ring)
	r.init()
}

// 因为绑定前驱和后驱节点为自己,没有循环,时间复杂度为:`O(1)`

// 创建一个指定大小N的循环链表,值全为空
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

// 会连续绑定前驱和后驱节点,时间复杂度为:O(n)

// 1.2获取上一个节点或下一个节点
// 获取下一个节点
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next

}

// 获取上一节点
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

// 获取前驱和后驱节点,时间复杂度为:O(1)

// 1.2 获取第n个节点
// 因为链表是循环的,当n为负数,表示从前面往前遍历,否则往后面遍历
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n++ {
			r = r.next
		}
	}
	return r
}

// 因为需要遍历n次,时间复杂度为:O(n)

// 1.3 添加节点
// 往节点A,链接一个节点,并且返回之前节点A的后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}
