package utils

import (
	"encoding/json"
	"os"
	"project-app-todo-list-cli-azwin/model"
)

func LoadTasksFromFile(fileName string) ([]model.Task, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		// Jika file tidak ada, return empty slice
		if os.IsNotExist(err) {
			return []model.Task{}, nil
		}
		return nil, err
	}

	var tasks []model.Task
	if len(data) == 0 {
		return []model.Task{}, nil
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasksToFile(fileName string, tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}