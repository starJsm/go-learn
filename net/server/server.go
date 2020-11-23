package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 100)
		fmt.Printf("server wait client:%s's message\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil || err == io.EOF {
			fmt.Println("server read err: ", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	fmt.Println("server start listening......")

	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err: ", err)
		return
	}
	defer listen.Close()

	// 循环等待客户端连接
	for {
		fmt.Println("wating clent......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("connection err: ", err)
			return
		} else {
			fmt.Printf("Accept() suc conn=%v, ip=%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}
