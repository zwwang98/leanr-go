package main

import "fmt"

func main() {
	declareArray()

	arr := [5]int{1, 2, 3, 4, 5}
	changeArr(&arr)
}

func declareArray() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr1, arr2, arr3)

	// 2d-array
	var grid [4][5]int
	fmt.Println(grid)

	fmt.Printf("\nIterate through the whole array - using for loop:\n")
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("arr3[%d] = %d\n", i, arr3[i])
	}
	fmt.Println()

	fmt.Printf("\nIterate through the whole array - using range.\n" +
		"1. If you want only index      - for i := range arr3 {}\n" +
		"2. If you want index and value - for i, v := range arr3 {}\n" +
		"3. If you want only value      - for _, v := range arr3 {}\n")
	for i, v := range arr3 {
		fmt.Printf("arr3[%d] = %d\n", i, v)
	}
	fmt.Println()
}

/*
数组是值类型
1. argument声明了是[5]int，那就必须是[5]int才是valid argumen，如果传入[3]int会报错
2. 因为是值类型，所以传递数组时都是通过复制传递，所以在拷贝处做的修改不会发生在原数组上
*/
func printArr(arr [5]int) {
	fmt.Printf("\nPrint array:\n")
	for i, v := range arr {
		fmt.Printf("arr[%d] = %d\n", i, v)
	}
	fmt.Println()
}

/*
数组是值类型，为了能在函数里修改愿数组，我们需要传入数组指针，而不是数组本身。
其实我们一般不使用数组。
*/
func changeArr(arr *[5]int) {
	fmt.Printf("\nBefore change:\n")
	for i, v := range arr {
		fmt.Printf("arr[%d] = %d\n", i, v)
	}
	fmt.Println()

	fmt.Printf("\nChange arr[0] to 100\n")
	arr[0] = 100

	fmt.Printf("\nAfter change:\n")
	for i, v := range arr {
		fmt.Printf("arr[%d] = %d\n", i, v)
	}
	fmt.Println()
}
