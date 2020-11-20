package main

import (
	"fmt"
	"time"
)

func main() {
	intchan := make(chan int, 5)
	for i := 0; i < 5; i++ {
		intchan <- i
	}

	strchan := make(chan string, 3)
	for i := 0; i < 3; i++ {
		strchan <- "string" + fmt.Sprintf("%d", i)
	}

label:
	for {
		select {
		case v := <-intchan:
			fmt.Println("intchan: ", v)
			time.Sleep(time.Millisecond * 100)
		case v := <-strchan:
			fmt.Println("string: ", v)
			time.Sleep(time.Millisecond * 100)
		default:
			fmt.Println("退出")
			// return
			break label
		}
	}
	fmt.Println("out label")
}
