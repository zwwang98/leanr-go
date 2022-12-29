package main

import "fmt"

func main() {
	declareSlice()

	arr := [7]int{0, 1, 2, 3, 4, 5, 6}

	// this is a way to get the slice of the array
	s := arr[:]

	updateSlice(s)

	reslice(s)

	extendSlice()

	sliceUnderTheHood()

	sliceAppend()

	commonSliceOperations()
}

/*
 1. 区间左开右闭
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
    s := arr[2:6]  // s = {2, 3, 4, 5}

2.
*/
func declareSlice() {
	fmt.Println("\n[declareSlice]")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])
}

/*
 1. Do not have to use a pointer to be able to update the slice.
    Pass the slice itself is enough.
 2. The slice is a view of the array under the hood, so any change happens on the slice happens on the array.
*/
func updateSlice(s []int) {
	fmt.Println("\n[updateSlice]")
	fmt.Println("Before update:")

	for i, v := range s {
		fmt.Printf("{i: %d}, {v: %d}\n", i, v)
	}

	fmt.Println("Change s[0] to 100")
	s[0] = 100

	fmt.Println("After update:")
	for i, v := range s {
		fmt.Printf("{i: %d}, {v: %d}\n", i, v)
	}
}

/*
Re-slice happens on the current version of slice.
So if at first s []int = {0, 1, 2, 3, 4, 5, 6}
after s[2:], it will be s = {2, 3, 4, 5, 6}
after s[:2], it will be s = {2, 3}
*/
func reslice(s []int) {
	fmt.Println("[reslice]")

	fmt.Println("Before reslice:")
	for _, v := range s {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	s = s[2:]
	fmt.Println("After reslice s = s[2:]:")
	for _, v := range s {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	s = s[:2]
	fmt.Println("After reslice s = s[:2]:")
	for _, v := range s {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

// Check /Users/CS/learn-go/extend-slice.png
// slice could only extend to the bigger index, extending to the smaller index is not allowed
func extendSlice() {
	fmt.Println("\n[extendSlice]")
	arr := [7]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("arr =", arr)
	s := arr[2:6]
	fmt.Println("s = arr[2:6] =", s)
	extendS := s[3:5]
	fmt.Println("extendS := s[3:5] =", extendS)
}

/*
Each slice has
1. a pointer to where it starts at the array under the hood
2. a length indicates how long the slice is
3. a capacity indicates how many places it has from the starting point to the end of the array under the hood
*/
func sliceUnderTheHood() {
	fmt.Println("\n[sliceUnderTheHood]")
	arr := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr =", arr)
	s := arr[2:6]
	fmt.Printf("s = arr[2:6] = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
	extendS := s[3:5]
	fmt.Printf("extendS := s[3:5] = %v, len(extendS) = %d, cap(extendS) = %d\n", extendS, len(extendS), cap(extendS))
}

/*
 */
func sliceAppend() {
	fmt.Println("\n[sliceAppend]")
	arr := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr =", arr)
	s := arr[2:6]
	fmt.Printf("s = arr[2:6] = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
	extendS := s[3:5]
	fmt.Printf("extendS := s[3:5] = %v, len(extendS) = %d, cap(extendS) = %d\n", extendS, len(extendS), cap(extendS))
	fmt.Println("arr =", arr)

	fmt.Println("If the slice has enough capacity, the append will modify the value in the array under the hood")
	s2 := append(extendS, 10)
	fmt.Printf("s2 := append(extendS, 10) = %v, len(s2) = %d, cap(s2) = %d\n", s2, len(s2), cap(s2))
	fmt.Println("arr =", arr)

	fmt.Println("If the slice has used up capacity, it will be assigned a new longer array.")
	s3 := append(s2, 11)
	fmt.Printf("s3 := append(s2, 11) = %v, len(s3) = %d, cap(s3) = %d\n", s3, len(s3), cap(s3))
	fmt.Println("arr =", arr)
}

func commonSliceOperations() {
	fmt.Println("\n[commonSliceOperations]")
	var s []int
	fmt.Println("Declare an empty slice by \"var s []int\", and s =", s)

	fmt.Println("Zero value of slice is nil")
	fmt.Println("(s == nil) =", s == nil)

	fmt.Println("Append odd numbers to the empty slice")
	for i := 0; i < 10; i++ {
		s = append(s, 2*i+1)
		printSlice(s)
	}
	fmt.Println("s =", s)

	fmt.Println("\nOther ways to declare slices:")
	s1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("1. s1 := []int{0, 1, 2, 3, 4}")
	printSlice(s1)

	s2 := make([]int, 15)
	fmt.Println("2. s2 := make([]int, 15), 15 is the length of the slice")
	printSlice(s2)

	s3 := make([]int, 10, 32)
	fmt.Println("3. s3 := make([]int, 10, 32), 10 is the length of the slice, 32 is the capacity")
	printSlice(s3)

	copy(s2, s1)
	fmt.Println("\ncopy(s2, s1)\n" +
		"s2 is the destination slice, s1 is the source slice")
	printSlice(s1)
	printSlice(s2)

	fmt.Println("\nTo delete the element at s2[3]")
	fmt.Println("Before deletion:")
	printSlice(s2)

	fmt.Println("Delete s2[3] by \"s2 = append(s2[:3], s2[4:]...)\"")
	s2 = append(s2[:3], s2[4:]...)

	fmt.Println("After deletion at s2[3]:")
	printSlice(s2)

	fmt.Println("\nDelete from front:")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println("front =", front)
	printSlice(s2)

	fmt.Println("\nDelete from tail:")
	tail := s2[len(s2)-1]
	fmt.Println("tail =", tail)
	s2 = s2[:len(s2)-1]
	printSlice(s2)
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cpa=%d, s=%v\n", len(s), cap(s), s)
}
