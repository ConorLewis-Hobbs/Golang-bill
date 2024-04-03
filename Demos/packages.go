package main

import (
	"fmt"
	"sort"
	"strings"
)

func packages() {
	greeting := "Hello world"

	fmt.Println(strings.Contains(greeting, "Hello"))
	//true

	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	//Hi world
	// Doesn't change the orginal value of the altered tring

	fmt.Println(strings.ToUpper(greeting))
	// HELLO WORLD

	fmt.Println(strings.Index(greeting, "ll"))
	//2
	//returns the position of the queried item (Hello world)

	fmt.Println(strings.Split(greeting, " "))
	//[Hello world]
	//returns an array at the defined split point

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)
	//[20 25 30 35 45 50 60 75]
	//sorts from high to low. Alteres the orginal slice

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)
	//2
	//returns the position of the searched item, 30, [20 25 30 35 45 50 60 75]

	names := []string{"Yoshi", "Peach", "Kooper", "Devin"}
	sort.Strings(names)
	fmt.Println(names)
	// [Devin Kooper Peach Yoshi]
	//orders alphabetically

	fmt.Println(sort.SearchStrings(names, "Kooper"))
	//1
	//returns the index of the searched name, [Devin Kooper Peach Yoshi]
	//
}
