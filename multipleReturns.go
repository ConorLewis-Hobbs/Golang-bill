package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning %v \n", name)
}

func sayGoodbye(name string) {
	fmt.Printf("Goodbye %v \n", name)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
	// take array as argument,
	//call function for each element within an array
}

func areaOfCircle(r float64) float64 {
	return math.Pi * r * r
	//takes
}

func main() {
	sayGreeting("Conor")
	//Good morning Conor

	sayGoodbye("Conor")
	//Goodbye Conor

	cycleNames([]string{"Jim", "Bob", "Carrot"}, sayGreeting)
	// Good morning Jim
	// Good morning Bob
	// Good morning Carrot

	cycleNames([]string{"Jim", "Bob", "Carrot"}, sayGoodbye)
	// Goodbye Jim
	// Goodbye Bob
	// Goodbye Carrot

	areaOne := areaOfCircle(10.5)
	fmt.Println("The area of the circle is:", areaOne)
	//The area of the circle is: 346.36059005827474
	fmt.Printf("The area of the circle is %0.3f \n", areaOne)
	//The area of the circle is 346.361
	//(3 decimal places)

}
