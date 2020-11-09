package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func testMonster() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      40,
		Birthday: "2011-11-1",
		Sal:      8000.0,
		Skill:    "究极进化",
	}

	// 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}

	// 序列化结果
	fmt.Printf("monster 序列化后=%v\n", string(data))
	// fmt.Printf("monster 序列化后=%v\n", data)
}

// 将map进行序列化
func testMap() {
	// 定义一个map
	var a map[string]interface{}

	// 使用map，需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "天"

	// 序列化
	data, err := json.Marshal(&a)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}

	// 序列化结果
	fmt.Printf("monster 序列化后=%v\n", string(data))
}

// 对切片进行序列化
func testSlice() {
	// 定义一个切片
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "红孩儿"
	m1["age"] = 30
	m1["address"] = "天"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "红孩儿"
	m2["age"] = 30
	m2["address"] = [2]string{"一个", "tian"}
	slice = append(slice, m2)

	data, err := json.Marshal(&slice)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}

	// 序列化结果
	fmt.Printf("monster 序列化后=%v\n", string(data))
}
func main() {
	//testMonster()
	//testMap()
	testSlice()
}
