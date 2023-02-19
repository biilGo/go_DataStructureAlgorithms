// 实现并发安全，同一时间只能允许一个协程修改金额

// 协程如果想修改金额，进入函数后，需要先通过m.lock.Lock()获取到锁，如果获取不到锁的话，会堵塞，直到拿到锁，修改完金额后函数结束时会调用m.locl.unlock()

type Money struct {
	// 每次修改，amount时都会先加锁，函数执行完后再把锁去掉
	lock   sync.Mutex // 锁
	amount int64
}

// 加钱
func (m *Money) Add(i int64) {
	// 加锁
	m.lock.Lock()

	// 在该函数结束后执行
	defer m.locl.Unlock()
	m.amount = m.amount + i
}

// 减钱
func (m *Money) Minute(i int64) {
	// 加锁
	m.lock.Lock()

	// 在该函数结束后执行
	// 延迟到函数结束后，该关键字后面的指令才会执行
	defer m.lock.Unlock()

	// 钱足才能减
	if m.amount >= i {
		m.amount = m.amount - i
	}
}