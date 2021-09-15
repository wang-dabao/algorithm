go中单利模式分为：懒汉式，饿汉式，懒汉加锁式、双重锁、sycn.Once

###懒汉式
每次获取的时候才会去初始化创建对象，缺点就是线程不安全
```go
type singleton struct{}
var ins *singleton

func getSingleton() *singleton {
    if ins == nil {
       return &singleton{}
    }
    return ins
}
```

###饿汉式
提前初始化好了，用到的时候直接返回,缺点就是如果初始化很费劲的话，很耗时
```go
type singleton struct{}
var ins *singleton = &singleton

func getSingleton() *singleton {
	return ins
}
```

###懒汉加锁式
在懒汉式的基础上，解决并发的问题，加锁、缺点是每次加锁需要付出代价
```go
type singleton struct{}
var ins *singleton
var mu sync.Mutex

func getSingleton() *singleton {
	mu.Lock()
	defer mu.Unlock()
	
	if ins == nil {
		ins = &singleton{}
	}
	return ins
}
```

###双重锁
```go
type singleton struct{}
var ins *singleton
var mu sync.Mutex

func getSingleton() *singleton {
	if ins == nil{
		mu.Lock()
        defer mu.Unlock()

        if ins == nil {
            ins = &singleton{}
        }
	}
    return ins
}
```
###sync.Once
```go
type singleton struct{}
var ins *singleton
var once sync.Once

func getSingleton() *singleton {
     once.Do(func(){
        ins = &singleton{}
     })
     return ins
}
```
