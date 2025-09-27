package main

import "fmt"

// Демонстрация базовых условий: if/else и switch.

func main() {
    n := 7

    // if / else: выполняется ровно одна ветка
    if n%2 == 0 {
        fmt.Println("even")
    } else {
        fmt.Println("odd")
    }

    // if с короткой инициализацией: x доступен только внутри веток
    if x := n * 2; x > 10 {
        fmt.Println("x > 10")
    } else {
        fmt.Println("x <= 10")
    }

    // switch по значению
    day := 6
    switch day {
    case 6, 7:
        fmt.Println("weekend")
    default:
        fmt.Println("workday")
    }

    // switch без выражения — case с условиями
    score := 82
    switch {
    case score >= 90:
        fmt.Println("A")
    case score >= 80:
        fmt.Println("B")
    default:
        fmt.Println("C or below")
    }
}


