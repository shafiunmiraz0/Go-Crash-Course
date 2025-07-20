package main

import "fmt"

func main() {
    age := 20
    if age >= 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }

    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("Start of the week")
    case "Friday":
        fmt.Println("Weekend is coming")
    default:
        fmt.Println("Midweek")
    }
}
