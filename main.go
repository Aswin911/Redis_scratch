package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("Server is Listening on Port : 6379")

	//creating a new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}

	//Listen for connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		//read the message from the client
		_, err = conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}

		//ignoring request and sending back a PONG
		conn.Write([]byte("+OK\r\n"))
	}
}
