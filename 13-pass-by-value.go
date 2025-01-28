package main

// import "fmt"

// //GO is a pass by value language not pass by reference
// //GO makes "copies" of values when passed into functions

// /*
// Group A: Non pointer values				Group B: Pointer wrapper values
// ---------------------------				-------------------------------
// Strings									Slices
// Ints									Maps
// Floats									Functions
// Booleans
// Arrays
// Structs
// */

// func updateName (x string) string{

// 	fmt.Println("init:", x) // original value which is passed
// 	x = "wedge"
// 	fmt.Println("final:", x) // changed value

// 	return x //return to get changed value
// }

// func updateMenu (y map[string]float64){
// 	y["coffee"] = 2.99 //add/update coffee in the main array
// }

// func main() {
// //Group-A: Examples of group-A: Strings, Ints, Floats, Booleans, Arrays, Structs

// 	name := "tifa"

// 	updateName(name)

// 	fmt.Println("var:", name) //original value, value of name variable will not change

// 	//To change on update we have to re-assign the value of the variable (and must be return the change value from function)

// 	name = updateName(name)

// 	fmt.Println("updated:", name)

// //Group-B: Examples of group-B: Slices, Maps, Functions

// 	menu := map[string]float64 {
// 		"pie": 4.99,
// 		"ice cream": 3.55,
// 	}

// 	updateMenu(menu) //It will update the original maps

// 	fmt.Println(menu)	//print: map[coffee:2.99 ice cream:3.55 pie:4.99]

// 	//In the group-B it creates a copy of the pointer (map) in the memory and also update the reference values in the memory when needed

// }

