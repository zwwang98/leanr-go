### 语法
#### 定义一个`channel`
```go
var c chan int
```


#### channel是一等公民
channel也是一等公民，参数、返回值都可以是channel。比如下面这段函数的参数就有`channel`
```go
func createWorker(id int, c chan int) chan int {
	for {
		// 将channel的值传出来给n
		n := <-c
		fmt.Printf("Woeker %d received %d.\n", id, n)
	}
}
```


#### 声明channel用途（收还是发）
比如下面这段函数返回的`channel`在函数外部就只能用来接受数据，不能够发出数据。
```go
func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Woeker %d received %c.\n", id, <-c)
		}
	}()
	return c
}
```


### buffered channel
如果是 unbuffered channel，发送到channel里的值必须要有人接，不然就会deadlock
```go
func unBufferedChanDemo() {
	c := make(chan int)
	c <- 1
}
```
比如上面这段代码运行的时候就会报错：
```command
fatal error: all goroutines are asleep - deadlock!
```

这就意味着，当我们在一个goroutine发送数据到channel后，我们就要马上切换到另一个goroutine来接受值，尽管goroutine很轻量，但是这样频繁的切换
也是非常低效的，所以我们有了buffered channel。
```go
func bufferedChanDemo() {
c := make(chan int, 3)
c <- 1
c <- 2
c <- 3
}
```
这段代码是可以运行完不报错的。缓冲大小是3，我们也只发了3个数据到channel中。如果再继续发，就会报错了。


### close channel
可以通过`close(c)`来关闭channel，一般是发送方来调用、以提醒接收方不再有数据发送。
接收方可以通过两种方式判断channel是否被发送方关闭:
1. `n, ok := <-c`，用`ok`来判断
2. `for n := range c { ... }`，用`range`


### channel理论基础
参考Communication Sequential Process(CSP)，go的并发模型就是基于这篇论文。

"Don't communicate by sharing memory, share memory by communicating."


### 使用channel来synchronize


### waitgroup



### select
