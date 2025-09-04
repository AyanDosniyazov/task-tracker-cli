package main

import (
	"fmt"
	"log"
	"os"
	"task-tracker-cli/task"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	action := os.Args[1]
	args := os.Args[2:]

	if err := handleAction(action, args); err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func handleAction(action string, args []string) error {
	switch action {
	case "add":
		if len(args) < 1 {
			return fmt.Errorf("usage: go run main.go add <task-name>")
		}
		if _, err := task.Add(args[0]); err != nil {
			return err
		}
		log.Println("Task added successfully")

	case "update":
		if len(args) < 2 {
			return fmt.Errorf("usage: go run main.go update <task-id> <new-status>")
		}
		return task.Update(args[0], "", args[1])

	case "delete":
		if len(args) < 1 {
			return fmt.Errorf("usage: go run main.go delete <task-id>")
		}
		return task.Delete(args[0])

	case "mark-in-progress":
		if len(args) < 1 {
			return fmt.Errorf("usage: go run main.go mark-in-progress <task-id>")
		}
		return task.Update(args[0], "InProgress", "")

	case "mark-done":
		if len(args) < 1 {
			return fmt.Errorf("usage: go run main.go mark-done <task-id>")
		}
		return task.Update(args[0], "Done", "")

	case "list":
		status := ""
		if len(args) >= 1 {
			status = args[0]
		}
		return task.Get(task.Status(status))
	case "help":
		printUsage()
	default:
		return fmt.Errorf("unknown action: %s", action)
	}
	return nil
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add <task-name>            - Add a new task")
	fmt.Println("  update <task-id> <status>  - Update task status")
	fmt.Println("  delete <task-id>           - Delete a task")
	fmt.Println("  mark-in-progress <task-id> - Mark task as in progress")
	fmt.Println("  mark-done <task-id>        - Mark task as done")
	fmt.Println("  list [status]              - List tasks (done, todo, in-progress)")
}
