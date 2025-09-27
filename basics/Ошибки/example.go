package main

import (
    "errors"
    "fmt"
)

func div(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    if q, err := div(10, 0); err != nil {
        fmt.Println("Ошибка:", err)
    } else {
        fmt.Println("Результат:", q)
    }
}


