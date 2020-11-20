package main

import (
	"fmt"
)

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("writeData ", i)
	}
	close(intChan)
}

//func readData(intChan chan int, exitChain chan bool) {
//	for {
//		v, ok := <-intChan
//		if !ok {
//			break
//		}
//		fmt.Println("readData ", v)
//	}
//	exitChain <- true
//	close(exitChain)
//}
func readData(intChan chan int, exit *bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readData ", v)
	}
	*exit = true
}

func main() {
	var intChan chan int
	intChan = make(chan int, 30)
	ok := false

	go writeData(intChan)
	go readData(intChan, &ok)

	//exitChain <- true
	//close(exitChain)

	for {
		//v, ok := <-exitChain
		if ok {
			fmt.Println("is ok? ", ok)
			break
		}
		//fmt.Println("value: ", v)
	}

}
