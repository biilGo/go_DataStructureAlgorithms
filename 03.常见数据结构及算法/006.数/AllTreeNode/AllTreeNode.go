package main

import (
	"fmt"
	"sync"
)

// 二叉树
type TreeNode struct {
	Data  string    // 节点用来存放数据
	Left  *TreeNode // 左子树
	Right *TreeNode // 右子树
}

func LayerOrder(TreeNode *TreeNode) {
	if TreeNode == nil {
		return
	}

	// 新建队列
	queue := new(LinkQueue)

	// 根节点先入队
	queue.Add(TreeNode)
	for queue.size > 0 {
		// 不断出队列
		element := queue.Remove()

		// 先打印节点值
		fmt.Println(element.Data, " ")

		// 左子树非空,入队列
		if element.Left != nil {
			queue.Add(element.Left)
		}

		// 右子树非空,入队列
		if element.Right != nil {
			queue.Add(element.Right)
		}
	}
}

// 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value *TreeNode
}

// 链表队列,先进先出
type LinkQueue struct {
	root *LinkNode  // 链表起点
	size int        // 队列的元素数量
	lock sync.Mutex // 为了并发安全使用的锁
}

// 入队
func (queue *LinkQueue) Add(v *TreeNode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 如果栈顶为空,那么增加节点
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

	queue.size = queue.size + 1
}

// 出队
func (queue *LinkQueue) Remove() *TreeNode {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 队中元素已空
	if queue.size == 0 {
		panic("over limit")
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

// 队列中元素数量
func (queue *LinkQueue) Size() int {
	return queue.size
}

func main() {
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}
	t.Left.Left = &TreeNode{Data: "D"}
	t.Left.Right = &TreeNode{Data: "E"}
	t.Right.Left = &TreeNode{Data: "F"}

	fmt.Println("\n层次排序")
	LayerOrder(t)
}
