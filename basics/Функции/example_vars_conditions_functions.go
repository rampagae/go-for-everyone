package main

import (
    "errors"
    "fmt"
)

// Сочетание: переменные + условия + функции

func safeDiv(a int, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func categoryByAge(age int) string {
    switch {
    case age < 0:
        return "invalid"
    case age < 12:
        return "child"
    case age < 18:
        return "teen"
    default:
        return "adult"
    }
}

func main() {
    // Переменные хранятся в памяти; читаем их и принимаем решения
    apples := 6
    eaten := 2
    apples = apples - eaten

    if apples > 0 {
        fmt.Println("Есть яблоки:", apples)
    } else {
        fmt.Println("Яблок нет")
    }

    // Используем функции, которые внутри применяют условия
    if q, err := safeDiv(10, 2); err != nil {
        fmt.Println("Ошибка деления:", err)
    } else {
        fmt.Println("Результат деления:", q)
    }

    fmt.Println("Категория по возрасту 15:", categoryByAge(15))

    // Под капотом:
    // - переменные размещены в памяти; операторы читают/записывают значения
    // - if/switch создают ветвления; выполняется подходящая ветка
    // - вызов функции создаёт стековый фрейм с аргументами и локальными переменными
    // - fmt выводит результат в stdout, ОС показывает его в терминале
}


