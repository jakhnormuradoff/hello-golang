package main

import "fmt"

func main() { 
    fmt.Println("hello", hello())
}

func hello() string {
    return "Hello go"
}