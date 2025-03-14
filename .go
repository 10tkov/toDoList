package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tasks := []string{}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("_____")
		fmt.Println("prod by 10tkov")
		fmt.Print("Ведите команду! ")
		fmt.Println("Доступные команды: create, read, update, delete, exit")
		fmt.Println("_____")

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "create":
			fmt.Print("Введите задачу: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)

			if task == "" {
				fmt.Println("Задача не может быть пустой.")
			} else {
				tasks = append(tasks, task)
				fmt.Println("Задача добавлена!")
			}
		}
		switch command {
		case "read":
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст.")
			} else {
				fmt.Println("Список задач:")
			}
			for i, task := range tasks {
				fmt.Println("_____")
				fmt.Printf("%d: %s\n", i+1, task)
			}
		}
		switch command {
		case "update":
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст. Нечего обновлять.")
			} else {
				fmt.Print("Введите название задачи, которую хотите обновить: ")
				oldTask, _ := reader.ReadString('\n')
				oldTask = strings.TrimSpace(oldTask)

				found := false
				for i, task := range tasks {
					if task == oldTask {
						fmt.Print("Введите новое название задачи: ")
						newTask, _ := reader.ReadString('\n')
						newTask = strings.TrimSpace(newTask)

						if len(newTask) < 3 {
							fmt.Println("Название задачи должно содержать не менее 3 символов.")
						} else {
							tasks[i] = newTask
							fmt.Println("Задача обновлена!")
						}
						found = true
						break
					}
				}

				if !found {
					fmt.Println("Задача не найдена.")
				}
			}
		}
		switch command {
		case "delete":
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст.")
			} else {
				fmt.Print("Введите название задачи, которую хотите удалить: ")
				taskToDelete, _ := reader.ReadString('\n')
				taskToDelete = strings.TrimSpace(taskToDelete)

				found := false
				for i, task := range tasks {
					if task == taskToDelete {
						tasks = append(tasks[:i], tasks[i+1:]...)
						fmt.Println("Задача удалена!")
						found = true
						break
					}
				}
				if !found {
					fmt.Println("Задача не найдена.")
				}
			}
		case "exit":
			fmt.Println("Выход из программы. Всего хорошего!")
			return

		}
	}
}
