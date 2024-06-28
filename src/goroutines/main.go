package main

import (
	"fmt"
	"sync"
	"time"
)

/*
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
*/

func findPrimes(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	if num == 1 {
		return
	} else if num == 2 {
		fmt.Println(num)
	} else {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return
			}
		}
		fmt.Println(num)
	}
}

func main() {
	wg := new(sync.WaitGroup)
	jobs := 300000
	wg.Add(jobs)
	start := time.Now().Unix()
	for i := 0; i <= jobs; i++ {
		go findPrimes(i, wg)
	}
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println(end-start, "seconds")
}
