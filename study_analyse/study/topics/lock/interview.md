# Go语言锁机制 - 面试题目

## 基础题目

### 题目1：Go的Mutex是可重入的吗？为什么？
**考察点**：Mutex的设计理念和实现原理
**难度**：中等
**答案**：
Go的Mutex是不可重入的。这与Java的synchronized不同，主要原因有：

1. **设计理念不同**：Go追求简洁性，避免复杂性
2. **避免死锁**：可重入锁可能导致更复杂的死锁场景
3. **性能考虑**：不可重入锁实现更简单，性能更好

```go
// 错误示例：会导致死锁
func recursiveFunction(mu *sync.Mutex) {
    mu.Lock()
    defer mu.Unlock()
    
    // 这里会死锁，因为同一个goroutine试图再次获取已持有的锁
    recursiveFunction(mu)
}
```

### 题目2：读写锁的适用场景是什么？
**考察点**：RWMutex的使用场景和性能特点
**难度**：简单
**答案**：
读写锁适用于读多写少的场景，具体包括：

1. **配置缓存**：配置信息经常被读取，偶尔更新
2. **数据库连接池**：连接信息频繁查询，偶尔变更
3. **缓存系统**：缓存数据大量读取，定期更新
4. **共享数据结构**：如map、slice等需要并发访问的数据

```go
// 典型应用：配置缓存
type ConfigCache struct {
    mu   sync.RWMutex
    data map[string]interface{}
}

func (cc *ConfigCache) Get(key string) (interface{}, bool) {
    cc.mu.RLock()
    defer cc.mu.RUnlock()
    return cc.data[key], true
}

func (cc *ConfigCache) Set(key string, value interface{}) {
    cc.mu.Lock()
    defer cc.mu.Unlock()
    cc.data[key] = value
}
```

### 题目3：如何避免死锁？
**考察点**：死锁预防和检测
**难度**：中等
**答案**：
避免死锁的方法：

1. **固定锁顺序**：总是按照相同的顺序获取锁
2. **避免嵌套锁**：尽量不要在持有锁时获取其他锁
3. **使用超时机制**：设置锁获取的超时时间
4. **使用单一锁**：用一个大锁替代多个小锁

```go
// 正确：固定锁顺序
func transfer(from, to *Account, amount int) {
    // 总是先获取ID小的账户的锁
    if from.id < to.id {
        from.mu.Lock()
        to.mu.Lock()
    } else {
        to.mu.Lock()
        from.mu.Lock()
    }
    defer from.mu.Unlock()
    defer to.mu.Unlock()
    
    // 转账逻辑
}

// 错误：可能导致死锁
func transferBad(from, to *Account, amount int) {
    from.mu.Lock()
    to.mu.Lock() // 如果另一个goroutine同时执行，可能死锁
    // ...
}
```

### 题目4：原子操作和锁的性能差异有多大？
**考察点**：性能对比和选择策略
**难度**：中等
**答案**：
原子操作比锁快5-10倍，具体差异取决于：

1. **操作复杂度**：简单操作差异更大
2. **竞争程度**：高竞争下差异更明显
3. **硬件架构**：不同CPU的原子指令性能不同

```go
// 性能对比示例
func benchmarkComparison() {
    const iterations = 1000000
    
    // 原子操作
    start := time.Now()
    var atomicCounter int64
    for i := 0; i < iterations; i++ {
        atomic.AddInt64(&atomicCounter, 1)
    }
    atomicTime := time.Since(start)
    
    // 互斥锁
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
    fmt.Printf("性能差异: %.2f倍\n", float64(mutexTime)/float64(atomicTime))
}
```

### 题目5：如何实现一个高性能的计数器？
**考察点**：性能优化和锁分离策略
**难度**：困难
**答案**：
可以使用锁分离策略实现高性能计数器：

```go
// 高性能计数器实现
type HighPerformanceCounter struct {
    counters [256]struct {
        value int64
        pad   [56]byte // 填充到64字节，避免缓存行冲突
    }
}

func (hpc *HighPerformanceCounter) Increment() {
    // 使用goroutine ID的哈希值选择计数器
    id := getGoroutineID()
    index := id % 256
    atomic.AddInt64(&hpc.counters[index].value, 1)
}

func (hpc *HighPerformanceCounter) Get() int64 {
    var total int64
    for i := 0; i < 256; i++ {
        total += atomic.LoadInt64(&hpc.counters[i].value)
    }
    return total
}

// 获取goroutine ID（简化实现）
func getGoroutineID() int {
    // 实际实现需要使用runtime包
    return 0
}
```

## 进阶题目

### 题目1：RWMutex的实现原理是什么？
**考察点**：读写锁的底层实现
**难度**：困难
**答案**：
RWMutex的核心实现基于读者计数机制：

```go
// 简化的RWMutex实现
type RWMutex struct {
    w           Mutex  // 写锁
    readerCount int32  // 读者计数
    readerWait  int32  // 等待的读者数量
    writerSem   uint32 // 写者信号量
    readerSem   uint32 // 读者信号量
}

const rwmutexMaxReaders = 1 << 30

func (rw *RWMutex) RLock() {
    // 增加读者计数
    if atomic.AddInt32(&rw.readerCount, 1) < 0 {
        // 有写者在等待，阻塞读者
        runtime_Semacquire(&rw.readerSem)
    }
}

func (rw *RWMutex) Lock() {
    // 获取写锁
    rw.w.Lock()
    
    // 将读者计数设为负数，阻止新读者
    r := atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders) + rwmutexMaxReaders
    
    // 等待现有读者完成
    if r != 0 && atomic.AddInt32(&rw.readerWait, r) != 0 {
        runtime_Semacquire(&rw.writerSem)
    }
}
```

### 题目2：如何实现一个无锁的队列？
**考察点**：无锁编程和CAS操作
**难度**：困难
**答案**：
可以使用CAS操作实现无锁队列：

```go
// 无锁队列实现
type LockFreeQueue struct {
    head *Node
    tail *Node
}

type Node struct {
    value interface{}
    next  *Node
}

func (lfq *LockFreeQueue) Enqueue(value interface{}) {
    newNode := &Node{value: value}
    
    for {
        tail := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&lfq.tail)))
        tailNode := (*Node)(tail)
        
        next := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&tailNode.next)))
        
        if tail == atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&lfq.tail))) {
            if next == nil {
                if atomic.CompareAndSwapPointer(
                    (*unsafe.Pointer)(unsafe.Pointer(&tailNode.next)),
                    next,
                    unsafe.Pointer(newNode)) {
                    atomic.CompareAndSwapPointer(
                        (*unsafe.Pointer)(unsafe.Pointer(&lfq.tail)),
                        tail,
                        unsafe.Pointer(newNode))
                    return
                }
            } else {
                atomic.CompareAndSwapPointer(
                    (*unsafe.Pointer)(unsafe.Pointer(&lfq.tail)),
                    tail,
                    next)
            }
        }
    }
}
```

### 题目3：Mutex的自旋等待机制是如何工作的？
**考察点**：Mutex的优化策略
**难度**：困难
**答案**：
Mutex使用自旋等待来减少上下文切换：

```go
// 简化的自旋等待实现
func (m *Mutex) Lock() {
    // 快速路径：直接尝试获取锁
    if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
        return
    }
    
    // 慢路径：自旋等待
    m.lockSlow()
}

func (m *Mutex) lockSlow() {
    var waitStartTime int64
    starving := false
    awoke := false
    iter := 0
    
    for {
        old := m.state
        new := old | mutexLocked
        
        if old&mutexLocked != 0 {
            // 锁已被占用，尝试自旋
            if runtime_canSpin(iter) {
                // 自旋等待
                if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 &&
                    atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
                    awoke = true
                }
                runtime_doSpin()
                iter++
                continue
            }
            // 自旋次数用完，进入等待队列
            new = old + 1<<mutexWaiterShift
        }
        
        // 尝试获取锁
        if atomic.CompareAndSwapInt32(&m.state, old, new) {
            if old&mutexLocked == 0 {
                break
            }
            // 进入等待队列
            runtime_Semacquire(&m.sema)
            awoke = true
            iter = 0
        }
    }
}
```

## 实战题目

### 题目1：设计一个线程安全的缓存系统
**场景**：高并发Web应用中的配置缓存
**要求**：
1. 支持并发读取和写入
2. 支持过期时间
3. 支持最大容量限制
4. 高性能

**解答**：
```go
// 线程安全缓存实现
type Cache struct {
    mu       sync.RWMutex
    data     map[string]*cacheItem
    maxSize  int
    size     int
}

type cacheItem struct {
    value      interface{}
    expireTime time.Time
}

func NewCache(maxSize int) *Cache {
    cache := &Cache{
        data:    make(map[string]*cacheItem),
        maxSize: maxSize,
    }
    
    // 启动清理过期项的goroutine
    go cache.cleanup()
    return cache
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    item, exists := c.data[key]
    if !exists {
        return nil, false
    }
    
    if time.Now().After(item.expireTime) {
        return nil, false
    }
    
    return item.value, true
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    // 检查容量限制
    if c.size >= c.maxSize {
        c.evictOldest()
    }
    
    expireTime := time.Now().Add(ttl)
    c.data[key] = &cacheItem{
        value:      value,
        expireTime: expireTime,
    }
    c.size++
}

func (c *Cache) Delete(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    if _, exists := c.data[key]; exists {
        delete(c.data, key)
        c.size--
    }
}

func (c *Cache) evictOldest() {
    var oldestKey string
    var oldestTime time.Time
    
    for key, item := range c.data {
        if oldestKey == "" || item.expireTime.Before(oldestTime) {
            oldestKey = key
            oldestTime = item.expireTime
        }
    }
    
    if oldestKey != "" {
        delete(c.data, oldestKey)
        c.size--
    }
}

func (c *Cache) cleanup() {
    ticker := time.NewTicker(time.Minute)
    defer ticker.Stop()
    
    for range ticker.C {
        c.mu.Lock()
        now := time.Now()
        for key, item := range c.data {
            if now.After(item.expireTime) {
                delete(c.data, key)
                c.size--
            }
        }
        c.mu.Unlock()
    }
}
```

### 题目2：实现一个高性能的连接池
**场景**：数据库连接池管理
**要求**：
1. 支持连接复用
2. 支持连接健康检查
3. 支持动态扩容
4. 线程安全

**解答**：
```go
// 高性能连接池实现
type ConnectionPool struct {
    mu          sync.RWMutex
    connections chan *Connection
    factory     func() (*Connection, error)
    maxIdle     int
    maxActive   int
    current     int
    closed      bool
}

type Connection struct {
    conn    interface{}
    pool    *ConnectionPool
    created time.Time
    lastUsed time.Time
}

func NewConnectionPool(factory func() (*Connection, error), maxIdle, maxActive int) *ConnectionPool {
    pool := &ConnectionPool{
        connections: make(chan *Connection, maxIdle),
        factory:     factory,
        maxIdle:     maxIdle,
        maxActive:   maxActive,
    }
    
    // 启动健康检查
    go pool.healthCheck()
    return pool
}

func (cp *ConnectionPool) Get() (*Connection, error) {
    if cp.closed {
        return nil, errors.New("pool is closed")
    }
    
    // 尝试从池中获取连接
    select {
    case conn := <-cp.connections:
        conn.lastUsed = time.Now()
        return conn, nil
    default:
        // 池中没有可用连接，创建新连接
        cp.mu.Lock()
        if cp.current >= cp.maxActive {
            cp.mu.Unlock()
            // 等待连接释放
            select {
            case conn := <-cp.connections:
                conn.lastUsed = time.Now()
                return conn, nil
            case <-time.After(time.Second * 5):
                return nil, errors.New("timeout waiting for connection")
            }
        }
        
        cp.current++
        cp.mu.Unlock()
        
        conn, err := cp.factory()
        if err != nil {
            cp.mu.Lock()
            cp.current--
            cp.mu.Unlock()
            return nil, err
        }
        
        conn.pool = cp
        conn.created = time.Now()
        conn.lastUsed = time.Now()
        return conn, nil
    }
}

func (cp *ConnectionPool) Put(conn *Connection) {
    if cp.closed {
        return
    }
    
    // 检查连接是否健康
    if !cp.isHealthy(conn) {
        cp.mu.Lock()
        cp.current--
        cp.mu.Unlock()
        return
    }
    
    // 尝试放回池中
    select {
    case cp.connections <- conn:
        // 成功放回
    default:
        // 池已满，关闭连接
        cp.mu.Lock()
        cp.current--
        cp.mu.Unlock()
    }
}

func (cp *ConnectionPool) isHealthy(conn *Connection) bool {
    // 检查连接是否过期
    if time.Since(conn.created) > time.Hour {
        return false
    }
    
    // 这里可以添加更多的健康检查逻辑
    return true
}

func (cp *ConnectionPool) healthCheck() {
    ticker := time.NewTicker(time.Minute)
    defer ticker.Stop()
    
    for range ticker.C {
        if cp.closed {
            return
        }
        
        cp.mu.Lock()
        // 清理过期连接
        for {
            select {
            case conn := <-cp.connections:
                if cp.isHealthy(conn) {
                    cp.connections <- conn
                    break
                } else {
                    cp.current--
                }
            default:
                goto done
            }
        }
    done:
        cp.mu.Unlock()
    }
}

func (cp *ConnectionPool) Close() {
    cp.mu.Lock()
    defer cp.mu.Unlock()
    
    cp.closed = true
    close(cp.connections)
    
    // 关闭所有连接
    for {
        select {
        case conn := <-cp.connections:
            // 关闭连接
        default:
            return
        }
    }
}
```

## 面试技巧

### 1. 回答要点
- **理论结合实践**：不仅要知道概念，还要能举出实际例子
- **性能数据**：记住关键的性能对比数据
- **优缺点分析**：能够分析不同方案的优缺点
- **实际应用**：能够描述在实际项目中的应用场景

### 2. 常见陷阱
- **死锁问题**：面试官可能会问具体的死锁场景
- **性能优化**：可能会要求优化现有代码
- **源码分析**：可能会要求分析具体的实现细节
- **设计模式**：可能会要求设计新的并发控制方案

### 3. 准备建议
- **多练习**：动手实现各种锁机制
- **性能测试**：对比不同方案的性能
- **源码阅读**：理解Go语言锁的底层实现
- **项目经验**：总结实际项目中的使用经验 