package main

import "fmt"

func updateName(newName *string) {
	*newName = "Bonnie"
}

func pointers() {
	name := "Rick"

	fmt.Println("memory address of name is: ", &name)
	//memory address of name is:  0x14000102020
	fmt.Println(name)
	//Rick

	memoryLocation := &name
	fmt.Println(memoryLocation)
	//0x14000102020

	fmt.Println("value at memory address:", *memoryLocation)
	//Rick

	fmt.Println(name)
	updateName(memoryLocation)
	fmt.Println(name)
	// Rick
	// Bonnie

}
