package main

import (
	"fmt"
	"strings"
)

func main() {

	str := new(string)
	fmt.Println("CTRL+C to close")

	for {
		fmt.Print("Search a string to for ian: ")

		_, err := fmt.Scan(str)

		if err != nil {
			error := fmt.Errorf("\nThere was an error with your entry: %v", &str)
			fmt.Println(error)
			return
		}

		newStr := strings.ToLower(*str)
		test1 := strings.Index(newStr, "i")
		if test1 != 0 {
			fmt.Println("Not Found!")
			continue
		}
		test2 := strings.Contains(newStr, "a")
		if !test2 {
			fmt.Println("Not Found!")
			continue
		}
		
		if "n" != string(newStr[len(newStr) - 1]) {
			fmt.Println("Not Found!")
			continue
		}
		
		fmt.Println("Found!")
	}


}