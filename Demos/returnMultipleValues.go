package main

import (
	"fmt"
	"strings"
)

func getInitials(name string) (string, string) {
	upperCaseName := strings.ToUpper(name)
	// convert input to upper case
	names := strings.Split(upperCaseName, " ")
	// split the input at a space " "

	var initials []string
	// make an array to put the split input
	for _, value := range names {
		initials = append(initials, value[:1])
	}
	// cycle through names, dont need index, do need value
	//for each name, get first letter, put it in initals array

	if len(initials) > 1 {
		return initials[0], initials[1]
		// if ititals array is more than 1 long, return both

	} else {
		return initials[0], "_"
		// else return the first and add a "_"
	}
}

func returnMultiple() {
	fn1, sn1 := getInitials("conor hobbs")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("david tennet")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("Prince")
	fmt.Println(fn3, sn3)
	// C H
	// D T
	// P _
}
