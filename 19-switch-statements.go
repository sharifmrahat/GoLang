package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// // This is custom type or structure
// type bill struct {
// 	name  string
// 	items map[string]float64
// 	tip   float64
// }

// // make new bills
// func newBill(name string) bill {
// 	//initial bill object
// 	yourBill := bill{
// 		name:  name,
// 		items: map[string]float64{},
// 		tip:   0,
// 	}

// 	return yourBill
// }

// // forma the bill | create a receiver function
// // we're limiting this bill only for bill objects
// func (b *bill) format() string {
// 	fs := "Bill Breakdown \n"
// 	var total float64 = 0

// 	//list items
// 	for k, v := range b.items {
// 		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
// 		//-25 is used to make the string 25 character long
// 		// Minus (-) is used to add space in right, Plus (+) is used to add space in left
// 		total += v
// 	}

// 	//To add full line -- after the items
// 	fs += fmt.Sprintf("%25v \n", "-------------------------------------")

// 	//subtotal
// 	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "Subtotal:", total)
// 	//tip
// 	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "Tip:", b.tip)

// 	fs += fmt.Sprintf("%25v \n", "-------------------------------------")
// 	//total
// 	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "Total:", total+b.tip)

// 	return fs
// }

// // update the tip
// func (b *bill) updateTip(tip float64) {
// 	// (*b).tip = tip
// 	b.tip = tip

// 	//Note:
// 	//(*b) is used to access the value of the bill pointer, it is called as dereferencing
// 	//b is a pointer to the bill object, and *b is the value of the bill object. but we can directly use b.tip = tip
// }

// // add item to the bill
// func (b *bill) addItem(name string, price float64) {
// 	b.items[name] = price
// }

// func getInput(prompt string, r *bufio.Reader) (string, error) {
// 	fmt.Print(prompt)
// 	input, err := r.ReadString('\n')

// 	return strings.TrimSpace(input), err
// }

// func promptOptions(b bill) {
// 	reader := bufio.NewReader(os.Stdin)
// 	option, _ := getInput("Choose  option (a - add item, s - save bill, t - add tip) ", reader)

// 	switch option {
// 	case "a":
// 		name, _ := getInput("Item name: ", reader)
// 		price, _ := getInput("Item price: ", reader)
// 		fmt.Println(name, price)
// 	case "t":
// 		tip, _ := getInput("Enter Tip: ", reader)
// 		fmt.Println(tip)
// 	case "s":
// 		fmt.Println("You choose s")
// 	default:
// 		fmt.Println("Error! invalid option")
// 		promptOptions(b)
// 	}

// }

// func createBill() bill {
// 	//Bufio is package for Reader and which will read from the Stdin (terminal)
// 	reader := bufio.NewReader(os.Stdin)

// 	// fmt.Print("Create a new bill name: ")
// 	// //When user type something and enter then it will read the input
// 	// name, _ := reader.ReadString('\n')
// 	// name = strings.TrimSpace(name) //to trim the space

// 	name, _ := getInput("Create a new bill name: ", reader)

// 	b := newBill(name)
// 	fmt.Println("Created the bill - ", b.name)

// 	return b
// }

// func main() {
// 	// myBill := newBill("John's Bill")

// 	// myBill.addItem("pie", 3.99)
// 	// myBill.addItem("cake", 5.99)
// 	// myBill.addItem("coffee", 2.99)

// 	// myBill.updateTip(10)
// 	// fmt.Println(myBill.format())
// 	myBill := createBill()
// 	promptOptions(myBill)

// 	fmt.Println(myBill)

// }
