package main

import "fmt"
func h(from string) {

	messagess <- "pings"
	fmt.Println("teaser ",from)
}
func main() {
	
	messagess := make(chan string)
	go h("action")
	//Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	 messages := make(chan string)
	// //Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.
	 go func() { messages <- "ping" }()
	// //he <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
	 msg := <-messages
	 m := <-messagess 
	 fmt.Println( m)
	 fmt.Scanln()
	 fmt.Println(msg)
	 
}