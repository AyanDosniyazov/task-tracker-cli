package task

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Status string

const (
	FileName                = "task.json"
	StatusDone       Status = "done"
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
)

type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

func loadTasks() (Tasks, error) {
	var tasks Tasks

	data, err := os.ReadFile(FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return Tasks{}, nil
		}
		return Tasks{}, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	if len(data) > 0 {
		if err := json.Unmarshal(data, &tasks); err != nil {
			return Tasks{}, fmt.Errorf("ошибка парсинга JSON: %w", err)
		}
	}
	return tasks, nil
}

func saveTasks(tasks Tasks) error {
	jsonTasks, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка кодирования JSON: %w", err)
	}
	if err := os.WriteFile(FileName, jsonTasks, 0644); err != nil {
		return fmt.Errorf("ошибка записи файла: %w", err)
	}
	return nil
}

func Get(status Status) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	var filtered []Task
	for _, t := range tasks.Tasks {
		if status == "" || t.Status == status {
			filtered = append(filtered, t)
		}
	}

	if len(filtered) == 0 {
		fmt.Println("Нет задач для отображения")
		return nil
	}

	fmt.Println("Tasks:")
	fmt.Println("---------------------------------------------------------")
	for _, t := range filtered {
		fmt.Printf("ID: %s | %s | Status: %s | Created: %s\n",
			t.ID, t.Description, t.Status, t.CreatedAt.Format("2006-01-02 15:04"))
	}
	fmt.Println("---------------------------------------------------------")

	return nil
}

func Add(description string) (string, error) {
	tasks, err := loadTasks()
	if err != nil {
		return "", err
	}

	newID := getNextID(tasks.Tasks)
	task := Task{
		ID:          newID,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks.Tasks = append(tasks.Tasks, task)

	if err = saveTasks(tasks); err != nil {
		return "", err
	}

	return newID, nil
}

func Update(id string, status Status, description string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	found := false
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			if status != "" {
				tasks.Tasks[i].Status = status
			}
			if description != "" {
				tasks.Tasks[i].Description = description
			}
			tasks.Tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("задача с ID %s не найдена", id)
	}

	return saveTasks(tasks)
}

// Удаляем задачу
func Delete(id string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks.Tasks {
		if t.ID == id {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("задача с ID %s не найдена", id)
	}

	return saveTasks(tasks)
}

func getNextID(tasks []Task) string {
	maxID := 0
	for _, t := range tasks {
		id, err := strconv.Atoi(t.ID)
		if err == nil && id > maxID {
			maxID = id
		}
	}
	return strconv.Itoa(maxID + 1)
}
