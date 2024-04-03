package main

import "fmt"

func arrays() {

	//arrays

	var ages [3]int = [3]int{20, 25, 30}
	//[3]int = explicitally state the array will be three
	//items long and only hold intigers

	fmt.Println(ages, len(ages))
	// [20 25 30] 3
	//prints the array and the length of the array

	names := [4]string{"Yoshi", "Peach", "Kooper", "Devin"}
	names[1] = "Bowser"
	//Changes the name at position one, Peach becomes Bowser
	fmt.Println(names, len(names))
	// [Yoshi Bowser Kooper Devin] 2

	//slices

	var scores = []int{100, 200, 300}
	//init the array
	scores[2] = 400
	scores = append(scores, 500)
	//you have to re attitolise the the array when adding / removing stuff

	fmt.Println(scores, len(scores))
	// [100 200 400 500] 4

	//Slice ranges, get a range of elements from a slice and save them to a var

	rangeOne := names[1:3]
	fmt.Println(rangeOne)
	// [Bowser Kooper]

	rangeTwo := names[:3]
	fmt.Println(rangeTwo)
	//[Yoshi Bowser Kooper]
	//[:3] if nothing before the : it starts from the beginning of the array

	rangeThree := names[1:]
	fmt.Println(rangeThree)
	//[Bowser Kooper Devin]
	// x number to the end of the array

	rangeThree = append(rangeThree, "Davos")
	fmt.Println(rangeThree)
	//[Bowser Kooper Devin Davos]
}
