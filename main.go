package main

import (
	"fmt"
	"sync"
	"time"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}


func main() {
	// main function itself is a goroutine are a verry light weight thread not on hardware level thread are specific to go itself
	// they take verry little memory run verry quickly and are all managed as a group of go routines is called co-routines 
	// and are managed by go-scheduler
	var wg sync.WaitGroup

	words := []string{
		"alpha", "beta", "delta", "gamma", "pi", "zeta", "eta", "theta", "epsilon",
	}

	wg.Add(9) // what happen if we add more than the number of goroutines call 
	// wg.Add(12) it will be in deadlock situation
	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}
	wg.Wait()

	wg.Add(2)
	go printSomething("This is the first thing to be printed!", &wg)
	time.Sleep(1 * time.Second) // shouldn't wait using time.sleep wait group is here to rescue
	printSomething("This is the Second thing to be printed!", &wg)
}