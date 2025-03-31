package main

import (
	"fmt"
	"sync"
	"time"
)

// Philosopher is a struct which stores information about a philosopher
type Philospher struct {
	name string
	rightFork int
	leftFork int
}

// philosophers is list of all philosophers.
var philosophers = []Philospher {
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// define some variables
var hunger = 3 // how many times does a person eat ?
var eatTime = 1*time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

// *** added this
var orderMutex sync.Mutex // a mutex for the slice orderFinished; part of challenge
var orderFinished []string // the order in which philosophers finish dining and leave; part of challenge!

func main() {
	// print out a welcome message
	fmt.Println("Dining Philosopher Problem")
	fmt.Println("--------------------------")
	fmt.Println("The table is empty .")

	// *** added this
	time.Sleep(sleepTime)
	
	// start the meal
	dine()

	// print out finished message
	fmt.Println("The table is empty.")


}

func dine() {
	// eatTime = 0 * time.Second
	// sleepTime = 0 * time.Second
	// thinkTime = 0 * time.Second
	
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}


	// start the meal.
	for i := 0; i < len(philosophers); i++ {
		// fire off a goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}
	wg.Wait()

}

func diningProblem(philospher Philospher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	//seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n", philospher.name)
	seated.Done()

	seated.Wait()

	// eat three times
	for i := hunger; i > 0; i-- {
		// get a lock on both forks
		if philospher.leftFork > philospher.rightFork {

			forks[philospher.rightFork].Lock()
			fmt.Printf("\t %s takes the right fork. \n", philospher.name)

			forks[philospher.leftFork].Lock()
			fmt.Printf("\t %s takes the left fork. \n", philospher.name)
		} else {

			forks[philospher.leftFork].Lock()
			fmt.Printf("\t %s takes the left fork. \n", philospher.name)

			forks[philospher.rightFork].Lock()
			fmt.Printf("\t %s takes the right fork. \n", philospher.name)
		}




		// forks[philospher.leftFork].Lock()
		// fmt.Printf("\t %s takes the left fork. \n", philospher.name)
		// forks[philospher.rightFork].Lock()
		// fmt.Printf("\t %s takes the right fork. \n", philospher.name)

		fmt.Printf("\t %s has both forks and is eating.\n", philospher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t %s is Thinking.\n", philospher.name)
		time.Sleep(thinkTime)

		forks[philospher.leftFork].Unlock()
		forks[philospher.rightFork].Unlock()

		fmt.Printf("\t %s put down the forks.\n", philospher.name)

	}

	fmt.Println(philospher.name, "is satisfied")
	fmt.Println(philospher.name, "left the table")


	// *** added this
	orderMutex.Lock()
	orderFinished = append(orderFinished, philospher.name)
	orderMutex.Unlock()

}





