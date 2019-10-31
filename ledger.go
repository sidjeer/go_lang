//in the previous example we saw how to manage simple counter state using atomic operations. For more complex state we can use a mutex to safely access data across multiple goroutines.
package main

import (
	"fmt"
	
	"sync"
	"sync/atomic"
	"time"
)

func main() {
var states int = 0;
	var mutex = &sync.Mutex{}
	var readOps uint64
	var writeOps uint64
	for r := 0; r < 100; r++ {
		go func() {
		
			for {
				mutex.Lock()
				states = states +10;
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for w := 0; w < 10; w++ {
		go func() {
			for {
			mutex.Lock()
				states = states -5;
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	//Let the 10 goroutines work on the state and mutex for a second.
	time.Sleep(time.Second)
	//Take and report final operation counts.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
	//With a final lock of state, show how it ended up.
	mutex.Lock()
	fmt.Println("state:", states);
	mutex.Unlock()
}