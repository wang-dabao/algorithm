### context的引入原因
在 Go http包的Server中，每一个请求在都有一个对应的 goroutine 去处理。
请求处理函数通常会启动额外的 goroutine 用来访问后端服务，比如数据库和RPC服务。用来处理一个请求的 goroutine 通常需要访问一些与请求特定的数据，
比如终端用户的身份认证信息、验证相关的token、请求的截止时间。 当一个请求被取消或超时时，所有用来处理该请求的 goroutine 都应该迅速退出，
然后系统才能释放这些 goroutine 占用的资源。

#### context 作用
* 用来设置截止日期: WithDeadline
* 设置同步信号: WithCancel
* 传递请求参数: WithValue

#### context.Context
是 Go 语言在 1.7 版本中引入标准库的接口1，该接口定义了四个需要实现的方法，其中包括：
```
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
```
* Deadline — 返回 context.Context 被取消的时间，也就是完成工作的截止日期；
* Done — 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消后关闭，多次调用 Done 方法会返回同一个 Channel；
* Err — 返回 context.Context 结束的原因，它只会在 Done 方法对应的 Channel 关闭时返回非空的值；
    * 如果 context.Context 被取消，会返回 Canceled 错误；
    * 如果 context.Context 超时，会返回 DeadlineExceeded 错误；
* Value — 从 context.Context 中获取键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法可以用来传递请求特定的数据；


#### 设计原理
在 Goroutine 构成的树形结构中对信号进行同步以`减少计算资源的浪费`是 `context.Context 的最大作用`。
每一个 context.Context 都会从最顶层的 Goroutine 一层一层传递到最下层。context.Context 可以在上层 Goroutine 执行出现错误时，将信号及时同步给下层。

* Context 与 Goroutine 树
* 不使用 Context 同步信号
* 使用 Context 同步信号
    context.Withtimeout(ctx, 1*time.Second)

#### 默认上下文 (context.Background() && context.Todo) 
这2个方法会返回预先初始化好的私有变量 background 和 todo. 
这两个私有变量都是通过 new(emptyCtx) 语句初始化的，它们是指向私有结构体 context.emptyCtx 的指针。是一个不可取消，没有设置截止时间，没有携带任何值的Context。

二者互为别名，没有太大差别，只是在使用和语义上稍有不同：
* context.Background: 主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。
* context.TODO ：应该在不确定使用哪种上下文时使用

#### 取消信号
context.WithCancel 函数能够从 context.Context 中衍生出一个新的子上下文并返回用于取消该上下文的函数。
一旦我们执行返回的取消函数，当前上下文以及它的子上下文都会被取消，所有的 Goroutine 都会同步收到这一取消信号。

context.WithDeadline 和 context.WithTimeout 也都能创建可以被取消的计时器上下文 context.timerCtx：
context.WithDeadline 在创建 context.timerCtx 的过程中判断了父上下文的截止日期与当前日期，并通过 time.AfterFunc 创建定时器，当时间超过了截止日期后会调用 context.timerCtx.cancel 同步取消信号。

#### 传值方法
context.WithValue 能从父上下文中创建一个子上下文，传值的子上下文使用 context.valueCtx 类型：
```
func WithValue(parent Context, key, val interface{}) Context {
	if key == nil {
		panic("nil key")
	}
	if !reflectlite.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}
	return &valueCtx{parent, key, val}
}
```
context.valueCtx 结构体会将除了 Value 之外的 Err、Deadline 等方法代理到父上下文中，它只会响应 context.valueCtx.Value 方法，该方法的实现也很简单：

```
type valueCtx struct {
	Context
	key, val interface{}
}

func (c *valueCtx) Value(key interface{}) interface{} {
	if c.key == key {
		return c.val
	}
	return c.Context.Value(key)
}
```
