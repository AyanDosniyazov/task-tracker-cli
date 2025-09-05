package task

import (
	"errors"
	tb "github.com/aquasecurity/table"
	"github.com/liamg/tml"
	"os"
	"strconv"
	"time"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

var StatusBeautify = map[Status]string{
	StatusTodo:       "<red>ToDo</red>",
	StatusInProgress: "<yellow>In Progress</yellow>",
	StatusDone:       "<green>Done</green>",
}

type Task struct {
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tasks struct {
	Task []Task `json:"tasks"`
}

func (t *Tasks) List(status Status) {
	t.print(status)
}
func (t *Tasks) Add(description string) {
	task := Task{
		Description: description,
		Status:      string(StatusTodo),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	t.Task = append(t.Task, task)
}

func (t *Tasks) Update(id int, description string, status Status) error {
	if status == "" || !status.isValid() {
		return errors.New("invalid status")
	}
	for i, _ := range t.Task {
		if i+1 == id {
			if description != "" {
				t.Task[i].Description = description
			}
			if status != "" {
				t.Task[i].Status = string(status)
			}
			t.Task[i].UpdatedAt = time.Now()
		}
	}

	return nil
}

func (t *Tasks) Delete(id int) {
	for i, _ := range t.Task {
		if i+1 == id {
			t.Task = append(t.Task[:i], t.Task[i+1:]...)
		}
	}
}

func (t *Tasks) DeleteAll() {
	t.Task = nil
}

func (t *Tasks) print(status Status) {
	table := tb.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("ID", "DESCRIPTION", "STATUS", "CREATED AT", "UPDATED AT")
	for index, td := range t.Task {
		currentStatus := Status(td.Status)

		if status != "" && currentStatus != status {
			continue
		}

		if beautified, ok := StatusBeautify[currentStatus]; ok {
			td.Status = tml.Sprintf(beautified)
		}

		table.AddRow(
			strconv.Itoa(index+1),
			td.Description,
			td.Status,
			td.CreatedAt.Format(time.RFC822),
			td.UpdatedAt.Format(time.RFC822),
		)
	}

	table.Render()
}

func (s Status) isValid() bool {
	switch s {
	case StatusTodo, StatusInProgress, StatusDone:
		return true
	}
	return false
}
