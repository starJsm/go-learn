package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	rVal := reflect.ValueOf(b)
	fmt.Println("n2=", rVal)

	n2 := 2 + rVal.Int()
	//n3 := "2" + rVal.String()
	fmt.Println("n2=", n2)
	//fmt.Println("n3=", n3)

	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)

	iV := rVal.Interface()
	num2 := iV.(int)
	fmt.Println("num2=", num2)

}

func reflectTest02(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	rVal := reflect.ValueOf(b)
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T\n", iV, iV)

	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

func main() {
	//var num int = 100
	//var str string = "111"
	//reflectTest01(num)
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
}
