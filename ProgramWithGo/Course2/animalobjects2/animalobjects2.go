package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("chirp")
}

type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func createAnimal(animalType, name string) Animal {
	switch animalType {
	case "cow":
		return Cow{name: name}
	case "bird":
		return Bird{name: name}
	case "snake":
		return Snake{name: name}
	}
	return nil
}

func checkExit(str string) {
	if str == "x" {
		os.Exit(0)
	}
}

func main() {
	animals := make(map[string]Animal)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the x symbol to close")
	fmt.Println("Enter commands: ")

	for {
		fmt.Print("> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		checkExit(line)
		words := strings.Fields(line)

		if len(words) != 3 {
			fmt.Println("Invalid input. Please enter commands in the format 'newanimal [name] [type]' or 'query [name] [action]'.")
			continue
		}

		command, name, arg := words[0], words[1], words[2]

		switch command {
		case "newanimal":
			animal := createAnimal(arg, name)
			if animal == nil {
				fmt.Println("Invalid animal type.")
			} else {
				animals[name] = animal
				fmt.Println("Created it!")
			}
		case "query":
			animal, ok := animals[name]
			if !ok {
				fmt.Println("No animal with that name exists.")
				continue
			}
			switch arg {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid query. Valid actions are 'eat', 'move', and 'speak'.")
			}
		default:
			fmt.Println("Invalid command. Valid commands are 'newanimal' and 'query'.")
		}
	}
}
