package greetings

// change this importable package to main program
// package main

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("Oops, empty name!")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// Add nil (meaning no error) as a second value in the successful return
	return message, nil
}

// the long way to Declare a message variable to hold your greeting
func HelloTest(name string) string {
	var message string
	message = fmt.Sprintf("Hi, %v. Welcome", name)
	return message
}

func Pause() {
	fmt.Println("---------------------------------------------------")
	var s string
	fmt.Println("输入exit退出:")
	fmt.Scan(&s)
	if s == "exit" {
	} else {
		Pause()
	}

	// fmt.Printf("输入任意键退出:")
	// b := make([]byte, 1)
	// os.Stdin.Read(b)
}

// change this importable package to main program
// func main() {
// 	fmt.Println("Hello, World!")
// 	message := Hello("Victor")
// 	fmt.Println(message)
// }
