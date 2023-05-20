package main

import (
	"fmt"
	// VS Code будет жаловаться, что не может найти пакет <module-name>/store,
	// если открыть не текущую директорию, как корневую, а уровнем выше (то есть весь репозиторий)
	// https://stackoverflow.com/questions/58518588/vscode-could-not-import-golang-package
	// Для того, чтобы обойти, можно использовать workspaces - go 1.18+
	"packages-test/store"
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
