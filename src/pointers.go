package main

import (
	"fmt"
	pk "golang_course/src/mypackage"
)

func main() {
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	myPC := pk.PcPublic{Ram: 16, Disk: 200, Brand: "msi"}
	fmt.Println(myPC)

	myPC.PingPublic()

	fmt.Println(myPC)
	myPC.DuplicateRamPublic()

	fmt.Println(myPC)
	myPC.DuplicateRamPublic()

	fmt.Println(myPC)

}
