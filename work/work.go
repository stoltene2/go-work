package work

import (
	"sync"
)

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	// Initialize the Pool

	// Make maxGoroutines
	// Register Goroutines with the waitgroup
	// When each worker finishes because the worker channel is closed then update the waitgroup
	return nil
}

// Enqueue another task to be run
// This call will block until a worker can pick up the Task
func (p *Pool) Run(w Worker) {

}

// Shutdown all the worker queue
// Block until all Goroutines finish
func (p *Pool) Shutdown() {

}
