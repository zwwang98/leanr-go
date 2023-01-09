package main

import (
	"bufio"
	"fmt"
	"learn-go/7_functional_programming/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
	defer fmt.Println(5)
	defer fmt.Println("Defer statements are in a stack:")

	for i := 6; i < 10; i++ {
		fmt.Println(i)
	}
}

func writeFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		//
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Println(pathError.Err)
		} else {
			fmt.Println("Unknown error", err)
		}
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

// 可以在handler上加一层wrapper来统一所有错误的处理

func main() {
	writeFile("fib.txt")
}
