package main

import "fmt"

func main() {
    // Массив фиксированного размера
    arr := [3]int{1, 2, 3}
    fmt.Println("array:", arr)

    // Срез — гибкий, растёт append'ом
    s := []int{1, 2}
    fmt.Println("slice:", s, "len:", len(s), "cap:", cap(s))

    s = append(s, 3)
    fmt.Println("after append:", s, "len:", len(s), "cap:", cap(s))

    // Создадим с запасом по ёмкости
    t := make([]int, 0, 5)
    t = append(t, 10, 20, 30)
    fmt.Println("t:", t, "len:", len(t), "cap:", cap(t))
}


