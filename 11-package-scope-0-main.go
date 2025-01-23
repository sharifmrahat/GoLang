package main

import "fmt"

var yourScore = 89
//But if we put this inside main scope then it will not be accessible to other files, should be in the package scope to access

func main() {
	// var yourScore = 89 : can't access from other files if we put this variable here inside main scope

	fmt.Println(packageLevelVariable)

	packageLevelFunction("Gopher")

	showScore()

		//To see the result we've to run both files: go run main.go 11-package-scope-1.go
}
