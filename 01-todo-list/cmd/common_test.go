package cmd

import (
	"testing"
	"time"
)

func TestGetFileData(t *testing.T) {
	_, err := getFileData()
	if err != nil {
		t.Fatal(err)
	}
}

func TestWriteFileDate(t *testing.T) {
	rawTasks, err := getFileData()
	newTask := Task{0, "Test", time.Now(), false}
	err = writeFileData(addNewTaskToList(newTask, rawTasks))
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestReWriteFileDate(t *testing.T) {
	rawTasks, err := getFileData()
	newTask := Task{0, "Test1", time.Now(), false}
	err = rewriteFile(addNewTaskToList(newTask, rawTasks))
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestAddTask(t *testing.T) {
	newTask := Task{1, "Test", time.Now(), false}
	err := addTask(newTask)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestChangeStatusTask(t *testing.T) {
	// Set the task
	rawTasks, err := getFileData()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	newTask := Task{0, "Test", time.Now(), false}
	err = writeFileData(addNewTaskToList(newTask, rawTasks))
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	rawTasks, err = getFileData()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	_, err = changeStatusTask(0, true)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestRemoveTask(t *testing.T) {
	// Set the task
	rawTasks, err := getFileData()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	newTask := Task{910, "Test", time.Now(), false}
	err = writeFileData(addNewTaskToList(newTask, rawTasks))
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	err = removeTask(910)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestGetListTasks(t *testing.T) {
	rawTasks, err := getFileData()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	_, err = getListTasks(rawTasks)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
