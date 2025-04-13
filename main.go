package main

import (
	"fmt"
	"slices"
)

func main(){
	var todos []string = []string{"Todo 1"}

	var command string;
	fmt.Println("Enter command")
	fmt.Scanln(&command)

	for {
		if command == "add" {
			var todoName string
			fmt.Println("Enter todo!")
			fmt.Scanln(&todoName)
		
			todos = append(todos, todoName)
	
			fmt.Println("Todo added!")
		} else if command == "edit" {
			var todoName string
			fmt.Println("Enter todo to edit: ")
			fmt.Scanln(&todoName)

			index := slices.Index(todos, todoName)

			if index == -1 {
				fmt.Println("Element not found")
			} else {
				fmt.Println("Todo was found")
				todos[index] = todoName
				fmt.Println("Todo was edited!")
			}
		} else if command == "delete" {
			var todoName string
			fmt.Println("Enter todo to delete:")
			fmt.Scanln(&todoName)

			index := slices.Index(todos, todoName)

			if index == -1 {
				fmt.Println("Element not found")
			} else {
				fmt.Println("Todo was found")
				slices.Delete(todoName, todoName)
			}
			
		}
	}

}