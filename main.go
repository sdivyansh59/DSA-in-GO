package main

import (
	"fmt"
	"sdivyansh59/dsa-in-go/stack"
)

func main() {
	fmt.Println("Subscribe to Go with Golang YT Channel.")
	myStack := stack.Stack{}

	myStack.Push(10)
	myStack.Push(20)
	myStack.Push(30)
	fmt.Println(myStack.IsEmpty())
	
	fmt.Println(myStack)

	myStack.Pop()
	myStack.Pop()
	myStack.Pop()


	fmt.Println(myStack)
	fmt.Println(myStack.IsEmpty())



}