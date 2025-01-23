package main

import "fmt"

// Any functions / variables / types declared in this file are available to all other files in the same package.

// This is a package-level variable
var packageLevelVariable = "I am a package level variable"

// This is a package-level function
func packageLevelFunction(name string) {
	fmt.Printf("Hello %v! I am a package level function \n", name)
}

func showScore () {
	fmt.Println("Your score is", yourScore)
}