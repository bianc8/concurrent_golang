package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x,y int) {
	x = sum * 4/9 // return values may be named, 
	y = sum - x // if so they are treated as variables,
	return //  defined at the top of the function
	// this is a naked return, should be used only in short functions (harm readibility)
}

/*
bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for uint32, represents a Unicode code point

float32 float64

complex64 complex128
*/
var (
	ToBe   bool = false
	MaxInt uint64 = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10000; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func mra() {
	rand.Seed(time.Now().UnixNano())
	var x int = rand.Intn(100)
	var y int = rand.Intn(100) 
	fmt.Println("x ", x, " y ",y, " add ", add(x, y))

	a, b := swap("hello", "world")
	fmt.Println("a ", a, " b ", b)

	zi, y := split(rand.Intn(100))
	fmt.Println("zi ", zi, " y ", y)

	var c, bo, s = 'c', true, "no!"
	fmt.Println(c, bo, s)

	fmt.Printf("\nType %T Value %v\n", ToBe, ToBe)
	fmt.Printf("Type %T Value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T Value %v\n", z, z)
	
	var in int // vars declared without an explicit init value are given their zero value
	// 0 for numeric types
	// false for boolean type
	// "" for strings
	fmt.Printf("\n%v\n", in)


	// For init statement, condition expr, post statement
	// it will stop once condition evaluates to false
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println("\n")
	// init and post statement are optional
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
	sum = 1
	// without semicolons, it's a while loop in Go
	// without condition it's a while true, infinite loop
	for sum < 1000 {
		sum += sum
	}
	fmt.Print(sum)

	if sum > 1000 {
		fmt.Println(" > 1000")
	} else {
		fmt.Println(" < 1000")
	}

	fmt.Println(Sqrt(2), "\n")

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s\n", os)
	}

	// defer statement defers the execution of a function until the surrounding function returns
	// deferred call's args are evaluated immediately, but the funcition call is not executed until the surrounding function returns
	// deferred calls are pushed onto a stack, LIFO

	// https://tour.golang.org/moretypes/1
}
