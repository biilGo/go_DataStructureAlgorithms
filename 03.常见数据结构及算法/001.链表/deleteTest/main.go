// 删除循环链表后面的三个节点:r.Unlink(3)
// 可以看到节点r后面的两个节点被切断了,然后分成了两个循环链表,r所在的链表变成了-1,1
// 而切除的那部分形成了一个新循环链表是4 3 2,并且返回给了用户
// 因为只要定位要删除的节点位置,然后进行链接:r.Link(r.Move(n+1))所以时间复杂度:O(n)+O(1)=O(n)

package main

import "fmt"

// 循环链表
type Ring struct {
	next, prev *Ring       // 前驱和后驱节点
	Value      interface{} // 数据
}

// 初始化空的循环链表，前驱和后驱都指向自己，因为是循环的
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// 创建N个节点的循环链表
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

// 获取下一个节点
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// 获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

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
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// 往节点A，链接一个节点，并且返回之前节点A的后驱节点
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

func linkNewTest() {
	// 第一个节点
	r := &Ring{Value: -1}

	// 链接新的五个节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})

	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)

		// 移动到下一个节点
		node = node.Next()

		// 如果节点回到了起点,结束
		if node == r {
			return
		}
	}

}

// 删除节点后面的 n 个节点
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

func deleteTest() {
	// 第一个节点
	r := &Ring{Value: -1}

	// 链接到新的五个节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})

	temp := r.Unlink(3) // 解除了后面两个节点

	// 打印原来的节点
	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)
		// 移动到下一节点
		node = node.Next()

		// 如果节点回到了起点,结束
		if node == r {
			break
		}
	}
	fmt.Println("------")

	// 打印被切断的节点
	node = temp
	for {
		// 打印节点值
		fmt.Println(node.Value)
		// 移到下一个节点
		node = node.Next()

		// 如果节点回到了起点,结束
		if node == temp {
			break
		}
	}
}

func main() {
	deleteTest()
}
