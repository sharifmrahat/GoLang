package main

// import "fmt"

// //We can manually create pointers for the data types like strings, ints, arrays, etc.
// //Pointers are data type themselves, when we create them they're stored in their own memory block
// //We can create a pointer for any variable which will point to the memory location the that variable

// func updateName (x string) string{

// 	fmt.Println("init:", x) // original value which is passed
// 	x = "wedge"
// 	fmt.Println("final:", x) // changed value

// 	return x //return to get changed value
// }

// func updateNameWithPointer (x *string){
// 	*x = "tada"
// }

// func main() {
// 	name := "tifa"

// 	updateName(name)

// //#Pointers

// 	fmt.Println("memory address of variable name is:", &name) //&variable will show the memory location

// 	//We can store pointers in a variable

// 	//Create a pointer
// 	memo := &name //variable memo is storing the pointer to name memory address where name's value ("tifa") is stored

// 	fmt.Println("memory address:", memo) //it will print the same memory address

// 	fmt.Println("value at memory address is:", *memo) //Using * with pointer will print the value of pointing address

// 	/*
// 	|--- memory ------|--- name ---|--- memo ----|
// 	|-----------------|------------|-------------|
// 	|--- address -----|--- 0x001---|--- 0x002 ---|
// 	|-----------------|------------|-------------|
// 	|--value/pointer--|--- tifa ---|--- p0x001---|
// 	|-----------------|------------|-------------|
// 	*/

// 	fmt.Println(name) // original value
// 	updateNameWithPointer(memo)
// 	fmt.Println(name) // changed value by pointer

// }

