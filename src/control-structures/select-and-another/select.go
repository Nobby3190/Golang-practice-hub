package selectandanother

import (
	"fmt"
	"time"
)

/*
讓外部使用的包要大寫
select用於非同步io操作，如果有多個case可執行，select只會隨機挑一個執行
沒有default，select語句不會停止，直到有符合case語句執行
*/

func SelectCase() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received %v from c1\n", i1)
	case c2 <- i2:
		fmt.Printf("sent %v to c2\n", i2)
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received %v from c3\n", i3)
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

// example two
var resChen = make(chan int)

func SelectCaseTwo() {
	go func() {
		resChen <- 0
	}()

	select {
	case data := <-resChen:
		doData(data)
	case <-time.After(time.Second * 3):
		fmt.Println("request time out")
	}
}

func doData(data int) {
	fmt.Printf("test number: %v\n", data)
}

// example three
var shouldQuit = make(chan struct{})

func SelectCaseThree() {
	go func() {
		simulateError()
	}()

	for {
		fmt.Println("Main program is running...")
		time.Sleep(time.Second)

		select {
		case <-shouldQuit:
			cleanup()
			return
		default:

		}
	}
}

func simulateError() {
	fmt.Println("Simulating error or illegal operation...")
	close(shouldQuit)
}

func cleanup() {
	fmt.Println("Cleaning up the resources.")
}

// example four
func SelectCaseFour() {
	ch := make(chan int, 5)
	for i := 0; i < 10; i++ {
		data := i
		select {
		case ch <- data:
			fmt.Printf("Sent %d to channel\n", data)
		default:
			fmt.Printf("Channel is full, dropped data %d\n", data)
		}
	}

	close(ch)
	for num := range ch {
		fmt.Printf("Received %d from channel\n", num)
	}
}
