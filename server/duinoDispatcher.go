package server

import (
	"fmt"
)

type DuinoDispatcher struct {
}

func (d *DuinoDispatcher) Relay(data []string) (err error) {

	return nil

}

func (d *DuinoDispatcher) GetCommands() (data []string) {
	c := []string{"ACC00", "HSI0C"}
	fmt.Printf("get Commands from Dispatcher called\n")
	return c
}
