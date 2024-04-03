package main

import "fmt"

func loops() {
	// x := 0
	// for x <= 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	// while x is >= 5 Println

	// for i := 0; i <= 5; i++ {
	// 	fmt.Println("value of x is:", i)
	// }

	// while x is >= 5 Println

	names := []string{"Yoshi", "Peach", "Kooper", "Devin"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println("The name is:", names[i])
	// }

	//while name index is less than names.length, print out name

	for index, value := range names {
		fmt.Printf("The position at index %v is %v \n", index, value)
		//The position at index 0 is Yoshi
		// The position at index 1 is Peach
		// The position at index 2 is Kooper
		// The position at index 3 is Devin
	}
	// get index and value for names.length

	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
		// The value is Yoshi
		// The value is Peach
		// The value is Kooper
		// The value is Devin
	}
	// get value for names.length

}
