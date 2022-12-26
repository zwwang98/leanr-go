package main

import (
	"fmt"
	"math"
	os "os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// package-scope variables
var (
	aa = 3
	bb = "bb"
	cc = true
)

// https://stackoverflow.com/questions/7052693/how-to-get-the-name-of-a-function-in-go
func GetFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func variableNoInitial() {
	var a int
	var s string
	fmt.Printf("\n[%s] %d %q\n", GetFunctionName(variableNoInitial), a, s)
}

func variableWithInitial() {
	var a int = 3
	var s string = "abc"
	fmt.Printf("\n[%s] %d %q\n", GetFunctionName(variableWithInitial), a, s)
}

func variableTypeDeduction() {
	var a, b, c = 1, "2", true
	fmt.Printf("\n[%s] {a, type: %T, value: %v}, {b, type: %T, value: %q}, {c, type: %T, value: %v}\n", GetFunctionName(variableTypeDeduction), a, a, b, b, c, c)
}

// https://go.dev/ref/spec#Short_variable_declarations
// cannot declare variables like this out of function scope
func variableShortDeclaration() {
	a, b, c := 1, "2", true
	fmt.Printf("\n[%s] {a, type: %T, value: %v}, {b, type: %T, value: %q}, {c, type: %T, value: %v}\n", GetFunctionName(variableShortDeclaration), a, a, b, b, c, c)
}

func consts() {
	// in go language, capital letters mean public, so constants are not given a capital name
	const filename = "abc.txt"
	// const variable could be used as any type, do not need to specify type
	const a, b = 3, 4
	var c int
	// math.Sqrt accept float64 but here is an integer
	// this works
	c = int(math.Sqrt(a*a + b*b))
	fmt.Printf("\n[consts] {filename: %s}, {c: %v}\n", filename, c)
}

func enum() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	fmt.Printf("\n[enum] {cpp: %v}, {java: %v}, {python: %v}, {golang: %v}\n", cpp, java, python, golang)

	// https://go.dev/ref/spec#Iota
	// constant generator
	const (
		aa = iota
		bb
		cc
		dd
	)
	fmt.Printf("\n[enum] test iota {aa: %v}, {bb: %v}, {cc: %v}, {dd: %v}\n", aa, bb, cc, dd)

	// iota could also be used in calculation
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Printf("\n[enum] iota with calculation {b: %v}, {kb: %v}, {mb: %v}, {gb: %v}, {tb: %v}, {pb: %v}\n", b, kb, mb, gb, tb, pb)
}

// show if statements
func readFile(filename string) {
	fmt.Printf("\n[readFile] Contents in %q: ", filename)
	// common way to handle error
	if contents, err := os.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%q\n", contents)
	}
}

// switch
// 1. break is default in every case
func grade(score int) string {
	g := ""

	switch {
	case score < 60:
		g = "F"
	case score < 70:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		// weird for panic method, it seems that panic could only accept type any
		panic(fmt.Sprintf("Wrong score: %d\n", score))
	}
	return g
}

// for loop
func convertToBinary(n int) string {
	binaryStr := ""

	for n > 0 {
		loBit := n % 2
		binaryStr = strconv.Itoa(loBit) + binaryStr
		n /= 2
	}

	return binaryStr
}

func forever() {
	// the same as 	while loop
	i := 10
	for i < 10 {
		fmt.Println(i)
	}

	// without condition, for loop will loop forever
	for {
		fmt.Println("abc")
	}
}

func eval(a int, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s\n", op)
	}
}

//  1. function could have more than one return values
//  2. all return values must be used, otherwise the compiler will report error
//  3. to escape compiler error on the unused one of the two returned values by one function,
//     could name it as "_"
func div(a, b int) (q, r int) {
	return a / b, a % b
}

// https://go.dev/ref/spec#Function_types
// variadic function
// The final incoming parameter in a function signature may have a type prefixed with ....
// A function with such a parameter is called variadic and may be invoked with zero or more arguments for that parameter.
func sum(numbers ...int) int {
	s := 0

	for i := range numbers {
		s += i
	}
	return s
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func swapWithoutPointer(a, b int) (int, int) {
	return b, a
}

//
//func main() {
//	fmt.Println("hello world")
//	variableNoInitial()
//	variableWithInitial()
//	variableTypeDeduction()
//	variableShortDeclaration()
//	fmt.Printf("\n[Print Package-Scope Variables] {aa, type: %T, value: %v}, {bb, type: %T, value: %q}, {cc, type: %T, value: %v}\n", aa, aa, bb, bb, cc, cc)
//	consts()
//	enum()
//	readFile("abc.txt")
//	grade(-1)
//	fmt.Println(convertToBinary(6))
//
//	// to escape error on the unused return value, use _
//	q, _ := div(1213, 13)
//	fmt.Println(q)
//
//	x, y := 4, 6
//	fmt.Printf("Before swap. {x: %v}, {y: %v}\n", x, y)
//	swap(&x, &y)
//	fmt.Printf("Swap once. {x: %v}, {y: %v}\n", x, y)
//	x, y = swapWithoutPointer(x, y)
//	fmt.Printf("Swap twice. {x: %v}, {y: %v}\n", x, y)
//
//	declareArray()
//
//	arr := [5]int{1, 2, 3, 4, 5}
//	changeArr(&arr)
//}
