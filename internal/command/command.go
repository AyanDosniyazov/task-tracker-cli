package command

import (
	"errors"
	tb "github.com/aquasecurity/table"
	"github.com/liamg/tml"
	"os"
	"strconv"
	"task-tracker-cli/internal/storage"
	"task-tracker-cli/internal/task"
)

type Command struct {
	Command     string
	Description string
}

type Commands struct {
	Commands []Command
}

func NewCommand() *Commands {
	return &Commands{}
}

func (c *Commands) Execute(args []string, tasks task.Tasks, storage storage.Storage[task.Tasks]) error {
	c.defaultCommand()

	switch args[1] {
	case "add":
		var description string
		if len(args) > 2 {
			description = args[2]
		} else {
			return errors.New("invalid argument")
		}
		tasks.Add(description)
	case "update":
		var description string
		var id int
		if len(args) > 3 {
			id, _ = strconv.Atoi(args[2])
			description = args[3]
		} else {
			return errors.New("invalid argument")
		}

		err := tasks.Update(id, description, "")
		if err != nil {
			return err
		}
	case "mark-in-progress":
		var id int
		if len(args) > 2 {
			id, _ = strconv.Atoi(args[2])
		} else {
			return errors.New("invalid argument")
		}

		err := tasks.Update(id, "", task.StatusInProgress)
		if err != nil {
			return err
		}
	case "mark-done":
		var id int
		if len(args) > 2 {
			id, _ = strconv.Atoi(args[2])
		} else {
			return errors.New("invalid argument")
		}

		err := tasks.Update(id, "", task.StatusDone)
		if err != nil {
			return err
		}
	case "delete":
		var id int
		if len(args) > 3 {
			id, _ = strconv.Atoi(args[2])
		} else {
			return errors.New("invalid argument")
		}

		tasks.Delete(id)
	case "list":
		var status task.Status
		if len(args) > 2 {
			status = task.Status(args[2])
		}
		tasks.List(status)
	default:
		c.printUsage()
	}
	err := storage.Save(tasks)
	if err != nil {
		return err
	}

	return nil
}

func (c *Commands) defaultCommand() {
	c.Commands = append(c.Commands,
		Command{
			Command:     tml.Sprintf("./task list <green><status></green>"),
			Description: "List tasks (done, task, in-progress)",
		},
		Command{
			Command:     tml.Sprintf("./task add <green><description><green> <green><status></green>"),
			Description: "Add a new task",
		},
		Command{
			Command:     tml.Sprintf("./task update <yellow><id></yellow> <green><status></green>"),
			Description: "Update task status",
		},
		Command{
			Command:     tml.Sprintf("./task delete <yellow><id></yellow>"),
			Description: "Delete a task",
		},
	)
}

func (c *Commands) printUsage() {
	table := tb.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("COMMANDS", "DESCRIPTION")

	for _, command := range c.Commands {
		table.AddRow(command.Command, command.Description)
	}

	table.Render()
}
