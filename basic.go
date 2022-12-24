package main

import (
	"fmt"
	"math"
	os "os"
	"reflect"
	"runtime"
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

func main() {
	fmt.Println("hello world")
	variableNoInitial()
	variableWithInitial()
	variableTypeDeduction()
	variableShortDeclaration()
	fmt.Printf("\n[Print Package-Scope Variables] {aa, type: %T, value: %v}, {bb, type: %T, value: %q}, {cc, type: %T, value: %v}\n", aa, aa, bb, bb, cc, cc)
	consts()
	enum()
	readFile("abc.txt")
	grade(-1)
}
