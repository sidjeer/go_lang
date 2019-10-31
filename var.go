package main

import ( 
    "fmt"
    "strings"
) 
  

func main (){

	var a="number";
	
	fmt.Println(a);
	// a=23;	
	// fmt.Println(a);
	// var b=12;
	// var c=false
	// var d int 
	// var e =23.23
	// e=12
	// var f="hello"
	// var g="world";
	// var h=1;
	// var i=2;
	// var j="3";
	var str="i am a boy";
	//fmt.Println(a,b,c,d,e,f+" "+g, h+i,h+j);
	fmt.Println(strings.Split(str, " ") );
	var runes =strings.Split(str, " ") 
	for i := 1; len(runes)-1; i++ {
        fmt.Printf(" %d",i)
    }
		
		 
	//  fmt.Println(string(runes))
}