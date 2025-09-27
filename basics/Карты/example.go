package main

import "fmt"

func main() {
    ages := make(map[string]int)
    ages["Masha"] = 20
    ages["Ivan"] = 25

    fmt.Println("Masha:", ages["Masha"]) // 20

    if v, ok := ages["Oleg"]; ok {
        fmt.Println("Oleg:", v)
    } else {
        fmt.Println("Oleg не найден")
    }

    delete(ages, "Ivan")
    for name, age := range ages {
        fmt.Println(name, age)
    }
}


