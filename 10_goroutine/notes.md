### 了解goroutine
Go的 coroutine 协程
1. 轻量级"线程"
2. 抢占式的线程控制权在于处理器，如果一个线程在任务执行中被处理器打断，会需要存储很多的上下文信息。  
   非抢占式的协程多任务处理，由协程主动交出控制权，会更加高效、对资源的消耗更小一些。
3. 是编译器/解释器/虚拟机层面的多任务，不是操作系统级别的多任务
4. 多个协程可能在一个或者多个处理器上运行

### "非抢占式"多任务处理
1. 非抢占式可能造成死机，因为协程不会主动交出控制权，如果一直没有机会交出控制权的话，可能就会死机。这算是非抢占式的一个坑。
2. go 1.14起开始支持抢占式了，算是填了坑

### 如何定义`goroutine`
1. 任何函数只要加上`go`关键字，这个函数就会被安排在单独的`goroutine`中运行。
2. 不需要像Python一样对异步函数进行区分，都是用`go`声明即可。
3. 调度器会在合适的点自己切换，这和传统的协程不一样，传统的协程是非抢占式的。
4. 可以在运行时加上flag `-race`来检测数据访问冲突。

### `goroutine`可能的切换点
仅供参考，不能保证切换或者不切换。
1. i/o，select
2. channel
3. 等待锁
4. 函数调用（有时候）
5. 调用`runtime.Gosched()`，主动交出控制权    