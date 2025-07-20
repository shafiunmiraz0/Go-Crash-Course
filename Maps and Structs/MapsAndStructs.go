package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    person := Person{Name: "Miraz", Age: 25}
    m := map[string]int{"apple": 5, "banana": 3}
    
    fmt.Println(person, m)
}
