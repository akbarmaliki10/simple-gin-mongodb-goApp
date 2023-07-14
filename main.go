package main

import "fmt"

func main() {
	fmt.Println(saySomething())
}

func saySomething() (string, string) {
	return "something ", "else"
}
