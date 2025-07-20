package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }

    names := []string{"Go", "Rust", "Python"}
    for index, name := range names {
        fmt.Println(index, name)
    }
}
