package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list items

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v £%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v £%0.2f \n", "tip:", b.tip)
	//add tip

	fs += fmt.Sprintf("%-25v £%0.2f \n", "total:", total+b.tip)
	//total

	return fs

	// Bill breakdown:
	// Pie:                      £5.99
	// Cake:                     £3.99
	// total:                    £9.98

}

func (b *bill) updateTip(tip float64) {
	(b).tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}