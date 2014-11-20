package main

import (
	"bytes"
	"encoding/binary"
	"github.com/tarm/goserial"
	"io"
	"log"
	"fmt"
	"time"
)

func main() {
	c := &serial.Config{Name: "COM7", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)
	sendArduinoCommand("on" , 0, s)
	time.Sleep(1 * time.Second)
	msg, err := readArduinoCommand(s)
	fmt.Println("got message from Arduino: ",msg)
	time.Sleep(1 * time.Second)
	sendArduinoCommand("off" , 0, s)
	time.Sleep(1 * time.Second)
	msg, err = readArduinoCommand(s)
	fmt.Println("got message from Arduino: ",msg)
	
	//      for{
	//      	  n, err := s.Write([]byte("test"))
	//      if err != nil {
	//              log.Fatal(err)
	//      }
	//
	//      buf := make([]byte, 128)
	//      n, err = s.Read(buf)
	//      if err != nil {
	//              log.Fatal(err)
	//      }
	//      log.Print("%q", buf[:n])
	//      }

}

// reads from Ardunio
func readArduinoCommand(serialPort io.ReadWriteCloser) (s string, e error) {
	buf := make([]byte, 128)
	_, err := serialPort.Read(buf)
	if err != nil {
		return "", err
	}
	sm := string(buf[:])
	return sm, nil
}

// sendArduinoCommand transmits a new command over the nominated serial
// port to the arduino. Returns an error on failure. Each command is
// identified by a single byte and may take one argument (a float).
func sendArduinoCommand(command string, argument float32, serialPort io.ReadWriteCloser) error {
	if serialPort == nil {
		return nil
	}

	// Package argument for transmission
	bufOut := new(bytes.Buffer)
	err := binary.Write(bufOut, binary.LittleEndian, argument)
	if err != nil {
		return err
	}

	// Transmit command and argument down the pipe.
	endline := []byte{'\n'}
	 
	 for _, v := range [][]byte{ []byte(command), bufOut.Bytes(), endline} {
	//for _, v := range [][]byte{[]byte{command}, bufOut.Bytes()} {
		_, err = serialPort.Write(v)
		if err != nil {
			return err
		}
	}


	return nil
}
