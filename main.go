package main

import (
	"solotodo/scanner"
	"solotodo/todo"
)

func main() {
	todoList := todo.NewList()

	scanner := scanner.NewScanner(todoList)

	scanner.Start()
}
