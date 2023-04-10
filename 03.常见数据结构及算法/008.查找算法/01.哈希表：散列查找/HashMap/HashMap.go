package main

import (
	"fmt"
	"math"
	"sync"
	"xxhash"
)

// 使用结构体 HashMap 来表示哈希表：

const (
	// 扩容因子
	expandFactor = 0.75
)

// 哈希表
type HashMap struct {
	array        []*keyPairs // 哈希表数组,每个元素是一个键值对
	capacity     int         // 数组容量
	len          int         // 已添加键值对元素数量
	capacityMask int         // 掩码,等于capacity-1

	// 增删键值对时,需要考虑并发安全
	lock sync.Mutex
}

// 键值对,连城一个链表
type keyPairs struct {
	key   string      // 键
	value interface{} // 值
	next  *keyPairs   // 下一个键值对
}

// 6.1 初始化哈希表

// 创建大小为capacity的哈希表
func NewHashMap(capacity int) *HashMap {
	// 默认大小为16
	defaultCapacity := 1 << 4
	if capacity <= defaultCapacity {
		// 如果传入的大小小于默认大小,那么使用默认大小16
		capacity = defaultCapacity
	} else {
		// 否则,实际大小为大于capacity的抵押给2^k
		capacity = 1 << (int(math.Ceil(math.Log2(float64(capacity)))))
	}

	// 新建一个哈希表
	m := new(HashMap)
	m.array = make([]*keyPairs, capacity, capacity)
	m.capacity = capacity
	m.capacity = capacity - 1
	return m
}

// 返回哈希表已添加元素数量
func (m *HashMap) Len() int {
	return m.len
}

// 我们可以传入capacity来初始化当前哈希表数组容量,容量掩码capacityMask = capacity-1主要用来计算数组下标
// 如果传入的容量大小默认容量16,那么将16作为哈希表的初始数组大小,否则将第一个大于capacity的2^k值作为数组的初始大小.

// 6.2 计算哈希值和数组下标

// 求key的哈希值
var hashAlgorithm = func(key []byte) uint64 {
	h := xxhash.New64()
	h.Write(key)
	return h.Sum64()
}

// 对键进行哈希求值,并计算下标
func (m *HashMap) hashIndex(key string, mask int) int {
	// 求哈希
	hash := hashAlgorithm([]byte(key))

	// 求下标
	index := hash & uint64(mask)
	return int(index)
}

// 首先,为结构体生成一个hashIndex方法
// 根据公式hash(key)&(2^x-1),使用xxhash哈希算法来计算键key的哈希值,并且和容量掩码mask进行&求得数组的下标,用来定位键值对该放在数组的哪个下标下.

// 6.2 添加键值对

// 哈希表添加键值对
func (m *HashMap) Put(key string, value interface{}) {
	// 实现并发安全
	m.lock.Lock()
	defer m.lock.Unlock()

	// 键值对要放的哈希数组下标
	index := m.hashIndex(key, m.capacityMask)

	// 哈希表数组下标的元素
	element := m.array[index]

	// 如果该元素为空表示链表是空的，不存在哈希冲突，直接将键值对作为链表的第一个元素：
	// 元素为空,表示空链表,没有哈希冲突,直接赋值
	if element == nil {
		m.array[index] = &keyPairs{
			key:   key,
			value: value,
		}
	} else {
		// 否则，则遍历链表，查找键是否存在：
		// 链表最后一个键值对
		var lastPairs *keyPairs

		// 遍历链表查看元素是否存在,存在则替换值,否则找到最后一个键值对
		for element != nil {
			// 键值对存在,那么更新值并返回
			// 当 element.key == key ，那么键存在，直接更新值，退出该函数。否则，继续往下遍历。
			if element.key == key {
				element.value = value
				return
			}

			lastPairs = element
			element = element.next
		}

		// 当跳出 for element != nil 时，表示找不到键值对，那么往链表尾部添加该键值对：
		// 找不到键值对,将新键值对添加到链表尾端
		lastPairs.next = &keyPairs{
			key:   key,
			value: value,
		}

		// 最后，检查是否需要扩容，如果需要则扩容：
		// 新的哈希表数量
		newLen := m.len + 1

		// 如果超出扩容因子,需要扩容
		if float64(newLen)/float64(m.capacity) >= expandFactor {
			// 创建了一个新的两倍大小的哈希表：newM := new(HashMap)，然后遍历老哈希表中的键值对，重新 Put 进新哈希表。
			// 新建一个原来两倍大小的哈希表
			newM := new(HashMap)
			newM.array = make([]*keyPairs, 2*m.capacity, 2*m.capacity)
			newM.capacity = 2 * m.capacity
			newM.capacityMask = 2*m.capacity - 1

			// 遍历老的哈希表,将键值对重新哈希到新哈希表
			for _, pairs := range m.array {
				for pairs != nil {
					// 直接递归Put
					newM.Put(pairs.key, pairs.value)
					pairs = pairs.next
				}
			}

			// 替换老的哈希表
			m.array = newM.array
			m.capacity = newM.capacity
			m.capacityMask = newM.capacityMask
		}

		m.len = newLen
	}
}

// 6.3 获取键值对
// 哈希表获取键值对
func (m *HashMap) Get(key string) (value interface{}, ok bool) {
	// 实现并发安全
	m.lock.Lock()
	defer m.lock.Unlock()

	// 同样先加锁实现并发安全，然后进行哈希算法计算出数组下标：index := m.hashIndex(key, m.capacityMask)，取出元素：element := m.array[index]。
	// 键值对要放的哈希表数组下标
	index := m.hashIndex(key, m.capacityMask)

	// 哈希表数组下标的元素
	element := m.array[index]

	// 对链表进行遍历：
	// 遍历链表查看元素是否存在,存在则返回
	for element != nil {
		if element.key == key {
			return element.value, true
		}

		element = element.next
	}

	return
}

// 删除键值对
// 哈希表删除键值对
func (m *HashMap) Delete(key string) {
	// 实现并发安全
	m.lock.Lock()
	defer m.lock.Unlock()

	// 键值对要放的哈希表数组下标
	index := m.hashIndex(key, m.capacityMask)

	// 哈希表数组下标的元素
	element := m.array[index]

	// 如果元素是空的，表示链表为空，那么直接返回：
	// 空链表,不用删除,直接返回
	if element == nil {
		return
	}

	// 否则查看链表第一个元素的键是否匹配：element.key == key，如果匹配，那么对链表头部进行替换，链表的第二个元素补位成为链表头部：
	// 链表的第一个元素就是要删除的元素
	if element.key == key {
		// 将第一个元素后面的键值对链上
		m.array[index] = element.next
		m.len = m.len - 1
		return
	}

	// 如果链表的第一个元素不匹配，那么从第二个元素开始遍历链表，找到时将该键值对删除，然后将前后两个键值对连接起来：
	// 下一个键值对
	nextElement := element.next
	for nextElement != nil {
		if nextElement.key == key {
			// 键值对匹配到,将该键值对从链中去掉
			element.next = nextElement.next
			m.len = m.len - 1
			return
		}

		element = nextElement
		nextElement = nextElement.next
	}
}

// 6.4 遍历哈希表
func (m *HashMap) Range() {
	// 实现并发安全
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, pairs := range m.array {
		for pairs != nil {
			fmt.Printf("'%v' = '%v'", pairs.key, pairs.value)
			pairs = pairs.next
		}
	}
	fmt.Println()
}

func main() {
	// 新建一个哈希表
	hashMap := NewHashMap(16)

	// 放35个值
	for i := 0; i < 35; i++ {
		hashMap.Put(fmt.Sprintf("%d", i), fmt.Sprintf("v%d", i))
	}
	fmt.Println("cap:", hashMap.capacity, "len:", hashMap.Len())

	// 打印全部键值对
	hashMap.Range()

	key := "4"
	value, ok := hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v'='%v'\n", key, value)
	} else {
		fmt.Printf("get %v not found\n", key)
	}

	// 删除键
	hashMap.Delete(key)
	fmt.Println("after delete cap:", hashMap.capacity, "len:", hashMap.Len())
	value, ok = hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v' = '%v'\n", key, value)
	} else {
		fmt.Printf("get %v not found\n", key)
	}
}
