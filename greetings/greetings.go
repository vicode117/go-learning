package greetings

// change this importable package to main program
// package main

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// change this importable package to main program
// func main() {
// 	fmt.Println("Hello, World!")
// 	message := Hello("Victor")
// 	fmt.Println(message)
// }
