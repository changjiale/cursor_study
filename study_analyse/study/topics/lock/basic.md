# Go语言锁机制 - 基础概念

## 什么是锁机制

### 定义和概念
锁机制是并发编程中用于保护共享资源的核心机制，确保在任意时刻只有一个或多个特定的goroutine能够访问共享资源，从而避免数据竞争和不一致性问题。

### 基本特性
- **互斥性**：同一时间只能有一个goroutine持有锁
- **可见性**：锁的获取和释放对其他goroutine可见
- **原子性**：锁的获取和释放操作是原子的
- **公平性**：保证等待的goroutine能够公平地获得锁

### 使用方式
```go
// 基本使用模式
var mu sync.Mutex
mu.Lock()
// 临界区代码
defer mu.Unlock()
```

## 核心原理

### 实现机制
Go语言的锁机制基于以下技术实现：

1. **原子操作**：使用CPU的原子指令确保操作的原子性
2. **自旋等待**：在获取锁失败时，短暂自旋等待而不是立即阻塞
3. **信号量机制**：使用操作系统的信号量进行阻塞和唤醒
4. **内存屏障**：确保内存操作的顺序性和可见性

### 工作流程
1. **尝试获取锁**：使用原子操作尝试获取锁
2. **自旋等待**：如果获取失败，短暂自旋等待
3. **进入等待队列**：自旋后仍未获取到锁，进入等待队列
4. **阻塞等待**：使用信号量机制阻塞当前goroutine
5. **被唤醒**：当锁被释放时，唤醒等待的goroutine

### 关键算法

#### Mutex算法
```go
// 简化的Mutex实现逻辑
type Mutex struct {
    state int32  // 锁状态：0=未锁定，1=已锁定
}

func (m *Mutex) Lock() {
    // 1. 尝试直接获取锁
    if atomic.CompareAndSwapInt32(&m.state, 0, 1) {
        return
    }
    
    // 2. 自旋等待
    for i := 0; i < spinCount; i++ {
        if atomic.CompareAndSwapInt32(&m.state, 0, 1) {
            return
        }
        // 短暂自旋
    }
    
    // 3. 进入等待队列
    m.waitForLock()
}
```

#### RWMutex算法
```go
// 简化的RWMutex实现逻辑
type RWMutex struct {
    readerCount int32  // 读者计数
    writerSem   uint32 // 写者信号量
}

func (rw *RWMutex) RLock() {
    // 增加读者计数
    if atomic.AddInt32(&rw.readerCount, 1) < 0 {
        // 有写者在等待，阻塞读者
        rw.waitForWriter()
    }
}

func (rw *RWMutex) Lock() {
    // 将读者计数设为负数，阻止新读者
    r := atomic.AddInt32(&rw.readerCount, -maxReaders)
    // 等待现有读者完成
    if r != 0 {
        rw.waitForReaders()
    }
}
```

## 基础用法

### 基本API

#### Mutex API
```go
type Mutex struct {
    // 私有字段
}

func (m *Mutex) Lock()     // 获取锁
func (m *Mutex) Unlock()   // 释放锁
func (m *Mutex) TryLock() bool // 尝试获取锁（Go 1.18+）
```

#### RWMutex API
```go
type RWMutex struct {
    // 私有字段
}

func (rw *RWMutex) Lock()      // 获取写锁
func (rw *RWMutex) Unlock()    // 释放写锁
func (rw *RWMutex) RLock()     // 获取读锁
func (rw *RWMutex) RUnlock()   // 释放读锁
func (rw *RWMutex) TryLock() bool   // 尝试获取写锁
func (rw *RWMutex) TryRLock() bool  // 尝试获取读锁
```

#### 原子操作API
```go
// 整数原子操作
func AddInt32(addr *int32, delta int32) (new int32)
func AddInt64(addr *int64, delta int64) (new int64)
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
func LoadInt32(addr *int32) (val int32)
func LoadInt64(addr *int64) (val int64)
func StoreInt32(addr *int32, val int32)
func StoreInt64(addr *int64, val int64)

// 指针原子操作
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
```

### 使用示例

#### Mutex基本使用
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Counter struct {
    mu    sync.Mutex
    count int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *Counter) Get() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

func main() {
    counter := &Counter{}
    
    // 启动多个goroutine并发增加计数
    for i := 0; i < 10; i++ {
        go func() {
            for j := 0; j < 1000; j++ {
                counter.Increment()
            }
        }()
    }
    
    time.Sleep(2 * time.Second)
    fmt.Printf("最终计数: %d\n", counter.Get())
}
```

#### RWMutex基本使用
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]string),
    }
}

func (c *Cache) Get(key string) (string, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    value, exists := c.data[key]
    return value, exists
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}

func main() {
    cache := NewCache()
    
    // 启动多个读goroutine
    for i := 0; i < 5; i++ {
        go func(id int) {
            for j := 0; j < 10; j++ {
                key := fmt.Sprintf("key%d", j)
                if value, exists := cache.Get(key); exists {
                    fmt.Printf("读者 %d 读取: %s = %s\n", id, key, value)
                }
                time.Sleep(10 * time.Millisecond)
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
            time.Sleep(50 * time.Millisecond)
        }
    }()
    
    time.Sleep(1 * time.Second)
}
```

#### 原子操作基本使用
```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

type AtomicCounter struct {
    value int64
}

func (ac *AtomicCounter) Increment() {
    atomic.AddInt64(&ac.value, 1)
}

func (ac *AtomicCounter) Get() int64 {
    return atomic.LoadInt64(&ac.value)
}

func (ac *AtomicCounter) CompareAndSwap(old, new int64) bool {
    return atomic.CompareAndSwapInt64(&ac.value, old, new)
}

func main() {
    counter := &AtomicCounter{}
    
    // 启动多个goroutine并发增加计数
    for i := 0; i < 10; i++ {
        go func() {
            for j := 0; j < 1000; j++ {
                counter.Increment()
            }
        }()
    }
    
    time.Sleep(2 * time.Second)
    fmt.Printf("最终计数: %d\n", counter.Get())
    
    // 演示CompareAndSwap
    oldValue := counter.Get()
    if counter.CompareAndSwap(oldValue, oldValue+100) {
        fmt.Printf("成功更新: %d -> %d\n", oldValue, oldValue+100)
    }
}
```

### 注意事项

#### 1. 锁的使用原则
- **最小化临界区**：只在必要的地方加锁，尽快释放锁
- **避免死锁**：固定锁的获取顺序，避免嵌套锁
- **使用defer**：确保锁能够正确释放，即使发生panic
- **避免长时间持有锁**：不要在持有锁时进行耗时操作

#### 2. 常见错误
```go
// 错误：忘记释放锁
func badExample() {
    var mu sync.Mutex
    mu.Lock()
    // 忘记调用 mu.Unlock()
}

// 错误：重复释放锁
func badExample2() {
    var mu sync.Mutex
    mu.Lock()
    mu.Unlock()
    mu.Unlock() // panic: sync: unlock of unlocked mutex
}

// 错误：在持有锁时调用可能阻塞的函数
func badExample3() {
    var mu sync.Mutex
    mu.Lock()
    time.Sleep(1 * time.Second) // 长时间持有锁
    mu.Unlock()
}
```

#### 3. 最佳实践
```go
// 正确：使用defer确保锁释放
func goodExample() {
    var mu sync.Mutex
    mu.Lock()
    defer mu.Unlock()
    // 临界区代码
}

// 正确：最小化临界区
func goodExample2() {
    var mu sync.Mutex
    var data []int
    
    // 在锁外进行耗时操作
    result := expensiveOperation()
    
    // 只在必要时加锁
    mu.Lock()
    data = append(data, result)
    mu.Unlock()
}

// 正确：使用RWMutex优化读多写少场景
func goodExample3() {
    var rw sync.RWMutex
    var data map[string]string = make(map[string]string)
    
    // 读操作使用读锁
    func get(key string) (string, bool) {
        rw.RLock()
        defer rw.RUnlock()
        value, exists := data[key]
        return value, exists
    }
    
    // 写操作使用写锁
    func set(key, value string) {
        rw.Lock()
        defer rw.Unlock()
        data[key] = value
    }
}
```

## 总结

Go语言的锁机制提供了三种主要的并发控制方式：

1. **Mutex**：适用于简单的互斥场景，实现简单，性能适中
2. **RWMutex**：适用于读多写少的场景，能够显著提升并发性能
3. **原子操作**：适用于简单的数值操作，性能最高，但功能有限

选择合适的锁机制需要根据具体的应用场景、性能要求和复杂度来决定。在实际开发中，应该优先考虑原子操作，然后是RWMutex，最后才是Mutex。 