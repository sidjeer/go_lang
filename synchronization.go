//We can use channels to synchronize execution across goroutines. Here’s an example of using a blocking receive to wait for a goroutine to finish.
package main

import "fmt"
import "time"

//This is the function we’ll run in a goroutine. The done channel will be used to notify another goroutine that this function’s work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	//end a value to notify that we’re done.
	done <- true
}
func workers(done1 chan bool) {
	fmt.Print("working...");
	time.Sleep(time.Second);
	fmt.Println("done");
	//end a value to notify that we’re done.
	done1 <- true
}
func main() {
	//Start a worker goroutine, giving it the channel to notify on.
//	done := make(chan bool, 1)
	done1 := make(chan bool, 1)
//	go worker(done)
	go workers(done1)
	//Block until we receive a notification from the worker on the channel.
	//<-done
	<-done1
}