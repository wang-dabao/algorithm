### sync.map 产生背景
Go 的内建 map 是不支持并发写操作的，原因是 map 写操作不是并发安全的，当你尝试多个 Goroutine 操作同一个 map，会产生报错：fatal error: concurrent map writes。
因此官方另外引入了 sync.Map 来满足并发编程中的应用。

### 实现原理
sync.Map 的实现原理可概括为：
1. 通过 read 和 dirty 两个字段将读写分离，读的数据存在只读字段 read 上，将最新写入的数据则存在 dirty 字段上
2. 读取时会先查询 read，不存在再查询 dirty，写入时则只写入 dirty
3. 读取 read 并不需要加锁，而读或写 dirty 都需要加锁
4. 另外有 misses 字段来统计 read 被穿透的次数（被穿透指需要读 dirty 的情况），超过一定次数（超过len(dirty)长度）则将 dirty 数据同步到 read 上
5. 对于删除数据则直接通过标记来延迟删除

### 数据结构
1. map 数据结构
```go
type Map struct {
    // 加锁作用，保护 dirty 字段
    mu Mutex
    // 只读的数据，实际数据类型为 readOnly
    read atomic.Value
    // 最新写入的数据
    dirty map[interface{}]*entry
    // 计数器，每次需要读 dirty 则 +1, 统计被穿透的次数
    misses int
}
```
Map 常用的有以下方法：
Load：读取指定 key 返回 value; 先从read中读取数据，读不到，再通过互斥锁从dirty读数据。

Store： 存储（增或改）key-value
如果read中包含该key, 说明是更新操作，更新read中的该值即可。
如果read中不存在，需要加锁访问dirty,如果dirty中存在该key,执行加锁更新；否则就是一个新key, 加锁添加到dirty中，并在read中标记 amended为true, 说明dirty中有该元素，而read中不包含；

Delete： 删除指定 key。如果 key 在 read 中，就将值置成 nil；如果在 dirty 中，直接删除 key。


2. readonly 数据结构
```go
type readOnly struct {
    // 内建 map
    m  map[interface{}]*entry
    // 表示 dirty 里存在 read 里没有的 key，通过该字段决定是否加锁读 dirty
    amended bool
}

// entry 数据结构用于存储值的指针
type entry struct {
    p unsafe.Pointer  // 等同于 *interface{}
}

```
其中：
entry 数据结构用于存储值的指针，属性p有三种状态：
* p == nil: 键值已经被删除，且 m.dirty == nil
* p == expunged: 键值已经被删除，但 m.dirty != nil 且 m.dirty 不存在该键值
* 除以上情况，则键值对存在，存在于 m.read.m 中，如果 m.dirty!=nil 则也存在于 m.dirty

### 总结
可见，通过这种读写分离的设计，解决了并发情况的写入安全，又使读取速度在大部分情况可以接近内建 map，非常适合`读多写少`的情况。

sync.Map 还有一些其他方法：
Range：遍历所有键值对，参数是回调函数
LoadOrStore：读取数据，若不存在则保存再读取
