package main

import "fmt"

func main() {
	age := 50

	fmt.Println(age <= 50)
	//is age less than 50
	fmt.Println(age >= 50)
	//more than 50
	fmt.Println(age == 50)
	//exactly 50
	fmt.Println(age != 50)
	//not 50

	if age < 30 {
		fmt.Println(age, "is less than 30")
	} else if age >= 70 {
		fmt.Println(age, "Age is greater than 70")
	} else {
		fmt.Println("the age is:", age)
	}

	names := []string{"Yoshi", "Peach", "Kooper", "Devin"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue
			// if condition is true, run this
		}
		if index > 2 {
			fmt.Println("Breaking at position", index)
			break
			//if condition is true, break out of loop
		}

		fmt.Printf("The value at postion %v is %v \n", index, value)
		// The value at postion 0 is Yoshi
		// Continuing at pos 1
		// The value at postion 2 is Kooper
		// Breaking at position 3
	}
}
