package main

import (
	"bufio"
	"fmt"
	"strings"
)

func loopApp(scanner *bufio.Scanner, firstRun bool) {
	for {
		if firstRun {
			printIntro()
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

			addedItems := findItemsToAdd(fields)
			printAddedItems(addedItems)
		case "удалить":
			if fieldsLength == 1 {
				fmt.Println("Вы не ввели значения, которое нужно удалить")
				continue
			}

			deletedItems := extractDisallowedItems(fields)
			printDeletedItems(deletedItems)
		case "список":
			if len(items) == 0 {
				fmt.Println("Ваш список пуст")
				continue
			}

			fmt.Printf("Ваш список: \n- %s\n", strings.Join(items, "\n- "))
		case "help":
			printHelpCommandOutput()
		case "выйти":
			fmt.Println("До скорого!")
			return
		default:
			fmt.Println("Вы ввели неизвестную команду")
		}
	}
}
