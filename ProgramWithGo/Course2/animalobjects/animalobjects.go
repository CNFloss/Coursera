package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal struct {
	food string
	locmotion string
	noise string
}

func (a animal) Eat() {
	fmt.Println(a.food)
}

func (a animal) Move() {
	fmt.Println(a.locmotion)
}

func (a animal) Speak() {
	fmt.Println(a.noise)
}

func verify(err error) {
	if err != nil {
		errorMsg := fmt.Errorf("\nThere was an error with your entry: %s", err)
		fmt.Println(errorMsg)
		os.Exit(1)
	}
}

func checkExit(str string) {
	if str == "x" {
		os.Exit(0)
	}
}

func checkAction(str string, a animal) {
	switch str {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()		
	}
}

func main() {
	fmt.Println("Enter the x symbol to close")

	cow := animal{
		food: "grass",
		locmotion: "walk",
		noise: "moo",
	}
	bird := animal{
		food: "worms",
		locmotion: "fly",
		noise: "peep",
	}
	snake := animal{
		food: "mice",
		locmotion: "slither",
		noise: "hsss",
	}

	for {
		fmt.Print("> ")
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
		verify(err)
		
		line = strings.TrimSpace(line)
		checkExit(line)

		command := strings.Split(line, " ")
		
		switch command[0] {
		case "cow":
			checkAction(command[1], cow)
		case "bird":
			checkAction(command[1], bird)
		case "snake":
			checkAction(command[1], snake)
		}
	}
}