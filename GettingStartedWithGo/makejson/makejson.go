package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	Name string `json:"name"`
	Addr string `json:"address"`
}

func verify(str string, err error) {
	if err != nil {
		errorMsg := fmt.Errorf("\nThere was an error with your entry: %s", err)
		fmt.Println(errorMsg)
		os.Exit(1)
	}
	
	if len(str) == 3 && strings.ToLower(string(str[0])) == "x" {
		os.Exit(0)
	}
}

func main() {
	addrBook := make(map[string]string, 50)
	fmt.Println("Enter the ! symbol to close")
	fmt.Println("Enter a name then a address to cache")
	fmt.Println("Max capacity is 50 of each pair")

	for {
		fmt.Print("Enter full name: ")
    reader := bufio.NewReader(os.Stdin)
    fullname, err := reader.ReadString('\n')
		verify(fullname, err)

		fmt.Print("Enter full address: ")
    fulladdr, err := reader.ReadString('\n')
		verify(fulladdr, err)
		fullname = fullname[0:len(fullname)-2]
		fulladdr = fulladdr[0:len(fulladdr)-2]
		fmt.Println(fullname, fulladdr)
		addrBook[fullname] = fulladdr

		for key, val := range addrBook {
			jsonEntry, err := json.Marshal(Entry{key, val})
			if err != nil {
				errorMsg := fmt.Errorf("\nThere was an error with your entry: %s", err)
				fmt.Println(errorMsg)
				return
			}
			strEntry := string(jsonEntry)

			fmt.Println(strEntry)
		}
	}
}