//Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time. Implementing timeouts in Go is easy and elegant thanks to channels and select.
package main

import "time"
import "fmt"

func main() {
	//For our example, suppose we’re executing an external call that returns its result on a channel image after 2s.
	image := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		image <- "image found"
	}()
	//Here’s the select implementing a timeout. res := <-image awaits the result and <-Time.After awaits a value to be sent after the timeout of 1s. Since select proceeds with the first receive that’s ready, we’ll take the timeout case if the operation takes more than the allowed 1s.
	select {
	case res := <-image:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("image  not found")
	}
	//If we allow a longer timeout of 3s, then the receive from text will succeed and we’ll print the result.
	text := make(chan string, 1)
	go func() {
		time.Sleep(7 * time.Second)
		text <- "text found"
	}()
	select {
	case res := <-text:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("text not found")
	}
}