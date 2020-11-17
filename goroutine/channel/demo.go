/*
使用rand方法创建10个Person结构体实例，放到channel中
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

var (
	firstName  = []string{"赵", "钱", "孙", "李", "周", "吴", "郑", "王"}
	secondName = []string{"金", "木", "水", "火", "土"}
)

func main() {
	var personChan chan Person
	personChan = make(chan Person, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		index := rand.Intn(len(firstName))
		name := firstName[index]
		index = rand.Intn(len(secondName))
		name += secondName[index]

		age := rand.Intn(100)
		personChan <- Person{Name: name, Age: age, Address: "河南"}
	}
	for i := 0; i < 10; i++ {
		per := <-personChan
		fmt.Println(per)
	}
}
