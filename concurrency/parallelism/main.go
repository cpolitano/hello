package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

var wg sync.WaitGroup

// init function runs before main.go
func init() {
	// will run concurrency with all CPUs
	// now included by default in Go, so adding it is not necessary
	// Go used to only run on 1 core by default unless this line was included
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go kale()
	go avocado()
	wg.Wait()
}

func kale() {
	for i := 0; i < 100; i++ {
		fmt.Println("kale:", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}

	wg.Done()
}

func avocado() {
	for j := 0; j < 100; j++ {
		fmt.Println("avocado:", j)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}

	wg.Done()
}

// race conditions
// go run --race main.go

