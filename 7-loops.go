package main

// import (
// 	"fmt"
// )

// func main(){

// //#Loop: Go always uses for keyword for any kind of loop

// 	x := 0

// 	//Ex:1 Same as while loop

// 	for x < 5 {
// 		fmt.Println("Value of x is", x)
// 		x++ //without this it will go infinite
// 	}

// 	//Ex:2 Traditional for loop
// 	for i := 0; i<= 10; i++ {
// 		fmt.Println("Value of i is:", i)
// 	}

// 	//Ex:3
// 	peoples := []string{"A", "B", "C", "D"}
// 	for i := 0; i< len(peoples); i++ {
// 		fmt.Println(peoples[i])
// 	}

// 	//Ex:4 Like for in loop
// 	for index, value := range peoples {
// 		fmt.Printf("The value at index %v is %v \n", index, value)
// 	}

// 	//Note if we don't want to use the index then we can put _ otherwise it will throw error
// 	// for _, value := range people{}

// }