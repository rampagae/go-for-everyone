package main

import "fmt"

func main() {
    s := "Привет"
    fmt.Println("bytes len:", len(s))

    // Проходим по рунам (символам)
    for i, r := range s {
        fmt.Println(i, string(r))
    }
}


