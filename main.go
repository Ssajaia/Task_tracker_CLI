package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Description string   `json:"description"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

const fileName = "tasks.json"

func loadTasks() ([]Task, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []Task{}, nil
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, data, 0644)
}

func addTask(desc string) {
	tasks, _ := loadTasks()
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}
	task := Task{
		ID: id,
		Description: desc,
		Status: "todo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tasks = append(tasks, task)
	saveTasks(tasks)
	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
}

func updateTask(id int, desc string) {
	tasks, _ := loadTasks()
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = desc
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Println("Task updated successfully")
			return
		}
	}
	fmt.Println("Task not found")
}

func deleteTask(id int) {
	tasks, _ := loadTasks()
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks(tasks)
			fmt.Println("Task deleted successfully")
			return
		}
	}
	fmt.Println("Task not found")
}

func markStatus(id int, status string) {
	tasks, _ := loadTasks()
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Println("Task status updated successfully")
			return
		}
	}
	fmt.Println("Task not found")
}

func listTasks(filter string) {
	tasks, _ := loadTasks()
	for _, t := range tasks {
		if filter == "" || t.Status == filter {
			fmt.Printf("ID: %d | Description: %s | Status: %s | CreatedAt: %s | UpdatedAt: %s\n",
				t.ID, t.Description, t.Status, t.CreatedAt.Format(time.RFC3339), t.UpdatedAt.Format(time.RFC3339))
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add \"description\"")
			return
		}
		addTask(strings.Join(os.Args[2:], " "))
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <id> \"new description\"")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		updateTask(id, strings.Join(os.Args[3:], " "))
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		deleteTask(id)
	case "mark-in-progress":
		id, _ := strconv.Atoi(os.Args[2])
		markStatus(id, "in-progress")
	case "mark-done":
		id, _ := strconv.Atoi(os.Args[2])
		markStatus(id, "done")
	case "list":
		filter := ""
		if len(os.Args) > 2 {
			filter = os.Args[2]
		}
		listTasks(filter)
	default:
		fmt.Println("Unknown command")
	}
}

