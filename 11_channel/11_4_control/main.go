package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string, done chan struct{}) chan string {
	fmt.Printf("[msgGen] [msgGen-%s] starts...\n", name)
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
				c <- fmt.Sprintf("[msgGen] service-{%s}: message %d", name, i)
				i++
			case <-done:
				fmt.Printf("[msgGen] Cleaning up [msggen-%s]...\n", name)
				// 睡2秒模拟关闭耗时
				time.Sleep(2 * time.Second)
				// 通过done通知main真的完成了清理
				done <- struct{}{}
				return
			}
		}
	}()
	return c
}

/*
*
 1. 用bool来标记是否收到数据
 2. 如果从c收到数据，返回数据和true
    如果未从c收到数据，返回空值和false
*/
func nonBlockingWait(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout):
		return "", false
	}
}

func main() {
	done := make(chan struct{})

	m1 := msgGen("Service1", done)
	m2 := msgGen("Service2", done)

	for i := 0; i < 10; i++ {
		fmt.Println(<-m1)
		if m, ok := nonBlockingWait(m2); ok {
			fmt.Println(m)
		} else {
			fmt.Println("no message from service2")
		}
	}

	for i := 0; i < 10; i++ {
		if m, ok := timeoutWait(m2, 1000*time.Millisecond); ok {
			fmt.Println(m)
		} else {
			fmt.Println("service2 timeout")
		}
	}

	done <- struct{}{}
	<-done
}
