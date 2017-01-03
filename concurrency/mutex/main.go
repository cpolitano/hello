package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64 // indicates atomic
// atomic creates values that can only be altered by one process at a time

// MUTEX
// mutually exclusive
var mutex sync.Mutex

// init function runs before main.go
func init() {
	// will run concurrency with all CPUs
	// now included by default in Go, so adding it is not necessary
	// Go used to only run on 1 core by default unless this line was included
	// this enables parallelism
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go incrementor("A")
	go incrementor("B")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)

		// mutex locks these variables from being changed by other processes
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()

		// atomicity
		// atomic.AddInt64(&counter, 1)
		// fmt.Println(s, i, "Counter:", counter)
	}

	wg.Done()
}
