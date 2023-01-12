package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service %s: message %d", name, i)
			i++
		}
	}()
	return c
}

func fanInGenerator(chs ...chan string) chan string {
	c := make(chan string)

	for _, ch := range chs {
		// for loop循环可能会改变ch，使得ch最后只是range里最后一个channel的值，所以fan-in看到service2的值
		// 可以通过拷贝一份channel来解决这个问题，这里是通过函数传参来实现一份拷贝，也可以自己手动拷贝
		go func(in chan string) {
			for {
				// 这里的两个箭头意思是，从c1取出数据，然后传入c，如果没有<-c1而是c1，就变成将channel c1传给channel c了
				c <- <-ch
			}
		}(ch)
	}

	return c
}

func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()

	return c
}

/*
*
这里写fan-in的两种方法：
1. 通过一个channel收集所有其他channel的信息，谁先传入先处理谁，可以想像成两条马路汇成一条马路
2. 通过select
两种的区别：
1. 第一个有多少上游channel就要新开多少goroutine，第二个则只需要新开一个goroutine，然后用select
2. 第一个可以有可变参数，第二个参数是固定的，所以在参数长度不确定的时候，用第一个
*/
func main() {
	m1 := msgGen("Service1")
	m2 := msgGen("Service2")
	m := fanInGenerator(m1, m2)
	//m := fanInBySelect(m1, m2)
	for {
		fmt.Println(<-m)
	}
}
