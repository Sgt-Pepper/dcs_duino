package server

import (
	"fmt"
	"net"
)
// Define the DCS Telemetric Server :)
type DTS struct {
	Adress        net.UDPAddr
	myDispatchers map[string]int
	connection    *net.UDPConn
}


func (server *DTS) StartListening() {
	conn, err := net.ListenUDP("udp", &server.Adress)
	server.connection = conn
	defer server.connection.Close()
	if err != nil {
		panic(err)
	}
	for {
		server.receiveDcsData()
	}
}

func (server *DTS) receiveDcsData() {
	buf := make([]byte, 65536)
	n, address, err := server.connection.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("error reading data from connection")
		fmt.Println(err)
		return
	}
	if address != nil {
		fmt.Println("got message from ", address, " with n = ", n)
		if n > 0 {
			fmt.Println("from address", address, "got message:", string(buf[0:n]))
		}
	}
}
func (server *DTS) CreateDispatcher() {
	d := new(DuinoDispatcher)
	commands := d.GetCommands()
	fmt.Println("Dispatcher accepting :");
	for _,s := range commands{
		fmt.Println(s);
	}
}












