package ArduinoTalker

import (

)

type ArduinoTalker struct {
	Accepting []string
}

func (a ArduinoTalker) init() {
	// this will come from the arduino
	Accepting = [3]string
	Accepting[0] = "A"
	Accepting[1] = "B"
	Accepting[2] = "C"
} 


