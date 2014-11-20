package main

import (
	"fmt"
	"net"
	"time"
)



func main() {
	fmt.Printf("hello, world\n")
	myUDPServer()
}

func myUDPServer() {
	addr := net.UDPAddr{
		Port: 9229,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", &addr)

	defer conn.Close()
	if err != nil {
		panic(err)
	}

	// Do something with `conn`#
	var a1 ArduinoTalker
	a1.init()
	
	var m map[string]ArduinoTalker
	for key, value := range a1.Accepting {
    	m[key] = a1
    }

	//var buf []byte = make([]byte, 64000)
	buf := make([]byte, 65536)

	for {
		time.Sleep(100 * time.Millisecond)
		receiveDcsData()

	}

}

func receiveDcsData() {

		n, address, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("error reading data from connection")
			fmt.Println(err)
			return
		}
		if address != nil {

			fmt.Println("got message from ", address, " with n = ", n)

			if n > 0 {
				fmt.Println("from address", address, "got message:", string(buf[0:n]), n)
			}
		}
}
