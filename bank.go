//Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.
package main

import "fmt"

type Configuration struct {
    Val string
    Proxy
}

type Proxy struct {
    Address string
    Port    string
}

func main() {

    c := &Configuration{
        Val: "test",
        Proxy: Proxy{
            Address: "addr",
            Port:    "port",
        },
    }
    fmt.Println(c)
}