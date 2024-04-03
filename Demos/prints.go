package main

import "fmt"

func prints() {

	// print statements

	fmt.Println("Hello, world", "Beep boop")
	//starts a new line automatically

	fmt.Print("Hello, world", "Beep boop")
	//Doesn't start a new line

	fmt.Print("\nHello, world \n", "Beep boop \n")
	// \n starts a new line

	age := 27
	name := "Conor"

	fmt.Println("Hello my name is", name, "and I am", age, "years old")

	fmt.Printf("Hello my name is %q and I am %v years old\n", name, age)
	// %v = a format specifier
	// %q = adds quotes around string variables

	fmt.Printf("Age is of type %T \n", age)
	// %T = adds the type of variable

	fmt.Printf("Putting a decimal here: %0.2f \n", 225.25)
	// %f = adds a float variable, 1f = 1decimal place, 2f = 2dp ect

	var str = fmt.Sprintf("Hello my name is %q and I am %v years old\n", name, age)
	fmt.Println(str)
	// save a print statement to a variable
}
