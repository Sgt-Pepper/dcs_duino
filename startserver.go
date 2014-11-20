package main

import (
	"github.com/Sgt-Pepper/dcs_duino/server"
	"fmt"
	"net"
)
func main() {
	fmt.Printf("new architecture\n")

	s := server.DTS{
		Adress: net.UDPAddr{
			Port: 9229,
			IP:   net.ParseIP("127.0.0.1"),
		},
	}
	s.CreateDispatcher()
	s.StartListening()

}

