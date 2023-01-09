package main

/*
*
返回的不仅是return statement里的函数，还有相关的变量。即返回的是一个闭包。
*/
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

//func main() {
//	a := adder()
//	for i := 0; i < 10; i++ {
//		fmt.Printf("The sum from 0 to %d is: %d\n", i, a(i))
//	}
//}
