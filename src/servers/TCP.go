package main

import (
	"net"
	"fmt"
	"encoding/gob"
)
//this package uses encoding/gob that can decode/encode input data.
// Additional variants of decode/encode are in encoding, as example encoding/json.
func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}

func server() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		//accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	//receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Received", msg)
	}
	c.Close()
}

func client() {
	//connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "Hello World!"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}
