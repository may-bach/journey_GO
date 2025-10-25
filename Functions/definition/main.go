package main

import "fmt"


func sayHello() {
    fmt.Println("Welcome.\n")
}

func greet(name string) {
    fmt.Printf("Nice to meet you, %s!\n", name)
}

func add(num1 int, num2 int) {
    result := num1 + num2
    fmt.Printf("%d + %d = %d\n", num1, num2, result)
}

func main() {
    
    sayHello() 
    
    greet("Izumi")

    add(10, 5)
    add(99, 1)
}