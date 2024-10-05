package main

import "fmt"

func main(){
//#String Type
	// one way of variable declaration
	var student string = "Abdullah"


	//another way of variable declaration
	var studentTwo = "Rahim"

	//3rd way of variable declaration
	var studentThree string // value is not assigned

	fmt.Println(student, studentTwo, studentThree)

	studentTwo = "Karim"
	studentThree = "Yousha"
	
	//Shorthand
	specialStudent := "Solomon" //can't use this shorthand outside of function

	fmt.Println(student, studentTwo, studentThree, specialStudent)

//#Int bit and float type

	var age int = 22
	var ageTwo = 24
	ageThree := 27

	fmt.Println(age, ageTwo, ageThree)

	//bits & memory
	var num int8 = 25 //8bit integer | Range: -128 to 127
	var numTwo int16 = 128 //16bit integer 
	var numThree uint = 24 //only positive number
	var numFour uint8 = 225 //8bit positive number | Range: 0 to 255

	//Floats
	var score float32 = 223.5 //32bit floating number
	var scoreTwo float64 = 434324.678
	//float only 32 and 64 bit

	println(num, numTwo, numThree, numFour, score, scoreTwo)



//#Printing and format specifier

	//Print: print without breaking line
	fmt.Print("Hello, ")
	fmt.Print("World \n") //\n for new line

	//Println : print with line break
	fmt.Println("My name is", student, "and my age is", age)
	fmt.Println("The End")

	//Printf : for formatting strings, variable interpolation with string
	fmt.Printf("My name is %v and my age is %v \n", student, age) //default formatting
	//%_ is format specifier. _ can be v or other letter
	//%v for variable, %q for quotes

	fmt.Printf("My name is %q and my age is %q \n", student, age) 
	//%q only add quote on string type
	
	//%T for type checking
	fmt.Printf("My variable types are %T and %T \n", student, age)
	
	//%f for print float

	fmt.Printf("Your score is %f \n", score)
	fmt.Printf("Your score is %0.1f \n", score) //for 1 decimal point

	//Sprintf : save formatted strings
	str := fmt.Sprintf("My name is %v and my age is %v \n", student, age)

	fmt.Println("The save str is:", str)

}