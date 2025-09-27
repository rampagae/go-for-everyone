package main

import "fmt"

// Классический цикл счётчика и альтернативные формы for

func main() {
    // Классический счётчик: от 1 до 10 включительно
    for i := 1; i <= 10; i++ {
        fmt.Println("i:", i)
    }

    // Форма как while: повторяем, пока условие истинно
    j := 0
    for j < 3 {
        fmt.Println("j:", j)
        j++
    }

    // Бесконечный цикл с break
    k := 0
    for {
        if k >= 2 {
            break
        }
        fmt.Println("k:", k)
        k++
    }
}


