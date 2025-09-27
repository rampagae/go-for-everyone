package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) Greeting() string {
    return "Привет, я " + p.Name
}

func main() {
    masha := Person{Name: "Маша", Age: 20}
    fmt.Println(masha.Name, masha.Age)
    fmt.Println(masha.Greeting())
}


