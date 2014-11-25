package server

import (
	"fmt"
	"github.com/tarm/goserial"
	"io"
	"log"
	"strings"
)

type DuinoDispatcher struct {
	serialPort io.ReadWriteCloser
}

func (d *DuinoDispatcher) relay(data []string) (err error) {

	return nil

}

func (d *DuinoDispatcher) Relay(command string) (err error) {

	//fmt.Println("Dispatching ", command)
	c := strings.Split(command, "=")
	f, ok := formatFuncs[c[0]]
	if ok {
		command = f(command)
	} else {
		//fmt.Println("No function defined for type", c[0])
	}

	fmt.Println(command)
	d.sendArduinoCommand(command)
	msg, readErr := d.readArduinoCommand()
	fmt.Println("got message from Arduino: ",msg,readErr)
	return nil

}

// sendArduinoCommand transmits a new command over the nominated serial
// port to the arduino. Returns an error on failure. Each command is
// identified by a single byte and may take one argument (a float).
func (d *DuinoDispatcher) sendArduinoCommand(command string) error {
	if d.serialPort == nil {
		return nil
	}

	// Package argument for transmission
	//	bufOut := new(bytes.Buffer)
	//	err := binary.Write(bufOut, binary.LittleEndian, argument)
	//	if err != nil {
	//		return err
	//	}

	// Transmit command and argument down the pipe.
	fmt.Println("serializing")
	endline := []byte{'\n'}

	for _, v := range [][]byte{[]byte(command), endline} {

		_, err := d.serialPort.Write(v)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	fmt.Println("EISHOKEAY")
	return nil
}

// reads from Ardunio
func (d *DuinoDispatcher) readArduinoCommand( ) (s string, e error) {
	buf := make([]byte, 128)
	_, err := d.serialPort.Read(buf)
	if err != nil {
		return "", err
	}
	sm := string(buf[:])
	return sm, nil
}


// TODO: get the commands from Arduino
func (d *DuinoDispatcher) init() {
	c := &serial.Config{Name: "COM7", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println("serial problem")
		log.Fatal(err)
	}
	d.serialPort = s
	fmt.Println("serial established")
}

// TODO: get the commands from Arduino
func (d *DuinoDispatcher) GetCommands() (data []string) {
	d.init()

	c := []string{"CMSP", "gaga23"}
//	c := []string{"CSMP"}
	//fmt.Printf("get Commands from Dispatcher called\n")
	return c
}

/*
 * Some commands should be reformated before they are sent to arduino.format_CSMP
 * Just declare formating functions and add them to the respective command via the map
 * formatFuncs
 */
type formatFunc func(string) string

func format_CMSP(s string) string {
	s = strings.Replace(s, "-", "", -1)  // remove -------------
	s = strings.Replace(s, "\n", "", -1) // remove all line breaks
	s = strings.Replace(s, "txt_UP", "", -1)
	fmt.Println("index of txt down 1 is now ",  strings.Index(s, "txt_DOWN1"))
	fmt.Println(s)
	for strings.Index(s, "txt_DOWN1") < 21 {//16 plus "CMSP="
		s = strings.Replace(s, "txt_DOWN1", " txt_DOWN1", -1)
		fmt.Println("index of txt down 1 is now ",  strings.Index(s, "txt_DOWN1"))
	}
	s = strings.Replace(s, "txt_DOWN1", "", -1)
	s = strings.Replace(s, "txt_DOWN2", " ", -1)
	s = strings.Replace(s, "txt_DOWN3", " ", -1)
	s = strings.Replace(s, "txt_DOWN4", " ", -1)
	s = strings.Replace(s, "CHAF","CHA",-1) 
	s = strings.Replace(s, "FLAR","FLR",-1)
	s = strings.Replace(s, "OTR1","OT1",-1)
	s = strings.Replace(s, "PROG","PROG",-1)
	return s
}

var formatFuncs = map[string]formatFunc{
	"CMSP": format_CMSP,
}
