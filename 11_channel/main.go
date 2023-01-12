package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Woeker %d received %c.\n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	// 将值传给channel
	for i := 0; i < 10; i++ {
		channels[i] <- i + 'a'
		//channels[i] <- i + 'A'
	}

	for i := 0; i < 10; i++ {
		channels[i] <- i + 'A'
	}

	// 先暂时这么写，使得1、2都能被n接收并打印出来
	time.Sleep(time.Second)
}

func bufferedChanDemo() {
	c := make(chan int, 3)

	go func() {
		for {
			fmt.Printf("Channel received %d\n", <-c)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5

	time.Sleep(time.Millisecond)
}

/*
*
1. 可以close channel
2. 一般是发送方关闭channel来提醒接收方不再会有新的数据
3. 接收方要自己辨别channel是否已经关闭，加上一个ok即可，或者用 range channel
4. 如果channel close、并且接收方没有进行判断，就会一直收到0
*/
func channelCloseDemo() {
	c := make(chan int, 3)

	go func() {
		for {
			n, ok := <-c
			if ok {
				fmt.Printf("Channel received %d\n", n)
			} else {
				break
			}
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5

	close(c)

	c1 := make(chan int, 3)

	go func() {
		for n1 := range c1 {
			fmt.Printf("Channel received %d\n", n1)
		}
	}()

	c1 <- 1
	c1 <- 2
	c1 <- 3
	c1 <- 4
	c1 <- 5

	close(c1)

	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	bufferedChanDemo()
}
