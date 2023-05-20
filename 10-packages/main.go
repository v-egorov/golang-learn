package main

import (
	"fmt"
	"packages-test/store"
	// VS Code будет жаловаться, что не может найти пакет <module-name>/store,
	// если открыть не текущую директорию, как корневую, а уровнем выше (то есть весь репозиторий)
	// https://stackoverflow.com/questions/58518588/vscode-could-not-import-golang-package
	// Для того, чтобы обойти, можно использовать workspaces - go 1.18+ (пока не удалось понять как это
	// должно работать).
	// Проблема не проявляется, если открыть в VS Code директорию 10-packages, то есть не открывать
	// 10-packages как часть репозитория. Если в VS Code открыта корневая директория репозитория, то
	// будем видеть во вкладке Problems вот такое:
	// could not import packages-test/store (cannot find package "packages-test/store" in any of ....
	// "packages-test/store"
)

func main() {
	fmt.Println("Hello modules and packages")

	product := store.Product{
		Name:     "Test1",
		Quantity: 10,
	}

	product2 := store.Product2{
		Name2:  "Name2 value",
		Price2: 11,
	}

	fmt.Println(product, product2)
}
