package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This is custom type or structure
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	//initial bill object
	yourBill := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return yourBill
}

// forma the bill | create a receiver function
// we're limiting this bill only for bill objects
func (b *bill) format() string {
	fs := "Bill Receipt | Customer Name: " + b.name + "\n"
	fs += fmt.Sprintf("%25v \n", "-------------------------------------")
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		//-25 is used to make the string 25 character long
		// Minus (-) is used to add space in right, Plus (+) is used to add space in left
		total += v
	}

	//To add full line -- after the items
	fs += fmt.Sprintf("%25v \n", "-------------------------------------")

	//subtotal
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "Subtotal:", total)
	//tip
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "Tip:", b.tip)

	fs += fmt.Sprintf("%25v \n", "-------------------------------------")
	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "Total:", total+b.tip)

	return fs
}

// update the tip
func (b *bill) updateTip(tip float64) {
	// (*b).tip = tip
	b.tip = tip

	//Note:
	//(*b) is used to access the value of the bill pointer, it is called as dereferencing
	//b is a pointer to the bill object, and *b is the value of the bill object. but we can directly use b.tip = tip
}

// add item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

//Save Bill

func (b *bill) save() {
	//We need to save the data in byte slice format
	data := []byte(b.format())

	//first: directory, second: data, third: permission
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		//Panic to stop the flow of the program/operation
		panic(err)
	}
	fmt.Println("Bill has been save to file")
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput("Choose  option (a - add item, s - save bill, t - add tip) ", reader)

	switch option {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		//strconv is a package to convert string
		p, err := strconv.ParseFloat(price, 64) //64 bit float

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		promptOptions(b)

	case "t":
		tip, _ := getInput("Enter Tip: ", reader)
		t, err := strconv.ParseFloat(tip, 64) //64 bit float

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("Tip added - ", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You saved the file", b.name)
	default:
		fmt.Println("Error! invalid option")
		promptOptions(b)
	}

}

func createBill() bill {
	//Bufio is package for Reader and which will read from the Stdin (terminal)
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// //When user type something and enter then it will read the input
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name) //to trim the space

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func main() {
	// myBill := newBill("John's Bill")

	// myBill.addItem("pie", 3.99)
	// myBill.addItem("cake", 5.99)
	// myBill.addItem("coffee", 2.99)

	// myBill.updateTip(10)
	// fmt.Println(myBill.format())
	myBill := createBill()
	promptOptions(myBill)

	fmt.Println(myBill)

}
