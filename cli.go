package main

import (
	"fmt"
	"strings"
)

func printIntro() {
	fmt.Println("Список доступных команд:")
	fmt.Println("- help")
	fmt.Println("- добавить {яблоко пальто книга}")
	fmt.Println("- удалить {пальто яблоко}")
	fmt.Println("- список")
	fmt.Println("- выйти")
	fmt.Println("")
}

func printAddedItems(addedItems []string) {
	if len(addedItems) == 1 {
		fmt.Println("Вы добавили", addedItems[0])
	} else if len(addedItems) > 1 {
		allButLast := strings.Join(addedItems[:len(addedItems)-1], ", ")
		last := addedItems[len(addedItems)-1]
		fmt.Printf("Вы добавили %s и %s\n", allButLast, last)
	} else {
		fmt.Println("Все указанные элементы уже есть в списке")
	}
}

func printDeletedItems(deletedItems []string) {

	if len(deletedItems) > 0 {
		fmt.Printf("Вы удалили: %s\n", strings.Join(deletedItems, ", "))
	} else {
		fmt.Println("Ни один из указанных элементов не найден в списке")
	}
}

func printHelpCommandOutput() {
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
}
