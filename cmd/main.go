package main

import (
	"log"
	"os"
	"task-tracker-cli/internal/command"
	"task-tracker-cli/internal/storage"
	"task-tracker-cli/internal/task"
)

func main() {
	tasks := task.Tasks{}
	st := storage.NewStorage[task.Tasks]("tasks.json")
	err := st.Load(&tasks)
	if err != nil {
		log.Fatal(err)
	}

	cmd := command.NewCommand()
	err = cmd.Execute(os.Args, tasks, *st)
	if err != nil {
		log.Fatal(err)
	}
}
