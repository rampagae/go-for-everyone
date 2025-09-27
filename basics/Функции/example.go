package main

import (
    "errors"
    "fmt"
)

// Примеры функций: параметры, возвраты, обработка ошибок.

// sum возвращает сумму двух целых чисел
func sum(a int, b int) int {
    return a + b
}

// div делит a на b и возвращает результат или ошибку, если b == 0
func div(a int, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// isAdult возвращает true, если возраст >= порога
func isAdult(age int, threshold int) bool {
    return age >= threshold
}

func main() {
    fmt.Println("sum:", sum(2, 3))

    q, err := div(10, 0)
    if err != nil {
        fmt.Println("error:", err)
    } else {
        fmt.Println("div:", q)
    }

    fmt.Println("isAdult 20>=18:", isAdult(20, 18))
}


