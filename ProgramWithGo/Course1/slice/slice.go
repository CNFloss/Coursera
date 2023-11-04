package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	str := new(string)
	intSlice := make([]int, 0, 3)
	current := 0
	fmt.Println("Enter the X key to close")
	fmt.Println("Provide one entry at a time to store and sort integers")

	for {
		fmt.Print("Integer to store and sort: ")

		_, err := fmt.Scan(str)

		if err != nil {
			error := fmt.Errorf("\nThere was an error with your entry: %v", &str)
			fmt.Println(error)
			return
		}

		*str = strings.ToLower(*str)

		if *str == "x" {
			os.Exit(0)
		}

		num, err := strconv.Atoi(*str)

		if err != nil {
			error := fmt.Errorf("\nThere was an error with your entry: %v", &str)
			fmt.Println(error)
			return
		}

		intSlice = append(intSlice, num)
		sort.Slice(intSlice, func(i, j int) bool { return intSlice[i] < intSlice[j] })
		fmt.Print("Sorted Integers: ")
		fmt.Println(intSlice)

		current++
	}


}