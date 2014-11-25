package server

import (
	"fmt"
	"net"
	"strings"
)
// Define the DCS Telemetric Server :)
type DTS struct {
	Adress        net.UDPAddr
	myDispatchers map[string]*DuinoDispatcher
	connection    *net.UDPConn
}


func (server *DTS) StartListening() {
//	initialize some stuff 
	server.myDispatchers = make ( map[string]*DuinoDispatcher)

	server.CreateDispatcher()
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
			a := strings.Split(string(buf[0:n]), ";")//split into several commands
			for _,s := range a{
				c := strings.Split(s,"=")
				d := server.myDispatchers[c[0]]
				if (d == nil){
					//fmt.Println("no dispatcher for ", c[0])
				}else{
					d.Relay(s)
				}
				
			}
			fmt.Println("from address", address, "got message:", string(buf[0:n]))
		}
	}
}
func (server *DTS) CreateDispatcher() {
	d := new(DuinoDispatcher)
	commands := d.GetCommands()
	fmt.Println("Dispatcher accepting :");
	for _,s := range commands{
		server.myDispatchers[s]=d //Assigning this Dispatcher to its commands
		fmt.Println(s);
	}
}












