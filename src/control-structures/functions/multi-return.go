package functions

import "fmt"

func Foo(name string) {
	fmt.Println("HI," + name + "\n")
}

func FooTwo(name1 string, name2 string) {
	fmt.Println("HI," + name1 + ", " + name2 + "\n")
}

func FooThree(name string) string {
	var str = "HI" + name
	return str
}

func FooFour(name string) (str string) {
	str = "HI" + name
	return
}

func FooFive(x, y int) (int, int) {
	return x + y, x - y
}

func Total(x ...int) int {
	var t int
	for _, n := range x {
		t += n
	}
	return t
}
