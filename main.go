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

	for {
		if firstRun {
			fmt.Println("Список доступных команд:")
			fmt.Println("- help")
			fmt.Println("- добавить {яблоко пальто книга}")
			fmt.Println("- удалить {пальто яблоко}")
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
		items := make([]string, 0)

		switch cmd {
		case "добавить":
			switch fieldsLength {
			case 2:
				items = append(items, fields[1])
				fmt.Println("Вы добавили", fields[1])
			case 3:
				word1 := fields[1]
				word2 := fields[2]

				items = append(items, word1, word2)
				fmt.Printf("Вы добавили %s и %s\n", word1, word2)
			default:
				var args strings.Builder

				for i := 1; i < fieldsLength; i++ {
					items = append(items, fields[i])

					switch i {
					case fieldsLength - 1:
						args.WriteString(" и ")
						args.WriteString(fields[i])
					case fieldsLength - 2:
						args.WriteString(fields[i])
					default:
						args.WriteString(fields[i])
						args.WriteString(", ")
					}
				}

				fmt.Println("Вы добавили", args.String())
			}
		case "удалить":
			fmt.Println("Вы хотите удалить что-то...")
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
