package main

/*
	This program simulates a DTM M = ({z_0,z_1,z_2}, {a,b}, {a,b,␣}, \delta, z_0, {z_2}),
   	with one tape and prints its contents after each step. The maschine holds for the given
   	input "aabba" this must not be the case for another input.
   	An DTM accepts if it holds or it reaches a finite state, the first condition is not implemented here.
*/

import (
	"fmt"
)

func main() {
	state := 0
	slot := 0
	tape := []rune{'a', 'a', 'b', 'b', 'a', '␣'}
	done := false
	for !done {

		out := ""
		for i, v := range tape {
			if i == slot {
				out += " z_" + string(state+'0') + " " + string(v)
			} else {
				out += string(v)
			}
		}
		fmt.Println("\\vdash " + out)

		switch state {
		case 0:
			// z_0a -> z_1bN
			if tape[slot] == 'a' {
				tape[slot] = 'b'
				state = 1
				// z_0b -> z_1aN
			} else if tape[slot] == 'b' {
				tape[slot] = 'a'
				state = 1
				// z_0␣ -> z_2␣N
			} else if tape[slot] == '␣' {
				state = 2
			}
		case 1:
			// z_1a -> z_0aR
			if tape[slot] == 'a' {
				slot++
				state = 0
				// z_1b -> z_0bR
			} else if tape[slot] == 'b' {
				slot++
				state = 0
				// z_1␣ -> z_0␣R
			} else if tape[slot] == '␣' {
				slot++
				state = 0
			}
		case 2:
			// finite state
			done = true
		}

	}
}
