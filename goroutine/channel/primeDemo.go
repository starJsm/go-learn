package main

import (
	"fmt"
	"time"
)

func primeNum(intchan chan int, primechan chan int, exitchan chan bool) {
	fmt.Println("in primeNum")
	var flag bool

	for {
		flag = true
		num, ok := <-intchan
		if !ok {
			break
		}
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primechan <- num
			fmt.Println("put primechan")
		}
	}

	exitchan <- true
}

func main() {
	var intchan chan int
	intchan = make(chan int, 1000)
	var exitchan chan bool
	exitchan = make(chan bool, 4)
	var primechan chan int
	primechan = make(chan int, 200)
	start := time.Now().Unix()
	go func(intchan chan int) {
		for i := 2; i <= 800000; i++ {
			intchan <- i
		}
		close(intchan)
		fmt.Println("close intchan")
	}(intchan)

	fmt.Println(len(intchan))

	for i := 0; i < 6; i++ {
		fmt.Println("in go primeNum")
		go primeNum(intchan, primechan, exitchan)
	}

	go func() {
		fmt.Println("in go func")
		for i := 0; i < 6; i++ {
			<-exitchan
		}
		end := time.Now().Unix()
		fmt.Println("time :", end-start)
		close(primechan)
	}()

	fmt.Println(len(primechan))
	primeCnt := 0
	for {
		_, ok := <-primechan
		if !ok {
			break
		}
		//fmt.Println(primeCnt, v)
		primeCnt++
	}

}
