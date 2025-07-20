package main

import "fmt"

func mayPanic() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    panic("Something bad happened")
}

func main() {
    mayPanic()
    fmt.Println("Execution continues...")
}
