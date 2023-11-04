package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var IntCache = make([]int, 0, 10)

func BubbleSort(islice []int) []int {
	n := len(islice)
	swapped := false

	for {
		swapped = false

		for i := 0; i < n - 1; i++ {
			if islice[i] > islice[i+1] {
				swapped = true
				Swap(islice, i)
			}
		}
		n--
		if !swapped {
			break
		}
	}
	return islice
}

func ErrorCheck(err error) {
	if err != nil {
		errorMsg := fmt.Errorf("\nThere was an error with your entry: %s", err)
		fmt.Println(errorMsg)
		os.Exit(1)
	}
}

func ExitCheck(str string) {
	if len(str) == 3 && strings.ToLower(string(str[0])) == "x" {
		os.Exit(0)
	}
}

func IntCheck(str string) {
	ssclice := strings.Fields(str)
	print_bool := false
	for _, val := range ssclice {
		if len(IntCache) == 10 {
			fmt.Println("Cache is full!")
			break
		}
		int, err := strconv.Atoi(val)
		if err != nil {
			continue
		}
		IntCache = append(IntCache, int)
		if !print_bool {print_bool = !print_bool}
	}
	if print_bool {
		fmt.Println(IntCache)
	}
}

func SortCheck(str string) bool {
	if (len(str) < 4) {
		return false
	}
	if str[:4] == "sort" {
		sorted := BubbleSort(IntCache)
		fmt.Println(sorted)
		return true
	}
	return false
}

func Swap(slice []int, num int) {
	temp := slice[num]
	slice[num] = slice[num+1]
	slice[num+1] = temp
}

func main() {
	fmt.Println("Enter 'x' to exit")
	fmt.Println("Enter integer(s) to store for sorting")
	fmt.Println("Seperate multiple entries with a space")
	fmt.Println("Enter 'sort' to sort")
	for {
		fmt.Print(">> ")
    reader := bufio.NewReader(os.Stdin)
    entry, err := reader.ReadString('\n')
		ErrorCheck(err)
		ExitCheck(entry)
		if SortCheck(entry) {continue}
		IntCheck(entry)
	}
}