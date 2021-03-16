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
	pool := Pool{
		work: make(chan Worker),
	}

	// Make maxGoroutines
	pool.wg.Add(maxGoroutines)

	// Register Goroutines with the waitgroup

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for t := range pool.work {
				t.Task()
			}

			pool.wg.Done()
		}()
	}

	// When each worker finishes because the worker channel is closed then update the waitgroup
	return &pool
}

// Enqueue another task to be run
// This call will block until a worker can pick up the Task
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown all the worker queue
// Block until all Goroutines finish
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
