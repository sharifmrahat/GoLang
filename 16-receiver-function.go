package main

// import "fmt"

// //This is custom type or structure
// type bill struct {
// 	name string
// 	items map[string]float64
// 	tip float64
// }

// //make new bills
// func newBill(name string) bill {
// 	//initial bill object
// 	yourBill := bill{
// 		name: name,
// 		items: map[string]float64{"pie": 3.99, "cake": 5.99},
// 		tip: 0,
// 	}

// 	return yourBill
// }

// //forma the bill | create a receiver function
// //we're limiting this bill only for bill objects
// func (b bill)format()string{
// 	fs := "Bill Breakdown \n"
// 	var total float64 = 0

// 	//list items
// 	for k, v := range b.items {
// 		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
// 		//-25 is used to make the string 25 character long
// 		// Minus (-) is used to add space in right, Plus (+) is used to add space in left
// 		total += v
// 	}

// 	//total
// 	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

// 	return fs
// }

// func main() {
//     myBill := newBill("John's Bill")
//     fmt.Println(myBill.format())
// }

