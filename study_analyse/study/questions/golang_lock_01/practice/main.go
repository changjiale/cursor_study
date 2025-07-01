package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
Go语言锁机制实践练习

本文件包含以下练习：
1. 实现线程安全的缓存
2. 死锁检测和避免
3. 性能测试对比
4. 锁优化策略
*/

// 练习1：实现一个线程安全的缓存
type Cache struct {
	mu    sync.RWMutex
	data  map[string]interface{}
	stats struct {
		hits   int64
		misses int64
	}
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	// TODO: 实现读操作，使用读锁
	// 提示：使用 RLock() 和 RUnlock()
	return nil, false
}

func (c *Cache) Set(key string, value interface{}) {
	// TODO: 实现写操作，使用写锁
	// 提示：使用 Lock() 和 Unlock()
}

func (c *Cache) Delete(key string) {
	// TODO: 实现删除操作，使用写锁
	// 提示：使用 Lock() 和 Unlock()
}

func (c *Cache) GetStats() (hits, misses int64) {
	// TODO: 实现统计信息获取，使用读锁
	return 0, 0
}

// 练习2：死锁检测和避免
func deadlockExample() {
	var mu1, mu2 sync.Mutex

	// TODO: 分析这个函数是否存在死锁风险
	// 如果存在，请修改代码避免死锁

	go func() {
		mu1.Lock()
		time.Sleep(100 * time.Millisecond)
		mu2.Lock()
		fmt.Println("Goroutine 1: 获取了两个锁")
		mu2.Unlock()
		mu1.Unlock()
	}()

	go func() {
		mu2.Lock()
		time.Sleep(100 * time.Millisecond)
		mu1.Lock()
		fmt.Println("Goroutine 2: 获取了两个锁")
		mu1.Unlock()
		mu2.Unlock()
	}()

	time.Sleep(1 * time.Second)
}

// 练习3：锁优化策略 - 锁分离
type OptimizedCounter struct {
	counters [256]struct {
		value int64
		mu    sync.Mutex
	}
}

func NewOptimizedCounter() *OptimizedCounter {
	return &OptimizedCounter{}
}

func (oc *OptimizedCounter) Increment(id int) {
	// TODO: 实现基于ID的锁分离
	// 提示：使用 id % 256 来选择不同的锁
}

func (oc *OptimizedCounter) GetValue(id int) int64 {
	// TODO: 实现获取指定ID的计数值
	return 0
}

func (oc *OptimizedCounter) GetTotalValue() int64 {
	// TODO: 实现获取所有计数的总和
	return 0
}

// 练习4：无锁编程 - 原子操作
type LockFreeCounter struct {
	value int64
}

func NewLockFreeCounter() *LockFreeCounter {
	return &LockFreeCounter{}
}

func (lfc *LockFreeCounter) Increment() {
	// TODO: 使用原子操作实现递增
	// 提示：使用 atomic.AddInt64()
}

func (lfc *LockFreeCounter) GetValue() int64 {
	// TODO: 使用原子操作获取值
	// 提示：使用 atomic.LoadInt64()
	return 0
}

// 练习5：性能测试对比
func benchmarkMutex(iterations int) time.Duration {
	var mu sync.Mutex
	var counter int64

	start := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations/10; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return time.Since(start)
}

func benchmarkRWMutex(iterations int) time.Duration {
	var mu sync.RWMutex
	var counter int64

	start := time.Now()

	var wg sync.WaitGroup
	// 启动写goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < iterations/10; j++ {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}()

	// 启动读goroutine
	for i := 0; i < 9; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations/9; j++ {
				mu.RLock()
				_ = counter
				mu.RUnlock()
			}
		}()
	}

	wg.Wait()
	return time.Since(start)
}

func benchmarkAtomic(iterations int) time.Duration {
	var counter int64

	start := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations/10; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	return time.Since(start)
}

// 练习6：实际应用 - 连接池
type Connection struct {
	ID       string
	IsActive bool
	LastUsed time.Time
}

type ConnectionPool struct {
	mu       sync.RWMutex
	conns    map[string]*Connection
	maxConns int
}

func NewConnectionPool(maxConns int) *ConnectionPool {
	return &ConnectionPool{
		conns:    make(map[string]*Connection),
		maxConns: maxConns,
	}
}

func (cp *ConnectionPool) GetConnection(id string) (*Connection, error) {
	// TODO: 实现双重检查锁定模式
	// 1. 首先使用读锁检查连接是否存在
	// 2. 如果不存在，使用写锁创建新连接
	// 3. 在写锁中再次检查（双重检查）

	return nil, fmt.Errorf("connection not found")
}

func (cp *ConnectionPool) ReleaseConnection(id string) {
	// TODO: 实现连接释放
}

func main() {
	fmt.Println("=== Go语言锁机制实践练习 ===\n")

	// 练习1：测试缓存
	fmt.Println("练习1：线程安全缓存")
	cache := NewCache()
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	if value, exists := cache.Get("key1"); exists {
		fmt.Printf("获取到值: %v\n", value)
	} else {
		fmt.Println("未找到key1")
	}

	// 练习2：测试死锁（注释掉，避免程序卡死）
	// fmt.Println("\n练习2：死锁检测")
	// deadlockExample()

	// 练习3：测试优化计数器
	fmt.Println("\n练习3：锁分离优化")
	optCounter := NewOptimizedCounter()
	for i := 0; i < 1000; i++ {
		optCounter.Increment(i)
	}
	fmt.Printf("总计数: %d\n", optCounter.GetTotalValue())

	// 练习4：测试无锁计数器
	fmt.Println("\n练习4：无锁编程")
	lockFreeCounter := NewLockFreeCounter()
	for i := 0; i < 1000; i++ {
		lockFreeCounter.Increment()
	}
	fmt.Printf("无锁计数器值: %d\n", lockFreeCounter.GetValue())

	// 练习5：性能测试
	fmt.Println("\n练习5：性能测试对比")
	iterations := 100000

	mutexTime := benchmarkMutex(iterations)
	rwMutexTime := benchmarkRWMutex(iterations)
	atomicTime := benchmarkAtomic(iterations)

	fmt.Printf("Mutex耗时: %v\n", mutexTime)
	fmt.Printf("RWMutex耗时: %v\n", rwMutexTime)
	fmt.Printf("Atomic耗时: %v\n", atomicTime)

	// 练习6：测试连接池
	fmt.Println("\n练习6：连接池")
	pool := NewConnectionPool(10)
	conn, err := pool.GetConnection("conn1")
	if err != nil {
		fmt.Printf("获取连接失败: %v\n", err)
	} else {
		fmt.Printf("获取连接成功: %s\n", conn.ID)
	}

	fmt.Println("\n=== 练习完成 ===")
	fmt.Println("请根据TODO注释完成各个练习的实现")
}

/*
预期输出示例：
=== Go语言锁机制实践练习 ===

练习1：线程安全缓存
获取到值: value1

练习3：锁分离优化
总计数: 1000

练习4：无锁编程
无锁计数器值: 1000

练习5：性能测试对比
Mutex耗时: 15.2ms
RWMutex耗时: 12.8ms
Atomic耗时: 8.5ms

练习6：连接池
获取连接成功: conn1

=== 练习完成 ===
请根据TODO注释完成各个练习的实现
*/
