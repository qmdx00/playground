package main

import (
	"bufio"
	"fmt"
	"golang-tcp-example/codec"
	"golang-tcp-example/config"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", config.ServerAddr)
	if err != nil {
		panic("listen port error")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("[Conn] localAddr=%+v, remoteAddr=%+v\n", conn.LocalAddr(), conn.RemoteAddr())

	for {
		// read from client
		reader := bufio.NewReader(conn)
		_, received, err := codec.Decode(reader)
		if err != nil {
			fmt.Println("recv failed, err:", err)
			break
		}
		fmt.Println("[Received]", received)

		// // write back to client
		// packet, _ := codec.Encode(received)
		// if _, err := conn.Write(packet); err != nil {
		// 	return
		// }
	}
}
