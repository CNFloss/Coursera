package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func GenDisplaceFn(a, v, s float64) func(t float64) float64 {
	// s = ½ a t2 + vot + so
	return func(t float64) float64 {
		return 0.5*a*(t*t) + (v*t) + s
	}
}

func main() {
	fmt.Println("Enter the x symbol to close")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter acceleration: ")
	accel, err := reader.ReadString('\n')
	accel = strings.TrimSpace(accel)
	verify(accel, err)
	a, err := strconv.ParseFloat(string(accel), 64)
	verify("", err)

	fmt.Print("Enter initial velocity: ")
	veloc, err := reader.ReadString('\n')
	veloc = strings.TrimSpace(veloc)
	verify(veloc, err)
	v, err := strconv.ParseFloat(string(veloc), 64)
	verify("", err)
	
	fmt.Print("Enter initial displacement: ")
	disp, err := reader.ReadString('\n')
	disp = strings.TrimSpace(disp)
	verify(disp, err)
	s, err := strconv.ParseFloat(string(disp), 64)
	verify("", err)

	displacementFunc := GenDisplaceFn(a, v, s)

	fmt.Print("Enter time: ")
	time, err := reader.ReadString('\n')
	time = strings.TrimSpace(time)
	verify(time, err)
	t, err := strconv.ParseFloat(string(time), 64)
	verify("", err)

	fmt.Println(displacementFunc(t))
}