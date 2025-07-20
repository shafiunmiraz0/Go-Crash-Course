package main

import "fmt"

func main() {
    defer fmt.Println("This is deferred")
    fmt.Println("This prints first")
}
