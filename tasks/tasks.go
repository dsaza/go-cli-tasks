package tasks

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var tasks []Task
var file *os.File

func Init() {
	fileRead, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	file = fileRead

	fileInfo, err := fileRead.Stat()
	if err != nil {
		panic(err)
	}

	if fileInfo.Size() == 0 {
		tasks = []Task{}
	} else {
		bytes, err := io.ReadAll(fileRead)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	}
}

func List() {
	if len(tasks) == 0 {
		fmt.Println("No tasks")
		return
	}

	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "✓"
		}

		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Name)
	}
}

func Add(name string) {
	task := Task{
		ID:        int(time.Now().UnixMilli()),
		Name:      name,
		Completed: false,
	}

	tasks = append(tasks, task)
	Save()
}

func Delete(id int) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			Save()
			return true
		}
	}

	fmt.Println("Task not found")
	return false
}

func ToggleCheck(id int, check bool) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = check
			Save()
			return true
		}
	}

	fmt.Println("Task not found")
	return false
}

func Save() {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
