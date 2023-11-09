package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func checkExit(str string) {
	if str == "x" {
		os.Exit(0)
	}
}

func mergeSortedSlices(a, b, c, d []int) []int {
	temp1 := mergeTwoSortedSlices(a, b)
	temp2 := mergeTwoSortedSlices(c, d)
	return mergeTwoSortedSlices(temp1, temp2)
}

func mergeTwoSortedSlices(a, b []int) []int {
	sorted := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			sorted = append(sorted, a[i])
			i++
		} else {
			sorted = append(sorted, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		sorted = append(sorted, a[i])
	}
	for ; j < len(b); j++ {
		sorted = append(sorted, b[j])
	}

	return sorted
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the x symbol to close")
	fmt.Println("Enter numbers to sort: ")

	for {
		fmt.Print("> ")
		scanned := scanner.Scan()
		if !scanned {
			fmt.Println("Error reading input")
			return
		}

		line := scanner.Text()
		checkExit(line)

		lineSlice := strings.Split(line, " ")
		ints := make([]int, 0, len(lineSlice))

		for _, str := range lineSlice {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Error converting '%s' to int: %s\n", str, err)
				return
			}
			ints = append(ints, num)
		}

		mid := len(ints) / 2
		firstHalf := ints[:mid]
		secondHalf := ints[mid:]

		firstQuarter := firstHalf[:len(firstHalf)/2]
		secondQuarter := firstHalf[len(firstHalf)/2:]
		thirdQuarter := secondHalf[:len(secondHalf)/2]
		fourthQuarter := secondHalf[len(secondHalf)/2:]

		var wg sync.WaitGroup
		sortSlice := func(slice []int) {
			defer wg.Done()
			slices.Sort(slice)
			fmt.Println("Sorted subarray:", slice)
		}

		wg.Add(4)
		go sortSlice(firstQuarter)
		go sortSlice(secondQuarter)
		go sortSlice(thirdQuarter)
		go sortSlice(fourthQuarter)
		wg.Wait()

		sortedNumbers := mergeSortedSlices(firstQuarter, secondQuarter, thirdQuarter, fourthQuarter)

		fmt.Println("Sorted array:", sortedNumbers)
	}
}
