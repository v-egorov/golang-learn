package main

import (
	"fmt"
)

func pointers() {
	first := 100
	// в second - будет копия значения first
	second := first

	// third - это указатель на int, и в third мы записываем адрес переменной first
	var third *int = &first
	// то же самое, но это сокращённая форма объявления переменной
	fourth := &first

	// Увеличим значение first на 1, и увидим, что в копии - second - ничего не изменилось,
	// при этом *fird и *fourth - также будут изменены, потому что указатели
	// fird и fourth содержат [один и тот же] адрес переменной first
	first++
	fmt.Println(first, second, *third, *fourth)

	// В go нет арифметики указателей, здесь мы увеличиваем значение,
	// которое хранится в int, адрес которого в переменной fird
	*third++
	fmt.Println(first, second, *third, *fourth)

	// попытка сделать third++ закончится ошибкой компиляции - non-numeric type *int
	// third++
}

func main() {
	pointers()
}
