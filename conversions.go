package main

import "fmt"

type person struct {
    name string
    age  int
}

func newPerson(name string) *person {

    p := person{name: name}
    p.age = 42
    return &p
}

func main() {
    // 1
    fmt.Printf("Hello, World: %d\n", 10 )
    // 2
    var byte []byte = []byte("this is a byte slice")
    //fmt.Printf("This is a %T", byte)
    fmt.Printf(string(byte))
    fmt.Printf("\n")

    fmt.Printf("%s\n", byte)
    fmt.Printf("%q\n", byte)
}