package main

import "fmt"

func main() {
    fmt.Println("Before panic")
    panic("Something went wrong")
    fmt.Println("After panic") // This won't run
}
