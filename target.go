package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {
	p := flag.String("p", "18081", "listen port")

	flag.Parse()

	fmt.Println(*p)

	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:"+*p)

	tcp, _ := net.ListenTCP("tcp", addr)

	for {

		acceptTCP, _ := tcp.AcceptTCP()

		println(acceptTCP.RemoteAddr().String())

		go ReadMsg(acceptTCP)

		time.Sleep(1 * time.Second)

	}

}

func ReadMsg(con *net.TCPConn) {

	for {
		rb := make([]byte, 4096)
		con.Read(rb)
		s := string(rb)
		fmt.Println(s)
		time.Sleep(1 * time.Second)
	}

}
