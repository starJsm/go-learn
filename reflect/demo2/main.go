package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func (s Monster) Print() {
	fmt.Println("-----start-----")
	fmt.Println(s)
	fmt.Println("-----end-----")
}

func TestSturct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()

	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)

	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))

		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("field %d: tag is %v\n", i, tagVal)
		}
	}

	numOfMetod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMetod)

	val.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println("res= ", res[0].Int())

	for i := 0; i < len(res); i++ {
		fmt.Printf("res[%d]: %v\n", i, res[i])
	}
}

func main() {
	var a Monster = Monster{
		Name:  "黄鼠狼",
		Age:   400,
		Score: 30.8,
	}
	TestSturct(a)
}
