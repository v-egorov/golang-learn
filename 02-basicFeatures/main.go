package main

import (
	"fmt"
)

func pointers() {
	first := 100
	// в second - будет копия значения first
	second := first

	// fird - это указатель на int, и в fird мы записываем адрес переменной first
	var fird *int = &first
	// то же самое, но это сокращённая форма объявления переменной
	fourth := &first

	// Увеличим значение first на 1, и увидим, что в копии - second - ничего не изменилось,
	// при этом *fird и *fourth - также будут изменены, потому что указатели
	// fird и fourth содержат [один и тот же] адрес переменной first
	first++
	fmt.Println(first, second, *fird, *fourth)

	// В go нет арифметики указателей, здесь мы увеличиваем значение,
	// которое хранится в int, адрес которого в переменной fird
	*fird++
	fmt.Println(first, second, *fird, *fourth)

	// попытка сделать fird++ закончится ошибкой компиляции - non-numeric type *int
	// fird++
}

func main() {
	pointers()
}
