![在这里插入图片描述](https://img-blog.csdnimg.cn/9f413e7efc7a4719838cc178bb635734.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5oiR5piv546L5aSn5a6d,size_18,color_FFFFFF,t_70,g_se,x_16)
**思考：为什么需要GO语言，好在哪里？**

如今的硬件性能越来越强，意味着向同样小的空间里放置更多晶体管的代价也越来越昂贵，为了有更大的提高，现在的处理器添加越来越多的内核，如四核和八核CPU，但是，大多数现代编程语言（如Java，Python等）都是90年代的单线程环境。虽然有许多框架在不断地提高多核资源使用效率，但还是需要花费大量的时间和精力搞懂这些框架的运行原理后才能熟练掌握。Go语言在多核并发上拥有原生的设计优势，Go语言从底层原生支持并发。Go语言的并发是基于 goroutine 的，goroutine 类似于线程，但并非线程。可以将 goroutine 理解为一种虚拟线程。Go 语言运行时会参与调度 goroutine，并将 goroutine 合理地分配到每个 CPU 中，最大限度地使用CPU性能。并且。Go语言简单易学，语法简单，代码风格统一，可以快速集成第三方包等等。

## 并发和并行

GO语言性能强大的原始功臣是 goroutine ，为了更好的理解什么是 goroutine ，我们先了解下并发和并行：

**并发**: 逻辑上具有处理多个同时任务的能力。

**并行**: 物理上同一时刻执行多个并发任务。

我们通常说的并发编程是说程序允许多个任务同时执行，但是实际上不一定是在同一时刻执行的，例如JAVA中，是通过多线程共享CPU时间片串行执行并发的，而并行则依赖于多核处理器（物理资源），让多个任务可以实现并发且并行执行。

操作系统的内核线程 和 编程语言的用户线程 之间，存在3中线程对应模型：

- M:1：多个（M）用户线程始终在一个内核线程上跑，context上下文切换很快，但是无法真正的利用多核。
- 1:1：一个用户线程就只在一个内核线程上跑，这时可以利用多核，但是上下文切换很慢，切换效率很低。
- M:N：多个goroutine在多个内核线程上跑，这个可以集齐上面两者的优势,既能快速切换上下文，也能利用多核的优势，而Go正是选择这种实现方式。
  ![在这里插入图片描述](https://img-blog.csdnimg.cn/59aaa45585f3447697bbf0b44fef5266.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5oiR5piv546L5aSn5a6d,size_20,color_FFFFFF,t_70,g_se,x_16)
  所以简单的将 goroutine 归纳为协程并不合适，，因为它运行时会创建多个线程来执行并发任务，且任务单元可被调度到其它线程执行。

## goroutine 调度器：
Go选择这种线程模式最重要的原因是拥有自己的调度器（GPM）由四部分组成：
![在这里插入图片描述](https://img-blog.csdnimg.cn/66acfb8b7486424da35c50d7858ead13.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5oiR5piv546L5aSn5a6d,size_18,color_FFFFFF,t_70,g_se,x_16)


**G**：

G是goroutine的缩写，goroutine的实体，包括了调用栈，重要的调度信息，例如channel等。

一个 goroutine 的栈在其生命周期开始时只有很小的栈（默认2KB），goroutine的栈不是固定的，他可以按需增大和缩小，goroutine的栈大小限制可以达到1GB，所以在Go语言中一次创建十万左右的goroutine也是可以的。

主要的数据结构：

```go

type g struct {
  stack       stack   // 描述了真实的栈内存，包括上下界

  m              *m     // 当前的m
  sched          gobuf   // goroutine切换时，用于保存g的上下文      
  param          unsafe.Pointer // 用于传递参数，睡眠时其他goroutine可以设置param，唤醒时该goroutine可以获取
  atomicstatus   uint32
  stackLock      uint32 
  goid           int64  // goroutine的ID
  waitsince      int64 // g被阻塞的大体时间
  lockedm        *m     // G被锁定只在这个m上运行
}


//sched了，保存了goroutine的上下文。goroutine切换的时候不同于线程有OS来负责这部分数据，而是由一个gobuf对象来保存，这样能够更加轻量级
type gobuf struct {
    sp   uintptr
    pc   uintptr
    g    guintptr
    ctxt unsafe.Pointer
    ret  sys.Uintreg
    lr   uintptr
    bp   uintptr // for GOEXPERIMENT=framepointer
}
```

**M**：

(Machine)是Go运行时（runtime）对操作系统内核线程的虚拟， M与内核线程是一一映射的关系， 一个groutine最终是要放到M上执行的，所有M是有线程栈的。如果不对该线程栈提供内存的话，系统会给该线程栈提供内存(不同操作系统提供的线程栈大小不同)。

主要的数据结构：

```go
type m struct {
    g0      *g     // 带有调度栈的goroutine

    gsignal       *g         // 处理信号的goroutine
    tls           [6]uintptr // thread-local storage
    mstartfn      func()
    curg          *g       // 当前运行的goroutine
    caughtsig     guintptr 
    p             puintptr // 关联p和执行的go代码
    nextp         puintptr
    id            int32
    mallocing     int32 // 状态

    spinning      bool // m是否out of work
    blocked       bool // m是否被阻塞
    inwb          bool // m是否在执行写屏蔽

    printlock     int8
    incgo         bool // m在执行cgo吗
    fastrand      uint32
    ncgocall      uint64      // cgo调用的总数
    ncgo          int32       // 当前cgo调用的数目
    park          note
    alllink       *m // 用于链接allm
    schedlink     muintptr
    mcache        *mcache // 当前m的内存缓存
    lockedg       *g // 锁定g在当前m上执行，而不会切换到其他m
    createstack   [32]uintptr // thread创建的栈
}
```
**P**：

P(Processor)是一个抽象的概念，并不是真正的物理CPU。所以当P有任务时需要创建或者唤醒一个系统线程来执行它队列里的任务。所以P/M需要进行绑定，构成一个执行单元。P决定了同时可以并发任务的数量，可通过GOMAXPROCS限制同时执行用户级任务的操作系统线程。可以通过runtime.GOMAXPROCS进行指定（最大256）。在Go1.5之后GOMAXPROCS被默认设置可用的核数，而之前则默认为1。

主要数据结构：

```go
type p struct {
    lock mutex

    id          int32
    status      uint32 // 状态，可以为pidle/prunning/...
    link        puintptr
    schedtick   uint32     // 每调度一次加1
    syscalltick uint32     // 每一次系统调用加1
    sysmontick  sysmontick 
    m           muintptr   // 回链到关联的m
    mcache      *mcache
    racectx     uintptr

    goidcache    uint64 // goroutine的ID的缓存
    goidcacheend uint64

    // 可运行的goroutine的队列
    runqhead uint32
    runqtail uint32
    runq     [256]guintptr

    runnext guintptr // 下一个运行的g

    sudogcache []*sudog
    sudogbuf   [128]*sudog

    palloc persistentAlloc // per-P to avoid mutex

    pad [sys.CacheLineSize]byte
}
```
**schedt**：
可以看做是一个全局的调度者。大多数需要的信息都已放在了结构体M、G和P中，schedt结构体只是一个壳。它保存有M的idle队列（idle G的一种状态，刚刚分配还未被初始化），P的idle队列，以及一个全局的就绪的G队列。schedt结构体中的Lock是非常必须的，如果M或P等做一些非局部的操作，它们一般需要先锁住调度器

主要的数据结构：

```go
type schedt struct {
   goidgen  uint64
    lastpoll uint64

    lock mutex

    midle        muintptr // idle状态的m
    nmidle       int32    // idle状态的m个数
    nmidlelocked int32    // lockde状态的m个数
    mcount       int32    // 创建的m的总数
    maxmcount    int32    // m允许的最大个数

    ngsys uint32 // 系统中goroutine的数目，会自动更新

    pidle      puintptr // idle的p
    npidle     uint32
    nmspinning uint32 

    // 全局的可运行的g队列
    runqhead guintptr
    runqtail guintptr
    runqsize int32

    // dead的G的全局缓存
    gflock       mutex
    gfreeStack   *g
    gfreeNoStack *g
    ngfree       int32

    // sudog的缓存中心
    sudoglock  mutex
    sudogcache *sudog
}
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/def0677615fb473f8c997851b2c484fd.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5oiR5piv546L5aSn5a6d,size_20,color_FFFFFF,t_70,g_se,x_16)

## Go调度器调度过程：

首先创建一个G对象，G对象保存到P本地队列或者是全局队列。P此时去唤醒一个M。M寻找是否有空闲的P，获取其资源并执行G。接下来M执行一个调度循环(调用G对象->执行->清理线程→继续找新的Goroutine执行)。

M执行过程中，随时会发生上下文切换。当发生上线文切换时，需要对执行现场进行保护，以便下次被调度执行时进行现场恢复。Go调度器M的栈保存在G对象上，只需要将M所需要的寄存器(SP、PC等)保存到G对象上就可以实现现场保护。当这些寄存器数据被保护起来，就随时可以做上下文切换了，在中断之前把现场保存起来。如果此时G任务还没有执行完，M可以将任务重新丢到P的任务队列，等待下一次被调度执行。当再次被调度执行时，M通过访问G的vdsoSP、vdsoPC寄存器进行现场恢复(从上次中断位置继续执行)。
![在这里插入图片描述](https://img-blog.csdnimg.cn/52e30893bb9c41e1b1513745a417dcdf.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5oiR5piv546L5aSn5a6d,size_20,color_FFFFFF,t_70,g_se,x_16)
我们粗略的看源码可知：

所有的 goroutine 都是有函数 runtime.newproc 来创建的，用系统堆栈，执行runtime.newproc1 来创建并将G绑定到一个M上，启动过程中 调度初始化runtime.schedinit 函数主要根据用户设置的GOMAXPROCS值来创建"P" 这些"p"创建好后都是闲置状态的，都存储在全局调度schedt的pidel字段中等待被使用。如果M的waiting队列中有就从队列中拿。没有的话，需要调用runtime.newm,创建一个。 M 创建好了，线程入口是runtime.mstart1。mstart1里面最重要的就是schedule了，在schedule中的动作大体就是找到一个等待运行的g，然后然后搬到m上，设置其状态为Grunning,直接切换到g的上下文环境,恢复g的执行.

调度的核心就是这个schedule函数了。newm函数会给新的M分配一个空闲的P，执行schedule函数前会拿到相应的P ，这一过程就是acquireq函数。schedule函数很复杂，大致应该分为四大步逻辑：

1. runqget(_g_.m.p.ptr()) , M尝试从自己的P中取出一个G，如果失败，说明P中没有G了。

2. findrunnable, 如果P中没有G了，那M也不能闲着不干活呀，所以M会试图跑去全局队列取一个G来处理；如果全局队里也没有G，M会尝试去别的P中"偷"一个G。如果多次尝试偷G都失败了，那说明实在没有了，这个时候M就会把P还回去，然后回到M池子里。

3. wakep, 到这个过程的时候，M发现自己的P里有好多G，根本处理不过来；再一看居然还有一些闲置的G，回到M池一看，居然还有睡觉的M，直接叫起来干活，分担点M的工作。有时候，M回到池子发现没有在睡觉的其他M，于是会很向系统说——”还有闲置的P啊，我快干不动了，赶紧弄点M吧。”，最后系统搞来一个新的M干活了。

4. execute，M拿起G 愉快的执行了起来
   那么问题来了，是什么时候发生调度呢，难道M等一个G执行完再执行P中的下一个嘛？当然不是，这里就涉及一个上下文切换的问题

**sysmon**：

了解上下文切换，不得不介绍一个sysmon线程。他也是抢占式调度的关键。在初始执行函数runtime.main的时候。第一件事是创建一个新的内核线程M 这个M没有分配P，专门做意见时间，就是系统监控，监控所有P

sysmon会进入一个无限循环, 第一轮回休眠20us, 之后每次休眠时间倍增, 最终休眠10ms. 如果发生过抢占成功，则又恢复中初始的20us的运行时间，sysmon中有netpool(获取fd事件), retake(抢占), forcegc(按时间强制执行gc), scavenge heap(释放自由列表中多余的项减少内存占用)等处理.

**上下文切换的调度点：(G的所有状态说明，在最后有图)**

对channel读写操作的时候调用runtime.park函数，这个会将G设置为Gwaiting的状态。放弃CPU 需要调用runtime.ready才能再次执行；runtime·gosched函数也可以让当前goroutine放弃cpu，但和park完全不同；gosched是将G设置为Grunnable状态，然后放入到调度器全局等待队列。除此之外，就是有些系统调用会将P的状态置为Psyscall。sysmon 线程会扫描所有的P 当发现有p处于Psyscall时，会对其进行抢占。正在执行系统调用的话，将P与M脱离（handoffp），正在执行Go代码的话，通知抢占(preemptone) ,  sysmon线程会创建一个新的M，并把刚才的P抢过来，P中剩下的G在新的M上执行，之前的M等系统调用返回后发现自己的P没有了，只能把自己拿的内个G放回全局队列里，自己回归M池子了
![在这里插入图片描述](https://img-blog.csdnimg.cn/735ca9c0076e4fbf9231e6d9f8e88d5a.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5oiR5piv546L5aSn5a6d,size_20,color_FFFFFF,t_70,g_se,x_16)


