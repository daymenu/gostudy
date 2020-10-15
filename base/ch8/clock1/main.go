package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("client is connection: ", conn.RemoteAddr())
		go handelConn(conn)
	}
}

func handelConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("2006-01-02 15:04:05\n"))
		if err != nil {
			fmt.Println("wirte to client is error : ", c.RemoteAddr(), err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
