package scanner

import (
	"fmt"
	"solotodo/todo"
)

func printResult(result string) {
	fmt.Println("Результат выполнения команды: ", result)
	fmt.Println(" ")
}

func printPrompt() {
	fmt.Print("Введите команду: ")
}

func printExit() {
	fmt.Println("Завершение работы...")
}

func printAdd(title string) {
	fmt.Println("Задача '" + title + "' успешно добавлена")
	fmt.Println(" ")
}

func printTasks(tasks map[string]todo.Task) {
	for _, task := range tasks {
		fmt.Println("Заголовок:", task.Title)
		fmt.Println("Текст:", task.Text)
		fmt.Println("Выполнена:", task.IsDone)
		fmt.Println("Создана:", task.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Println("---")
	}
}

func printDone(title string) {
	fmt.Println("Задача '" + title + "' выполнена")
	fmt.Println("")
}

func printDel(title string) {
	fmt.Println("Задача '" + title + "' удалена")
}

func printHelp() {
	fmt.Println("Список команд:")
	fmt.Println("1. help")
	fmt.Println("— эта команда позволяет узнать доступные команды")
	fmt.Println("")
	fmt.Println("2. add {заголовок задачи из одного слова} {текст}")
	fmt.Println("— эта команда позволяет добавлять новые задачи")
	fmt.Println("")
	fmt.Println("3. list")
	fmt.Println("— эта команда позволяет получить полный список задач")
	fmt.Println("")
	fmt.Println("4. del {заголовок существующей задачи}")
	fmt.Println("— эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("")
	fmt.Println("5. done {заголовок существующей задачи)")
	fmt.Println("— эта команда позволяет выполнить задачу")
	fmt.Println("")
	fmt.Println("6. events")
	fmt.Println("— эта команда позволяет получить список всех событий")
	fmt.Println("")
	fmt.Println("7. exit")
	fmt.Println("— эта-команда позволяет завершить выполнение программы")
	fmt.Println("")
}

func printEvents(events []Event) {
	fmt.Println("События: ", events)
	fmt.Println("")
}
