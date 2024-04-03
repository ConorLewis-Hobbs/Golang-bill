package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Create a new bill named: ")
	name, _ := reader.ReadString('\n')
	//set user input to name and save it when they start a new line
	//such as on enter
	name = strings.TrimSpace(name)
	//trims white space from user input

	b := newBill(name)
	fmt.Println("Created new bill named: ", b.name)

	return b
}

func main() {
	mybill := createBill()

	fmt.Println(mybill)
}

// mybill := newBill("mario's bill")

// mybill.updateTip(10)
// mybill.addItem("onion soup", 4.50)
// mybill.addItem("pot pie", 5.90)
// mybill.addItem("carrot cake", 2.50)
// mybill.addItem("coffee", 1.50)

// fmt.Println(mybill.format())
