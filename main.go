package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task_manager/src"
)

func main() {
	task := src.NewTaskManager()
	var (
		flag    = true
		request int
	)

	for flag {
		fmt.Println("Если Вы хотите добавить задачу нажмите 1")
		fmt.Println("Если Вы хотите выполнить задачу нажмите 2")
		fmt.Println("Если Вы хотите просмотреть задачи 3")
		fmt.Println("Если Вы посмотреть сколько задач Вы можете выполнить за отведенное время нажмите 4")
		fmt.Println("Если Вы хотите сменить приоритет какой-то из введенных задач нажмите 5")
		fmt.Println("Если Вы хотите завершить сеанс нажмите 6")

		fmt.Scan(&request)
		switch request {
		case 1:
			var (
				x        string
				priority int
				time     int
			)
			fmt.Scanln() // Очищаем буфер

			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Введите описание задачи:")
			x, _ = reader.ReadString('\n')
			x = strings.TrimSpace(x)
			fmt.Println("Введите приоритет задачи")
			fmt.Scan(&priority)
			fmt.Println("Введите время выполнения задачи")
			fmt.Scan(&time)
			task.Insert(x, priority, time)
		case 2:
			fmt.Println(task.Delete())
		case 3:
			task.View()
		case 4:
			var freeTime int
			fmt.Println("Введите свободное время для выполнения задач")
			fmt.Scan(&freeTime)
			fmt.Println(task.SelectTasksByTimer(freeTime))
		case 5:
			var (
				id       int
				priority int
			)
			fmt.Println("Введите id задачи, у которой вы хотите изменить приоритет:")
			fmt.Scan(&id)
			fmt.Println("Введите новое значение приоритета:")
			fmt.Scan(&priority)
			task.Queue.UpdatePriority(id, priority)
		case 6:
			flag = false
		}

	}
}
