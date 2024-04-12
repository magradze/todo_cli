package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/magradze/todo_cli"
)

// toDoFilename is the name of the file used to store the todo list.
const (
	toDoFilename = ".todo.json"
)

func main() {
	add := flag.String("add", "", "Add a task to the list")
	complete := flag.Int("complete", 0, "Complete a task on the list")
	del := flag.Int("del", 0, "Delete a task from the list")
	list := flag.Bool("list", false, "List all tasks")
	flag.Parse()

	todos := &todo_cli.List{}

	if err := todos.Load(toDoFilename); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err.Error())
		if err != nil {
			return
		}
		os.Exit(1)
	}

	switch {
	case *add != "":
		todos.AddTask(*add)
		if err := todos.Save(toDoFilename); err != nil {
			_, err := fmt.Fprintln(os.Stderr, err.Error())
			if err != nil {
				return
			}
			os.Exit(1)
		}
	case *complete > 0:
		if err := todos.CompleteTask(*complete); err != nil {
			_, err := fmt.Fprintln(os.Stderr, err.Error())
			if err != nil {
				return
			}
			os.Exit(1)
		}
		if err := todos.Save(toDoFilename); err != nil {
			_, err := fmt.Fprintln(os.Stderr, err.Error())
			if err != nil {
				return
			}
			os.Exit(1)
		}
	case *del > 0:
		if err := todos.DeleteTask(*del); err != nil {
			_, err := fmt.Fprintln(os.Stderr, err.Error())
			if err != nil {
				return
			}
			os.Exit(1)
		}
		if err := todos.Save(toDoFilename); err != nil {
			_, err := fmt.Fprintln(os.Stderr, err.Error())
			if err != nil {
				return
			}
			os.Exit(1)
		}
	case *list:
		if len(*todos) == 0 {
			fmt.Println("You have no tasks to complete!")
		}
		todos.PrintTable()
	default:
		_, err := fmt.Fprintln(os.Stdout, "Invalid command!")
		if err != nil {
			return
		}
		os.Exit(0)
	}
}
