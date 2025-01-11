package main

// import (
// 	"fmt"
// )

// func main(){
// //#Boolean & Conditionals

// 	age := 18

// 	fmt.Println(age <= 21)
// 	fmt.Println(age >= 21)
// 	fmt.Println(age == 18)
// 	fmt.Println(age != 18)

// 	if age >= 21 {
// 		fmt.Println("Your are mature")
// 	} else if age >= 18 {
// 		fmt.Println("You are 18")
// 	}else {
// 		fmt.Println("Your are below 18")
// 	}

// 	names := [5]string{"Rahim", "Karim", "Rihan", "Saad", "Muaz"}

// 	for index, value := range names {
// 		if index == 1 {
// 			fmt.Println("Skipping/Continuing the position", index)
// 			continue
// 			//if index === 1 then it will return from here, so Karim will not print
// 		}
// 		if index > 2 {
// 			fmt.Println("Breaking at position", index)
// 			break
// 			//if index >= 3 then it will return from here, so Sadd & Muaz will not print
// 		}

// 		fmt.Printf("The value at position %v is %v \n", index, value)
// 	}
// }