package main

import "fmt"

// Указатели: адрес (&) и разыменование (*)

func increment(p *int) {
    // p указывает на ячейку с числом; меняем исходное значение
    *p = *p + 1
}

func main() {
    count := 3
    ptr := &count // адрес переменной count

    fmt.Println("До:", count)
    increment(ptr) // передаём «адрес коробочки»
    fmt.Println("После:", count)

    // Считываем через указатель
    fmt.Println("Через *ptr:", *ptr)
}


