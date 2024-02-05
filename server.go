package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

func main() {
	// 使用两个 channel 来暂存 agent 和 user 的连接请求
	userConnChan := make(chan net.Conn, 10)
	agentConnChan := make(chan net.Conn, 10)
	// 监听 agent 服务端口
	go ListenService(agentConnChan, "127.0.0.1:8988")
	// 监听 user 服务端口
	go ListenService(userConnChan, "127.0.0.1:1080")
	for agentConn := range agentConnChan {
		userConn := <-userConnChan
		go copyConn(userConn, agentConn)
	}
}

func ListenService(c chan net.Conn, ListenAddress string) {
	listener, err := net.Listen("tcp", ListenAddress)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		c <- conn
	}
}

func copyConn(srcConn, dstConn net.Conn) {
	_ = srcConn.SetDeadline(time.Time{})
	_ = dstConn.SetDeadline(time.Time{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer srcConn.Close()
		defer dstConn.Close()
		_, err := io.Copy(srcConn, dstConn)
		if err != nil {
			return
		}
	}()
	go func() {
		defer wg.Done()
		defer dstConn.Close()
		defer srcConn.Close()
		_, err := io.Copy(dstConn, srcConn)
		if err != nil {
			return
		}
	}()
	wg.Wait()
}
