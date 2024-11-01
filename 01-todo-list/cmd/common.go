package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
)

type Task struct {
	id         int
	name       string
	created    time.Time
	isComplete bool
}

var done = map[bool]string{true: "DONE", false: "UNDONE"}

var header []string = []string{"Id", "Name", "Created", "IsComplete"}

func getFileData() ([][]string, error) {
	file, err := os.OpenFile("allTasks.csv", os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}
	result := csv.NewReader(file)
	result.FieldsPerRecord = -1
	records, err := result.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}
	return records, nil
}

func writeFileData(tasks [][]string) error {
	file, err := os.OpenFile("allTasks.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("Error opening file: %v", err)
	}
	w := csv.NewWriter(file)
	w.WriteAll(tasks)
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}
	return err
}

func rewriteFile(tasks [][]string) error {
	file, err := os.OpenFile("allTasks.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("Error opening file: %v", err)
	}
	w := csv.NewWriter(file)
	w.WriteAll(tasks)
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}
	return nil
}

func addTask(task Task) error {
	rawTask, err := getFileData()
	if err != nil {
		return err
	}
	tasks := addNewTaskToList(task, rawTask)
	err = writeFileData(tasks)
	return err
}

func changeStatusTask(taskId int, status bool) (Task, error) {
	rawTasks, err := getFileData()
	if err != nil {
		return Task{}, err
	}
	tasks, err := getListTasks(rawTasks)
	number := -1
	for i, task := range tasks {
		if task.id == taskId {
			number = i
			tasks[i].isComplete = status
		}
	}
	err = rewriteFile(convertListTaskToListString(tasks))
	if err != nil {
		return Task{}, err
	}
	if number == -1 {
		return Task{}, fmt.Errorf("Task is not found")
	}

	return tasks[number], nil
}

func removeTask(taskId int) error {
	rawTasks, err := getFileData()
	if err != nil {
		return err
	}
	tasks, err := getListTasks(rawTasks)
	beforeDeleteNumber := len(tasks)
	for i, task := range tasks {
		if task.id == taskId {
			tasks = slices.Delete(tasks, i, i+1)
		}
	}
	afterDeleteNumber := len(tasks)
	if beforeDeleteNumber == afterDeleteNumber {
		return fmt.Errorf("The task doesn't exist to be delete")
	}
	err = rewriteFile(convertListTaskToListString(tasks))
	return err
}

func convertListTaskToListString(tasks []Task) [][]string {
	// fmt.Println(strconv.Itoa(len(tasks)))
	result := make([][]string, len(tasks)+1)
	result[0] = header
	for i, task := range tasks {
		result[i+1] = []string{strconv.Itoa(task.id), task.name, task.created.Format(time.RFC1123Z), strconv.FormatBool(task.isComplete)}
	}
	return result
}

func addNewTaskToList(task Task, allTasks [][]string) [][]string {
	// Header of the csv file
	result := make([][]string, 1)
	if allTasks == nil {
		result = make([][]string, 2)
		result[0] = header
		result[1] = []string{strconv.Itoa(task.id), task.name, task.created.Format(time.RFC1123Z), strconv.FormatBool(task.isComplete)}
	} else {
		// Add data rows
		result[0] = []string{strconv.Itoa(task.id), task.name, task.created.Format(time.RFC1123Z), strconv.FormatBool(task.isComplete)}
	}
	return result
}

// TransformSliceToData
func getListTasks(allTasks [][]string) ([]Task, error) {
	tasks := make([]Task, 0)
	for i := 1; i < len(allTasks); i++ {
		parseTime, err := time.Parse(time.RFC1123Z, allTasks[i][2])
		if err != nil {
			return nil, fmt.Errorf("Error to parseTime: %v", err)
		}
		parseInt, err := strconv.Atoi(allTasks[i][0])
		if err != nil {
			return nil, fmt.Errorf("Error to parseInt: %v", err)
		}
		parseBool, err := strconv.ParseBool(allTasks[i][3])
		if err != nil {
			return nil, fmt.Errorf("Error to parseBool: %v", err)
		}
		tasks = append(tasks,
			Task{
				parseInt,
				allTasks[i][1],
				parseTime,
				parseBool,
			})
	}
	return tasks, nil
}
