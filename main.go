package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill named: ", reader)

	b := newBill(name)
	fmt.Println("Created new bill named: ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, t - add tip, s - save bill) ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number...")
			promptOptions(b)
		} else {
			b.addItem(name, p)
			fmt.Println("Item added - ", name, price)
			promptOptions(b)
		}

	case "t":
		tip, _ := getInput("Enter tip ammount: (£) ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number...")
			promptOptions(b)
		} else {
			b.updateTip(t)
			fmt.Println("Tip added - ", tip)
			promptOptions(b)
		}

	case "s":
		b.save()
		fmt.Println("You chose to save the bill in file - ", b.name)
	default:
		fmt.Println("Entered invalid option...")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)

	// fmt.Println(mybill)
}

// mybill := newBill("mario's bill")

// mybill.updateTip(10)
// mybill.addItem("onion soup", 4.50)
// mybill.addItem("pot pie", 5.90)
// mybill.addItem("carrot cake", 2.50)
// mybill.addItem("coffee", 1.50)

// fmt.Println(mybill.format())
