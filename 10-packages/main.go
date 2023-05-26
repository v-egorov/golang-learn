package main

import (
	"fmt"
	// VS Code будет жаловаться, что не может найти пакет <module-name>/store,
	// если открыть не текущую директорию, как корневую, а уровнем выше (то есть весь репозиторий)
	// https://stackoverflow.com/questions/58518588/vscode-could-not-import-golang-package
	// Для того, чтобы обойти, можно использовать workspaces - go 1.18+ (пока не удалось понять как это
	// должно работать).
	// Проблема не проявляется, если открыть в VS Code директорию 10-packages, то есть не открывать
	// 10-packages как часть репозитория. Если в VS Code открыта корневая директория репозитория, то
	// будем видеть во вкладке Problems вот такое:
	// could not import packages-test/store (cannot find package "packages-test/store" in any of ....
	"packages-test/store"
)

func main() {
	fmt.Println("Hello modules and packages")

	product1 := store.Product{
		Name:     "Test1",
		Quantity: 10,
	}

	product2 := store.Product2{
		Name2:  "Name2 value",
		Price2: 11,
	}

	product3 := store.NewProduct("Test 3", 10, 12.3)

	fmt.Println(product1, product2, product3)

	// Это ошибка - поле price не экспортировано
	// fmt.Println(product3.price)

	// А метод Price() - экспортирован
	fmt.Printf("Price %f\n", product3.Price())

	fmt.Printf("Price with tax: %f\n", product3.PriceWithTax())
}
