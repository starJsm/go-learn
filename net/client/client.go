package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("cleat start......")

	conn, err := net.Dial("tcp", "10.12.37.157:8888")
	if err != nil {
		fmt.Println("conn fail err: ", err)
		return
	}
	fmt.Println("conn success: ", conn)

	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入[终端]

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ReadString err: ", err)
	}

	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("send message fail")
	}
	fmt.Printf("send %d data\n", n)
}
