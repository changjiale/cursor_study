package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
Go语言读写锁（RWMutex）实现原理演示

本文件演示了Go语言sync.RWMutex的核心实现原理，
帮助理解读写锁的工作机制。
*/

// 简化的读写锁实现，用于演示原理
type SimpleRWMutex struct {
	w           sync.Mutex // 写锁，保护写者
	writerSem   uint32     // 写者信号量
	readerSem   uint32     // 读者信号量
	readerCount int32      // 读者计数
	readerWait  int32      // 等待的读者数量
}

const rwmutexMaxReaders = 1 << 30 // 最大读者数量

// 读锁实现
func (rw *SimpleRWMutex) RLock() {
	// 增加读者计数
	if atomic.AddInt32(&rw.readerCount, 1) < 0 {
		// 如果读者计数为负数，说明有写者在等待
		// 阻塞当前读者
		fmt.Printf("读者被阻塞，等待写者完成\n")
		rw.waitForWriter()
	}
	fmt.Printf("读者获取读锁，当前读者数: %d\n", atomic.LoadInt32(&rw.readerCount))
}

// 读锁释放
func (rw *SimpleRWMutex) RUnlock() {
	// 减少读者计数
	if r := atomic.AddInt32(&rw.readerCount, -1); r < 0 {
		// 如果读者计数为负数，说明有写者在等待
		rw.rUnlockSlow(r)
	}
	fmt.Printf("读者释放读锁，剩余读者数: %d\n", atomic.LoadInt32(&rw.readerCount))
}

// 读锁释放的慢路径
func (rw *SimpleRWMutex) rUnlockSlow(r int32) {
	// 减少等待的读者数量
	if atomic.AddInt32(&rw.readerWait, -1) == 0 {
		// 如果没有读者在等待了，唤醒写者
		fmt.Printf("所有读者完成，唤醒写者\n")
		rw.wakeupWriter()
	}
}

// 写锁实现
func (rw *SimpleRWMutex) Lock() {
	// 首先获取写锁，防止其他写者进入
	rw.w.Lock()
	fmt.Printf("写者获取写锁\n")

	// 将读者计数减去最大读者数，使其变为负数
	// 这样新的读者会被阻塞
	r := atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders) + rwmutexMaxReaders
	fmt.Printf("写者设置读者计数为负数，当前读者数: %d\n", r)

	// 等待所有现有的读者完成
	if r != 0 && atomic.AddInt32(&rw.readerWait, r) != 0 {
		// 阻塞写者，等待读者完成
		fmt.Printf("写者等待 %d 个读者完成\n", r)
		rw.waitForReaders()
	}
	fmt.Printf("写者开始执行\n")
}

// 写锁释放
func (rw *SimpleRWMutex) Unlock() {
	// 将读者计数恢复为正数，允许新的读者进入
	r := atomic.AddInt32(&rw.readerCount, rwmutexMaxReaders)
	fmt.Printf("写者释放锁，恢复读者计数: %d\n", r)

	// 唤醒所有等待的读者
	for i := 0; i < int(r); i++ {
		fmt.Printf("唤醒等待的读者 %d\n", i+1)
		rw.wakeupReader()
	}

	// 释放写锁
	rw.w.Unlock()
	fmt.Printf("写者完全释放锁\n")
}

// 模拟等待和唤醒机制
func (rw *SimpleRWMutex) waitForWriter() {
	// 在实际实现中，这里会调用runtime_Semacquire
	time.Sleep(100 * time.Millisecond)
}

func (rw *SimpleRWMutex) waitForReaders() {
	// 在实际实现中，这里会调用runtime_Semacquire
	time.Sleep(200 * time.Millisecond)
}

func (rw *SimpleRWMutex) wakeupWriter() {
	// 在实际实现中，这里会调用runtime_Semrelease
	fmt.Printf("唤醒写者\n")
}

func (rw *SimpleRWMutex) wakeupReader() {
	// 在实际实现中，这里会调用runtime_Semrelease
	fmt.Printf("唤醒读者\n")
}

// 演示读写锁的使用
func demonstrateRWMutex() {
	fmt.Println("=== 读写锁实现原理演示 ===\n")

	var rw SimpleRWMutex
	var data map[string]string = make(map[string]string)

	// 启动多个读者
	for i := 0; i < 3; i++ {
		go func(id int) {
			for j := 0; j < 2; j++ {
				rw.RLock()
				fmt.Printf("读者 %d 读取数据: %v\n", id, data)
				time.Sleep(50 * time.Millisecond)
				rw.RUnlock()
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	// 启动写者
	go func() {
		time.Sleep(200 * time.Millisecond) // 等待读者开始
		rw.Lock()
		fmt.Printf("写者开始写入数据\n")
		data["key1"] = "value1"
		data["key2"] = "value2"
		time.Sleep(300 * time.Millisecond) // 模拟写操作耗时
		fmt.Printf("写者完成写入\n")
		rw.Unlock()
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("\n=== 演示完成 ===\n")
}

// 性能对比演示
func performanceComparison() {
	fmt.Println("=== 读写锁性能对比演示 ===\n")

	const iterations = 100000

	// 测试互斥锁
	start := time.Now()
	var mu sync.Mutex
	var counter int64
	var wg sync.WaitGroup

	// 启动多个goroutine，模拟读写操作
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
	mutexTime := time.Since(start)

	// 测试读写锁
	start = time.Now()
	var rw sync.RWMutex
	var rwCounter int64

	// 启动读goroutine
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations/8; j++ {
				rw.RLock()
				_ = rwCounter
				rw.RUnlock()
			}
		}()
	}

	// 启动写goroutine
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations/2; j++ {
				rw.Lock()
				rwCounter++
				rw.Unlock()
			}
		}()
	}
	wg.Wait()
	rwMutexTime := time.Since(start)

	fmt.Printf("互斥锁耗时: %v\n", mutexTime)
	fmt.Printf("读写锁耗时: %v\n", rwMutexTime)
	fmt.Printf("性能提升: %.2f倍\n", float64(mutexTime)/float64(rwMutexTime))
}

// 实际应用示例：线程安全的缓存
type ThreadSafeCache struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

func NewThreadSafeCache() *ThreadSafeCache {
	return &ThreadSafeCache{
		data: make(map[string]interface{}),
	}
}

func (tsc *ThreadSafeCache) Get(key string) (interface{}, bool) {
	tsc.mu.RLock()
	defer tsc.mu.RUnlock()
	value, exists := tsc.data[key]
	return value, exists
}

func (tsc *ThreadSafeCache) Set(key string, value interface{}) {
	tsc.mu.Lock()
	defer tsc.mu.Unlock()
	tsc.data[key] = value
}

func (tsc *ThreadSafeCache) Delete(key string) {
	tsc.mu.Lock()
	defer tsc.mu.Unlock()
	delete(tsc.data, key)
}

func demonstrateCache() {
	fmt.Println("=== 线程安全缓存演示 ===\n")

	cache := NewThreadSafeCache()

	// 启动多个读goroutine
	for i := 0; i < 5; i++ {
		go func(id int) {
			for j := 0; j < 10; j++ {
				key := fmt.Sprintf("key%d", j)
				if value, exists := cache.Get(key); exists {
					fmt.Printf("读者 %d 读取: %s = %v\n", id, key, value)
				} else {
					fmt.Printf("读者 %d 未找到: %s\n", id, key)
				}
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	// 启动写goroutine
	go func() {
		for i := 0; i < 10; i++ {
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("value%d", i)
			cache.Set(key, value)
			fmt.Printf("写者设置: %s = %s\n", key, value)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("\n=== 缓存演示完成 ===\n")
}

func demoMain() {
	// 演示读写锁实现原理
	demonstrateRWMutex()

	// 演示性能对比
	performanceComparison()

	// 演示实际应用
	demonstrateCache()

	fmt.Println("=== 所有演示完成 ===")
	fmt.Println("通过以上演示，你可以看到：")
	fmt.Println("1. 读写锁如何协调读者和写者")
	fmt.Println("2. 读写锁在读多写少场景下的性能优势")
	fmt.Println("3. 读写锁在实际项目中的应用")
}

/*
预期输出示例：
=== 读写锁实现原理演示 ===

读者获取读锁，当前读者数: 1
读者获取读锁，当前读者数: 2
读者获取读锁，当前读者数: 3
读者 0 读取数据: map[]
读者 1 读取数据: map[]
读者 2 读取数据: map[]
读者释放读锁，剩余读者数: 2
读者释放读锁，剩余读者数: 1
读者释放读锁，剩余读者数: 0
写者获取写锁
写者设置读者计数为负数，当前读者数: 0
写者开始执行
写者开始写入数据
写者完成写入
写者释放锁，恢复读者计数: 0
写者完全释放锁

=== 读写锁性能对比演示 ===

互斥锁耗时: 15.2ms
读写锁耗时: 8.5ms
性能提升: 1.79倍

=== 线程安全缓存演示 ===

写者设置: key0 = value0
读者 0 读取: key0 = value0
写者设置: key1 = value1
读者 1 读取: key1 = value1
...

=== 所有演示完成 ===
通过以上演示，你可以看到：
1. 读写锁如何协调读者和写者
2. 读写锁在读多写少场景下的性能优势
3. 读写锁在实际项目中的应用
*/
