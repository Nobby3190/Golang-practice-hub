package main

import (
	"fmt"
)

func variables() {
	var (
		foo string
		bar string
	)
	var age int = 25
	name := "Tony"

	_, _ = foo, bar                           // 變數不使用
	fmt.Printf("%s's age is %d\n", name, age) // 格式化輸出
}

func constants() {
	const pi = 3.14
	const (
		A = 'A'
		B = 'B'
	)
	const (
		C = "C"
		D
		E
	)
	_ = A
	_ = C
	fmt.Printf("%f\n", pi)

	// iota
	const (
		a = iota
		b = 'b'
		c = iota
		d
	)
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
	)
	type Weekday int
	const (
		sunday Weekday = iota
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	)
	fmt.Printf("numbers are %d, %c, %d, %d\n", a, b, c, d)
	fmt.Println(KB, MB, GB)
	fmt.Println(sunday, monday, tuesday, wednesday, thursday, friday, saturday)
}

func main() {
	variables()
	constants()
}
