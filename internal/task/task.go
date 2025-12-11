package task

import (
	"errors"
	"fmt"
)

// Task struct
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var tasks []Task
var nextID int = 1

// AddTask nambahan pagawean anyar
func AddTask(description string) (Task, error) {
	newTask := Task{
		ID:          nextID,
		Description: description,
		Status:      "todo",
	}
	tasks = append(tasks, newTask)
	nextID++
	return newTask, nil
}

// UpdateTask nganyarkeun deui pagawean
func UpdateTask(id int, newDescription string) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = newDescription
			return nil
		}
	}
	return errors.New("task eweuh teu kapanggih")
}

// DeleteTask ngahapus pagawean
func DeleteTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task eweuh teu kapanggih")
}

// MarkInProgress nandaan pagawean keur jalan
func MarkInProgress(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = "in-progress"
			return nil
		}
	}
	return errors.New("task eweuh teu kapanggih")
}

// MarkDone nandaan pagawean geus beres
func MarkDone(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = "done"
			return nil
		}
	}
	return errors.New("task eweuh teu kapanggih")
}

// ListTasks nampilkeun daptar pagawean
func ListTasks(statusFilter string) {
	if len(tasks) == 0 {
		fmt.Println("Teu aya pagawean nu kapanggih.")
		return
	}
	for _, task := range tasks {
		if statusFilter == "" || task.Status == statusFilter {
			fmt.Printf("ID: %d, Description: %s, Status: %s\n", task.ID, task.Description, task.Status)
		}
	}
}
