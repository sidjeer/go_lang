package main

import (
	"io"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond*6);
	io.WriteString(w, "<!DOCTYPE html><html><body><h2>JavaScript Alert</h2><button onclick='myFunction()'>Try it</button><script>function myFunction(){alert('I am an alert box!');}</script></body></html>")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8000", nil)
}