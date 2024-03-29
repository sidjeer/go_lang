 package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
	go h(from );
}
func h(from string) {
	fmt.Println("teaser ",from)
}
func main() {
	go f("test")
	//Suppose we have a function call f(s). Here’s how we’d call that in the usual way, running it synchronously.
	f("direct")
	go func(msg string) {
		fmt.Println(msg)
	}("gone")
	//To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.
	go f("goroutine")
	//You can also start a goroutine for an anonymous function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	
	//Our two function calls are running asynchronously in separate goroutines now, so execution falls through to here. This Scanln requires we press a key before the program exits.
	fmt.Scanln()
	fmt.Println("done")
}