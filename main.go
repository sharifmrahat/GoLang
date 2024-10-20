package main

import (
	"fmt"
	"sort"
	"strings"
)

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


//#Array

	// var ages [3]int = [3]int {18, 23, 28}
	var ages = [3]int {18, 23, 28} //should declare type in right

	fmt.Println(ages, len(ages))

	names := [4]string{"Rahim", "Karim", "Rihan", "Saad"}
	names[1] = "Abdullah"

	fmt.Println(names, len(names))

//#Slices

	//slices use arrays under the hood

	var scores = []int{23, 435, 98} //it will create a slices format not an array

	//Note: both in array and slices we can change the any element with index, 
	//In slices we can add or remove element but in array we can not remove or add element

	//to push

	scores = append(scores, 78)

	fmt.Println(scores, len(scores))

//#Ranges

	rangeOne := names[1:3] //from range 1 to 3 || index 1 and 2
	rangeTwo := names[2:] //from 2 to end
	rangeThree := names[:3] //from start to 3

	//we can append to range since it is a slice

	fmt.Println(rangeOne, rangeTwo, rangeThree)


//#The Standard Library

	//Strings package

	greetings := "Good Morning!"

	fmt.Println(strings.Contains(greetings, "Good")) //It will return true/false in the text contains in the string, case sensitive

	fmt.Println(strings.ReplaceAll(greetings, "Morning", "Afternoon")) //replace morning with afternoon

	fmt.Println(strings.ToUpper(greetings)) // to make it uppercase

	fmt.Println(strings.Index(greetings, "Mo")) // to find the index

	fmt.Println(strings.Split(greetings, " ")) // to split the text base on space, it will create an array with split values.

	//Sort

	widths := []int {32, 345, 12, 56, 87, 80, 23, 65, 43, 67, 9, 49}

	sort.Ints(widths) // to sort out the slices of integers, it will change original slices/arrays

	fmt.Println(widths)

	findIndex := sort.SearchInts(widths, 43) //to search the integer, it will return the position of 30

	fmt.Println(findIndex) // if not available then it will return length+1 of the slices/arrays

	employee := []string {"Nizul", "Wayne", "Hasif", "Shamil", "Farah", "Sharif", "Asad", "Faiz", "Hisham"}

	sort.Strings(employee)

	fmt.Println(employee)

	findTextIndex := sort.SearchStrings(employee, "Sharif")

	fmt.Println(findTextIndex)

	
//#Loop: Go always uses for keyword for any kind of loop

	x := 0

	//Ex:1 Same as while loop

	for x < 5 {
		fmt.Println("Value of x is", x)
		x++ //without this it will go infinite
	}

	//Ex:2 Traditional for loopo
	for i := 0; i<= 10; i++ {
		fmt.Println("Value of i is:", i)
	}

	//Ex:3
	peoples := []string{"A", "B", "C", "D"}
	for i := 0; i< len(peoples); i++ {
		fmt.Println(peoples[i])
	}

	//Ex:4 Like for in loop
	for index, value := range peoples {
		fmt.Printf("The value at index %v is %v \n", index, value)
	}

	//Note if we don't want to use the index then we can put _ otherwise it will throw error
	// for _, value := range people{}

}