package main

import (
	"fmt"
	"strings"
)


func getInitials(name string) (string, string){
	upperCase := strings.ToUpper(name)

	names := strings.Split(upperCase, "")

	var initials []string

	for _, value := range names {
		initials = append(initials, value[:1])  //0:1 range. (0 to 1 range)
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"


}

func main() {
	initial1, initial2 := getInitials("John Doe")
	fmt.Println(initial1, initial2)

	name1, name2 := getInitials("Abdullah")

	fmt.Println(name1, name2)
}

