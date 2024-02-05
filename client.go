package main

import (
	"golang.org/x/net/proxy"
	"net"
	"net/url"
	"time"
)

func main() {
	// 解析代理地址

	u, _ := url.Parse("sock5://127.0.0.1:1080")

	d := &net.Dialer{Timeout: 3 * time.Second}

	socks5, _ := proxy.SOCKS5("tcp", u.Host, nil, d)

	c, err := socks5.Dial("tcp", "127.0.0.1:18081")

	if err != nil {
		panic(err)
	}

	for {

		c.Write([]byte("abc"))

		time.Sleep(1 * time.Second)
	}

}
