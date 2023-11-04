package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
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
	namesList := make([]Name, 0, 3)
	fmt.Println("Enter the x symbol to close")
	fmt.Println("Enter a filename with first last names on each line")

	for {
		fmt.Print("Filename to read: ")
    reader := bufio.NewReader(os.Stdin)
    filename, err := reader.ReadString('\n')
		verify(filename, err)

		f, err := os.Open(filename[:len(filename)-2])
    if err != nil {
        verify("", err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fslice := make([]string, 0, 20)
			lslice := make([]string, 0, 20)
			line := []rune(scanner.Text())
			toggle := false
			for i := 0; i < len(line); i++ {
				if !toggle {
					if string(line[i]) == " " {
						toggle = true
						continue
					}
					fslice = append(fslice, string(line[i]))
				} else {
					lslice = append(lslice, string(line[i]))
				}
			}
			namesList = append(namesList, Name{fname:strings.Join(fslice[:20], ""), lname:strings.Join(lslice[:20], "")})
		}
		
		if err := scanner.Err(); err != nil {
			verify("", err)
		}

		for _, val := range namesList {
			fmt.Printf("%s ", val.fname)
			fmt.Printf("%s\n", val.lname)
		}
	}
}