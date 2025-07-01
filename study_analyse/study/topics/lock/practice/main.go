package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

/*
Go语言锁机制实践练习

本文件包含以下练习：
1. 基础用法练习
2. 进阶特性练习
3. 性能测试练习
4. 实际应用练习
*/

// 练习1：基础用法练习
func basicPractice() {
	fmt.Println("=== 基础用法练习 ===\n")

	// 1.1 Mutex基础使用
	fmt.Println("1.1 Mutex基础使用")
	var mu sync.Mutex
	var counter int

	// 启动多个goroutine并发增加计数
	for i := 0; i < 5; i++ {
		go func(id int) {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
			fmt.Printf("Goroutine %d 完成\n", id)
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("最终计数: %d (期望: 5000)\n\n", counter)

	// 1.2 RWMutex基础使用
	fmt.Println("1.2 RWMutex基础使用")
	var rw sync.RWMutex
	var data map[string]string = make(map[string]string)

	// 启动多个读goroutine
	for i := 0; i < 3; i++ {
		go func(id int) {
			for j := 0; j < 5; j++ {
				rw.RLock()
				fmt.Printf("读者 %d 读取数据: %v\n", id, data)
				rw.RUnlock()
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	// 启动写goroutine
	go func() {
		for i := 0; i < 3; i++ {
			rw.Lock()
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("value%d", i)
			data[key] = value
			fmt.Printf("写者设置: %s = %s\n", key, value)
			rw.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println()

	// 1.3 原子操作基础使用
	fmt.Println("1.3 原子操作基础使用")
	var atomicCounter int64

	// 启动多个goroutine并发增加计数
	for i := 0; i < 5; i++ {
		go func(id int) {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&atomicCounter, 1)
			}
			fmt.Printf("Goroutine %d 完成\n", id)
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("最终计数: %d (期望: 5000)\n\n", atomicCounter)
}

// 练习2：进阶特性练习
func advancedPractice() {
	fmt.Println("=== 进阶特性练习 ===\n")

	// 2.1 死锁演示
	fmt.Println("2.1 死锁演示（注释掉以避免程序卡死）")
	/*
		var mu1, mu2 sync.Mutex

		go func() {
			mu1.Lock()
			time.Sleep(100 * time.Millisecond)
			mu2.Lock()
			mu2.Unlock()
			mu1.Unlock()
		}()

		go func() {
			mu2.Lock()
			time.Sleep(100 * time.Millisecond)
			mu1.Lock()
			mu1.Unlock()
			mu2.Unlock()
		}()

		time.Sleep(1 * time.Second)
	*/
	fmt.Println("死锁演示已注释，避免程序卡死\n")

	// 2.2 锁分离策略
	fmt.Println("2.2 锁分离策略")
	type ShardedCounter struct {
		counters [4]struct {
			value int64
			mu    sync.Mutex
		}
	}

	sc := &ShardedCounter{}

	// 启动多个goroutine，每个使用不同的分片
	for i := 0; i < 8; i++ {
		go func(id int) {
			shard := id % 4
			for j := 0; j < 1000; j++ {
				sc.counters[shard].mu.Lock()
				sc.counters[shard].value++
				sc.counters[shard].mu.Unlock()
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
	var total int64
	for i := 0; i < 4; i++ {
		total += sc.counters[i].value
	}
	fmt.Printf("分片计数器总和: %d (期望: 8000)\n\n", total)

	// 2.3 内存对齐优化
	fmt.Println("2.3 内存对齐优化")
	type UnalignedStruct struct {
		value int64
		flag  bool
	}

	type AlignedStruct struct {
		value int64
		pad   [56]byte // 填充到64字节
	}

	fmt.Printf("未对齐结构体大小: %d 字节\n", unsafe.Sizeof(UnalignedStruct{}))
	fmt.Printf("对齐结构体大小: %d 字节\n", unsafe.Sizeof(AlignedStruct{}))
	fmt.Println()
}

// 练习3：性能测试练习
func performanceTest() {
	fmt.Println("=== 性能测试练习 ===\n")

	const iterations = 1000000

	// 3.1 原子操作 vs 互斥锁性能对比
	fmt.Println("3.1 原子操作 vs 互斥锁性能对比")

	// 原子操作测试
	start := time.Now()
	var atomicCounter int64
	for i := 0; i < iterations; i++ {
		atomic.AddInt64(&atomicCounter, 1)
	}
	atomicTime := time.Since(start)

	// 互斥锁测试
	start = time.Now()
	var mutexCounter int64
	var mu sync.Mutex
	for i := 0; i < iterations; i++ {
		mu.Lock()
		mutexCounter++
		mu.Unlock()
	}
	mutexTime := time.Since(start)

	fmt.Printf("原子操作耗时: %v\n", atomicTime)
	fmt.Printf("互斥锁耗时: %v\n", mutexTime)
	fmt.Printf("性能差异: %.2f倍\n\n", float64(mutexTime)/float64(atomicTime))

	// 3.2 RWMutex vs Mutex性能对比
	fmt.Println("3.2 RWMutex vs Mutex性能对比（读多写少场景）")

	// RWMutex测试
	start = time.Now()
	var rw sync.RWMutex
	var rwData map[string]string = make(map[string]string)
	rwData["test"] = "value"

	var wg sync.WaitGroup

	// 启动多个读goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations/10; j++ {
				rw.RLock()
				_ = rwData["test"]
				rw.RUnlock()
			}
		}()
	}

	// 启动写goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < iterations/100; j++ {
			rw.Lock()
			rwData["test"] = "newvalue"
			rw.Unlock()
		}
	}()

	wg.Wait()
	rwMutexTime := time.Since(start)

	// Mutex测试
	start = time.Now()
	var mutexData map[string]string = make(map[string]string)
	mutexData["test"] = "value"

	// 启动多个goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations/10; j++ {
				mu.Lock()
				_ = mutexData["test"]
				mu.Unlock()
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < iterations/100; j++ {
			mu.Lock()
			mutexData["test"] = "newvalue"
			mu.Unlock()
		}
	}()

	wg.Wait()
	mutexTime2 := time.Since(start)

	fmt.Printf("RWMutex耗时: %v\n", rwMutexTime)
	fmt.Printf("Mutex耗时: %v\n", mutexTime2)
	fmt.Printf("性能差异: %.2f倍\n\n", float64(mutexTime2)/float64(rwMutexTime))
}

// 练习4：实际应用练习
func practicalApplication() {
	fmt.Println("=== 实际应用练习 ===\n")

	// 4.1 线程安全缓存
	fmt.Println("4.1 线程安全缓存")
	cache := NewThreadSafeCache()

	// 启动多个读goroutine
	for i := 0; i < 5; i++ {
		go func(id int) {
			for j := 0; j < 10; j++ {
				key := fmt.Sprintf("key%d", j)
				if value, exists := cache.Get(key); exists {
					fmt.Printf("读者 %d 读取: %s = %v\n", id, key, value)
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
	fmt.Println()

	// 4.2 高性能计数器
	fmt.Println("4.2 高性能计数器")
	hpc := NewHighPerformanceCounter()

	// 启动多个goroutine
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 10000; j++ {
				hpc.Increment()
			}
			fmt.Printf("Goroutine %d 完成\n", id)
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("高性能计数器总和: %d (期望: 100000)\n\n", hpc.Get())
}

// 线程安全缓存实现
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

// 高性能计数器实现
type HighPerformanceCounter struct {
	counters [256]struct {
		value int64
		pad   [56]byte // 填充到64字节，避免缓存行冲突
	}
}

func NewHighPerformanceCounter() *HighPerformanceCounter {
	return &HighPerformanceCounter{}
}

func (hpc *HighPerformanceCounter) Increment() {
	// 使用简单的哈希函数选择计数器
	// 在实际应用中，可以使用更复杂的哈希函数
	id := time.Now().UnixNano() % 256
	atomic.AddInt64(&hpc.counters[id].value, 1)
}

func (hpc *HighPerformanceCounter) Get() int64 {
	var total int64
	for i := 0; i < 256; i++ {
		total += atomic.LoadInt64(&hpc.counters[i].value)
	}
	return total
}

func main() {
	fmt.Println("=== Go语言锁机制实践练习 ===\n")

	// 执行练习
	basicPractice()
	advancedPractice()
	performanceTest()
	practicalApplication()

	fmt.Println("=== 练习完成 ===")
	fmt.Println("通过以上练习，你可以：")
	fmt.Println("1. 掌握锁的基本用法")
	fmt.Println("2. 理解锁的性能特点")
	fmt.Println("3. 学会性能优化技巧")
	fmt.Println("4. 应用锁解决实际问题")
}
