package main

import (
	"fmt"
)

// Player comment
type Player struct {
	name, instrument string
}

func (x *Player) setPlayer(n, i string) {
	// makes itself an available method on variables of type names. it is called using the dot operator

	pf := fmt.Printf

	pf("\tsetPlayer -> before change -> name: %s  instrument: %s\n", x.name, x.instrument)

	x.name = n
	x.instrument = i

	pf("\tsetPlayer -> after change -> name: %s  instrument: %s\n", x.name, x.instrument)
}

func main() {
	pf := fmt.Printf

	beatle := Player{name: "John Lennon", instrument: "Guitar"}

	pf("type: %T, name: %s instrument: %s\n", beatle, beatle.name, beatle.instrument)

	beatle.setPlayer("George Harrison", "Lead Guitar")

	pf("type: %T, name: %s instrument: %s\n", beatle, beatle.name, beatle.instrument)

}
