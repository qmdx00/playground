package main

import (
	"bufio"
	"golang-tcp-example/codec"
	"golang-tcp-example/config"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", config.ServerAddr)
	if err != nil {
		panic("dial tcp error")
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputStr := strings.TrimSpace(input)
		if strings.ToUpper(inputStr) == "Q" {
			return
		}

		// write to server
		packet, _ := codec.Encode(inputStr)
		if _, err := conn.Write(packet); err != nil {
			return
		}

		// // read from server
		// reader := bufio.NewReader(conn)
		// _, received, err := codec.Decode(reader)
		// if err != nil {
		// 	fmt.Println("recv failed, err:", err)
		// 	break
		// }
		// fmt.Println("[Read]", received)
	}
}
