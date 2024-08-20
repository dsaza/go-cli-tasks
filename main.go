package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dsaza/go-cli-tasks/tasks"
)

func main() {
	tasks.Init()

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "list":
		tasks.List()
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter task name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		tasks.Add(name)
		tasks.List()
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go-cli-crud delete [task_id]")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		taskDeleted := tasks.Delete(id)
		if taskDeleted {
			tasks.List()
		}
	case "check":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go-cli-crud check [task_id]")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		taskEdited := tasks.ToggleCheck(id, true)
		if taskEdited {
			tasks.List()
		}
	case "uncheck":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go-cli-crud uncheck [task_id]")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		taskEdited := tasks.ToggleCheck(id, false)
		if taskEdited {
			tasks.List()
		}
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage: go-cli-crud [list|add|delete|check|uncheck]")
}
