	package main

	import "fmt"
    import s "strings"
    var p = fmt.Println
	func main (){
	p("hello World ",s.Contains("hello","lo"));
	p("hello World ",s.Count("hello","lo"));
	p("hello World ",s.HasPrefix("hello","he"));
	p("hello World ",s.HasSuffix("hello","lo"));
	}