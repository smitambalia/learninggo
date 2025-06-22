package main

import (
	"fmt"
	"sync"
)


type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) increaseViews(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock()
	p.views += 1
}

func main() {
	fmt.Println("Mutex in Go") 
	 // Mutexes (Mutual Exclusion Locks) are used to protect shared resources from concurrent access.

	myPost := post{views: 0}
	fmt.Println("Initial views:", myPost.views)
	wg := sync.WaitGroup{}

	for i:= range 10 {
		wg.Add(1) // Increment the WaitGroup counter
		i += 0 // Simulate some processing
		go myPost.increaseViews(&wg)
	}
	wg.Wait()
	fmt.Println("Updated views:", myPost.views) // Print the updated views
}