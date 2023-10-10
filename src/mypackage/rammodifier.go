package mypackage

import "fmt"

// PcPublic represents a publicly accessible pc
type PcPublic struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPC PcPublic) PingPublic() {
	fmt.Println(myPC.Brand, "Pong!")
}

func (myPC *PcPublic) DuplicateRamPublic() {
	myPC.Ram = myPC.Ram * 2
}
