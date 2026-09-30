package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	firstRun := true

	items := make([]string, 0)

	for {
		if firstRun {
			fmt.Println("Список доступных команд:")
			fmt.Println("- help")
			fmt.Println("- добавить {яблоко пальто книга}")
			fmt.Println("- удалить {пальто яблоко}")
			fmt.Println("- список")
			fmt.Println("- выйти")
			fmt.Println()

			firstRun = false
		}

		fmt.Print("Введите команду: ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Ошибка ввода:", err)
			}
			return
		}

		text := scanner.Text()
		fields := strings.Fields(text)
		fieldsLength := len(fields)

		if fieldsLength == 0 {
			continue
		}

		cmd := fields[0]

		switch cmd {
		case "добавить":
			if fieldsLength == 1 {
				fmt.Println("Вы не ввели значения, которое нужно добавить")
				continue
			}

			newItems := fields[1:]

			items = append(items, newItems...)

			if len(newItems) == 1 {
				fmt.Println("Вы добавили", newItems[0])
			} else {
				allButLast := strings.Join(newItems[:len(newItems)-1], ", ")
				last := newItems[len(newItems)-1]
				fmt.Printf("Вы добавили %s и %s\n", allButLast, last)
			}
		case "удалить":
			if fieldsLength == 1 {
				fmt.Println("Вы не ввели значения, которое нужно удалить")
			} else {
				var deletedItems []string

				for i := 1; i < fieldsLength; i++ {
					valueToRemove := fields[i]

					foundIndex := -1
					for idx, item := range items {
						if item == valueToRemove {
							foundIndex = idx
							break
						}
					}

					if foundIndex != -1 {
						deletedItems = append(deletedItems, valueToRemove)
						items = append(items[:foundIndex], items[foundIndex+1:]...)
					}
				}

				if len(deletedItems) > 0 {
					fmt.Printf("Вы удалили: %s\n", strings.Join(deletedItems, ", "))
				} else {
					fmt.Println("Ни один из указанных элементов не найден в списке")
				}
			}
		case "список":
			if len(items) == 0 {
				fmt.Println("Ваш список пуст")
				continue
			}

			fmt.Printf("Ваш список: \n- %s\n", strings.Join(items, "\n"))
		case "help":
			fmt.Println("Команда: help")
			fmt.Println("--- эта команда выводит список доступных команд")
			fmt.Println("")
			fmt.Println("Команда: добавить {что нужно добавить}")
			fmt.Println("--- эта команда позволяет добавлять что-либо")
			fmt.Println("")
			fmt.Println("Команда: удалить {что нужно удалить}")
			fmt.Println("--- эта команда позволяет удалять что-либо")
			fmt.Println("")
			fmt.Println("Команда: выйти")
			fmt.Println("--- эта команда позволяет выйти из приложения")
		case "выйти":
			fmt.Println("До скорого!")
			return
		default:
			fmt.Println("Вы ввели неизвестную команду")
		}
	}

}
