package service

import (
	"errors"
	"fmt"
	"project-app-todo-list-cli-azwin/model"
	"project-app-todo-list-cli-azwin/utils"
	"time"
)

type TodoService struct {
	dataFilePath string
	tasks        []model.Task
}

func NewTodoService(filePath string) *TodoService {
	return &TodoService{
		dataFilePath: filePath,
		tasks:        []model.Task{},
	}
}

// LoadTasks - memuat tasks dari file
func (s *TodoService) LoadTasks() error {
	tasks, err := utils.LoadTasksFromFile(s.dataFilePath)
	if err != nil {
		return err
	}
	s.tasks = tasks
	return nil
}

// SaveTasks - menyimpan tasks ke file
func (s *TodoService) SaveTasks() error {
	return utils.SaveTasksToFile(s.dataFilePath, s.tasks)
}

// AddTask - menambah task baru
func (s *TodoService) AddTask(taskName string) error {
	if err := s.LoadTasks(); err != nil {
		return err
	}

	// Generate ID baru
	newID := 1
	if len(s.tasks) > 0 {
		newID = s.tasks[len(s.tasks)-1].ID + 1
	}

	newTask := model.Task{
		ID:        newID,
		Task:      taskName,
		Status:    false,
		CreatedAt: time.Now(),
	}

	s.tasks = append(s.tasks, newTask)
	return s.SaveTasks()
}

// DisplayTasks - menampilkan semua tasks
func (s *TodoService) DisplayTasks() error {
	if err := s.LoadTasks(); err != nil {
		return err
	}

	if len(s.tasks) == 0 {
		fmt.Println("Tidak ada task.")
		return nil
	}

	fmt.Println("\n=== Daftar Task ===")
	for _, task := range s.tasks {
		status := "[ ]"
		if task.Status {
			status = "[✓]"
		}
		fmt.Printf("%d. %s %s (Created: %s)\n",
			task.ID,
			status,
			task.Task,
			task.CreatedAt.Format("2006-01-02 15:04:05"),
		)
	}
	fmt.Println()
	return nil
}

// UpdateTask - mengupdate nama task
func (s *TodoService) UpdateTask(id int, newTaskName string) error {
	if err := s.LoadTasks(); err != nil {
		return err
	}

	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks[i].Task = newTaskName
			return s.SaveTasks()
		}
	}

	return errors.New("task tidak ditemukan")
}

// DeleteTask - menghapus task
func (s *TodoService) DeleteTask(id int) error {
	if err := s.LoadTasks(); err != nil {
		return err
	}

	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return s.SaveTasks()
		}
	}

	return errors.New("task tidak ditemukan")
}

// MarkComplete - menandai task sebagai selesai
func (s *TodoService) MarkComplete(id int) error {
	if err := s.LoadTasks(); err != nil {
		return err
	}

	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks[i].Status = true
			return s.SaveTasks()
		}
	}

	return errors.New("task tidak ditemukan")
}

// MarkIncomplete - menandai task sebagai belum selesai
func (s *TodoService) MarkIncomplete(id int) error {
	if err := s.LoadTasks(); err != nil {
		return err
	}

	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks[i].Status = false
			return s.SaveTasks()
		}
	}

	return errors.New("task tidak ditemukan")
}