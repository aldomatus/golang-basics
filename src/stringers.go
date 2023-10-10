package main

import "fmt"

// pc represents a personal computer with RAM, brand, and disk storage.
type pc struct {
	ram   int
	brand string
	disk  int
}

// String returns a string representation of the pc struct.
func (myPC pc) String() string {
	return fmt.Sprintf("%s with %d GB of RAM and an SSD of %d GB", myPC.brand, myPC.ram, myPC.disk)
}

func main() {
	myPC := pc{ram: 16, brand: "MacBook", disk: 512}
	fmt.Println(myPC)

	// Another way to create a pc struct
	var anotherPc pc
	anotherPc.brand = "DELL"
	anotherPc.disk = 1000
	anotherPc.ram = 16

	fmt.Println(anotherPc)
}
