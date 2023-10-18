package main

import "fmt"

func main() {

	trunc := new(float64)
	fmt.Println("CTRL+C to close")

	for {
		fmt.Print("Enter a float to be truncated: ")

		_, err := fmt.Scan(trunc)

		if err != nil {
			error := fmt.Errorf("\nThere was an error with your entry: %v", &trunc)
			fmt.Println(error)
			return
		}

		fmt.Printf("%d is the truncated float \n", int64(*trunc))
	}


}