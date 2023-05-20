package main

import (
	"fmt"
	"reflect"
)

func f1(p1, p2 int) {
	fmt.Println("--- f1 ---")
	v := p1 * p2
	fmt.Println(v)
}

// Пустой параметр может использоваться, например, при реализации интерфейсов,
// когда нужно показать, к-л параметр не используется
func blankParam(p1 int, _ int) {
	fmt.Println("--- blankParam ---")
	fmt.Println(p1 * 10)
}

// Параметры могут не имен вовсе, тогда из значения невозможно достать внутри функции
// Такое опять же используется вместе с интерфейсами
func noParamNames(int, string, *[]byte) {
	fmt.Println("--- noParamNames ---")
	fmt.Println("No param names at all")
}

// variadic parameter должен быть _последним_ в списке параметров
func variadicParams(p1 string, strVals ...string) {
	fmt.Println("--- variadicParams ---")
	fmt.Println("type of strVals -", reflect.TypeOf(strVals))
	fmt.Println("len of strVals =", len(strVals))
	for i, strVal := range strVals {
		fmt.Println(i, strVal)
	}
}

func swapByValue(p1, p2 int) {
	fmt.Println("--- swapByValue ---")
	fmt.Println("Before swap: ", p1, p2)
	p := p1
	p1 = p2
	p2 = p
	fmt.Println("After swap: ", p1, p2)
}

func swapByRef(p1, p2 *int) {
	fmt.Println("--- swapByRef ---")
	fmt.Println("Before swap: ", *p1, *p2)
	p := *p1
	*p1 = *p2
	*p2 = p
	fmt.Println("After swap: ", *p1, *p2)
}

func returnOneValue(i int) int {
	fmt.Println("--- returnOneValue ---")
	return i * i
}

func returnTwoValues(i, j int) (int, int) {
	fmt.Println("--- returnTwoValues ---")
	// как будто swap
	return j, i
}

func namedResult(i ...int) (count, product int) {
	fmt.Println("--- namedResult ---")
	product = 1
	for _, v := range i {
		count += v
		product *= v
	}
	return
}

func deferFunc() int {
	fmt.Println("---- deferFunc ----")
	defer fmt.Println("1st defer call")
	result := 0
	for i := 0; i < 10; {
		result += i
		i++
	}
	defer fmt.Println("2nd defer call")
	fmt.Println("Function about to return")
	return result
}

func main() {
	f1(132, 345)
	blankParam(10, 20)
	noParamNames(10, "test", &[]byte{10, 20})
	variadicParams("p1", "p2", "p3", "p4")

	// в слайс strVals будет передан nil
	variadicParams("p11")

	s := []string{"p22", "p33", "p44", "p55"}
	// а вот так на место variadic параметра передаётся готовый слайс
	variadicParams("p1", s...)

	fmt.Println("----------")
	fmt.Println("---------- Swaps")
	p1 := 10
	p2 := 20

	fmt.Println("Before call: ", p1, p2)
	swapByValue(p1, p2)
	fmt.Println("After call: ", p1, p2)

	fmt.Println(" Swap by refernce ---")
	fmt.Println("Before call: ", p1, p2)
	swapByRef(&p1, &p2)
	fmt.Println("After call: ", p1, p2)

	fmt.Println(returnOneValue(17))

	p1 = 11
	p2 = 22
	p1, p2 = returnTwoValues(p1, p2)
	fmt.Println(p1, p2)

	fmt.Println(namedResult(1, 2, 3, 4, 5, 10))

	_, r2 := namedResult(5, 7, 8, 19)
	fmt.Println("first result discarded, second =", r2)

	fmt.Println(deferFunc())
}
