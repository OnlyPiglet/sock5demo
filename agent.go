package main

import (
	"fmt"
	"github.com/armon/go-socks5"
	"net"
	"time"
)

var server *socks5.Server

func main() {
	// 起一个简单的 socks5 服务
	var err error
	server, err = socks5.New(&socks5.Config{})
	if err != nil {
		panic(err)
	}
	// 不断向 server 发起连接请求, server 的连接池满了之后, 会阻塞在 dial 这一步
	for {
		conn, err := net.Dial("tcp", "127.0.0.1:8988")
		if err != nil {
			continue
		}
		// 连接成功之后, 使用 socks5 库处理该连接
		go handleSocks5(conn)
	}
}

func handleSocks5(conn net.Conn) {
	defer conn.Close()
	_ = conn.SetDeadline(time.Time{})
	// 使用该 socks5 库提供的 ServeConn 方法
	println(conn.LocalAddr().String())
	err := server.ServeConn(conn)
	if err != nil {
		fmt.Println(err)
	}
}
