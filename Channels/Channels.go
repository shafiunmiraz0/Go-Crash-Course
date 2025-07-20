package main

import "fmt"

func sendData(ch chan string) {
    ch <- "Hello from channel"
}

func main() {
    ch := make(chan string)
    go sendData(ch)
    msg := <-ch
    fmt.Println(msg)
}
