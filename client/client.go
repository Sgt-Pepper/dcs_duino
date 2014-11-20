package main

import (
	"fmt"
	"net"

	"time"
)

func main() {
	service := "127.0.0.1:9229"
	fmt.Println("Connecting to server at ", service)
	conn, err := net.Dial("udp", service)

	if err != nil {
		fmt.Println("Could not resolve udp address or connect to it  on ", service)
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to server at ", service)
	defer conn.Close()
	fmt.Println("About to write to connection")
	for {

		time.Sleep(1000 * time.Millisecond)
		n, err := conn.Write([]byte("A=1;B=2;C=3\n"))
		if err != nil {
			fmt.Println("error writing data to server")
			fmt.Println(err)
			return
		}
		if n > 0 {
			fmt.Println("Wrote ", n, " bytes to server")
		}
	}

}
