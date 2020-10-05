package main

import (
	"fmt"
	runner_pattern "learn-concurrency/runner-pattern/runner"
	"time"
)

func createTask() func(int) {
	return func(id int) {
		fmt.Println("Đang xử lý nhiệm vụ", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}

func main() {
	duration := time.Duration(5 * time.Second)
	newRunner := runner_pattern.New(duration)

	task1 := createTask()
	task2 := createTask()
	task3 := createTask()

	newRunner.AddTask(task1, task2, task3)

	err := newRunner.Start()
	fmt.Println(err)
}
