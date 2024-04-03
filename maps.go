package main

import (
	"fmt"
)

func main() {

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	//map[pie:7.99 salad:6.99 soup:4.99 toffee pudding:3.55]

	fmt.Println(menu["pie"])
	//	7.99

	for key, value := range menu {
		fmt.Printf("%v - £%v \n", key, value)
	}
	// salad - £6.99
	// toffee pudding - £3.55
	// soup - £4.99
	// pie - £7.99

	phoneBook := map[int]string{
		07576603300: "Conor",
		85746603440: "Jack",
		07576577472: "Will",
	}

	fmt.Println(phoneBook)
	//map[1039859514:Will 1039861440:Conor 85746603440:Jack]

	fmt.Println(phoneBook[85746603440])
	//Jack

	phoneBook[85746603440] = "Donnie"

	fmt.Println(phoneBook[85746603440])
	//Donnie

}
