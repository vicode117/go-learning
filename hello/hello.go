package main

import "fmt"

import "rsc.io/quote"

import "example.com/greetings"

func main() {
	fmt.Println("Hello, World!")
    fmt.Println(quote.Go())
    message := greetings.Hello("Victor")
    fmt.Println(message)
}
