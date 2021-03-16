package main

import (
	"go-work/work"
	"log"
	"math/rand"
	"sync"
	"time"
)


var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// Create a type to adhere to the Worker interface
type namePrinter struct {
	name string
}

func (n *namePrinter) Task() {
	log.Println(n.name)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func main() {
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(20*len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{name}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	p.Shutdown()
}
