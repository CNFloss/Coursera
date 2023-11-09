package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
A race condition in concurrent programming occurs when two or more threads can access shared data and they try to change it at the same time
*/

func main() {
	var count int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			count++
			runtime.Gosched()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			count--
			runtime.Gosched()
		}
	}()

	wg.Wait()

	fmt.Println("count:", count)
}
