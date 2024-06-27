package main

import (
	"fmt"
	"time"
)

func main() {
	go printA()
	go printB()
	time.Sleep(time.Second)
}

func printA() {
	for i := 0; i < 5; i++ {
		fmt.Println("A goroutine")
	}
}

func printB() {
	for i := 0; i < 5; i++ {
		fmt.Println("B goroutine")
	}
}
