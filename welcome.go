package main

import "fmt"

func welcome(name string) {
	fmt.Println("Welcome, " + name + "!")
}

func main() {
	welcome("Sergey")
	welcome("Alex")
}
