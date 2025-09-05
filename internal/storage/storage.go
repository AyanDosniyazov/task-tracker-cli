package storage

import (
	"encoding/json"
	"os"
	"task-tracker-cli/internal/task"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{fileName}
}

func (s *Storage[T]) Save(tasks T) error {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	if err = os.WriteFile(s.FileName, jsonData, 0644); err != nil {
		return err
	}

	return nil
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {

		if os.IsNotExist(err) {
			_, err = os.Create(s.FileName)
			if err != nil {
				return err
			}
		}
		return err
	}

	if err = json.Unmarshal(fileData, data); err != nil {
		return err
	}

	return nil
}

func (s *Storage[T]) GetLastID() (int, error) {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return 0, err
	}

	var tasks task.Tasks

	err = json.Unmarshal(fileData, &tasks)
	if err != nil {
		return 0, err
	}

	if len(tasks.Task) == 0 {
		return 0, nil
	}

	return len(tasks.Task) - 1, nil
}
