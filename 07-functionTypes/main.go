package main

import "fmt"

// калькулятор ф1
func f1(i int) (int, int) {
	fmt.Println("-- f1 --")
	return i + 10, i + 20
}

// калькулятор ф2
func f2(i int) (int, int) {
	fmt.Println("-- f2 --")
	return i * i, i * i * i
}

//  вызвать функцию-калькулятор
func callCalculator(m string, i int, calculator func(int) (int, int)) {
	fmt.Println("--- callCalculator ---")
	fmt.Println("Message: ", m)
	_, result := calculator(i)
	fmt.Println(result)
	fmt.Println()
}

// вычислить и вернуть функцию-калькулятор
// если b == true, то ф1, иначе ф2
func getCalculator(b bool) func(int) (int, int) {
	fmt.Println("--- Function as a result being called ---")
	fmt.Println("Input param value:", b)
	fmt.Println()
	if b {
		return f1
	} else {
		return f2
	}
}

// тип-алиас для типа функции-калькулятора.
// Сигнатура совпадает с ф1 и ф2
type calcFunc func(int) (int, int)

// Вызвать функцию-калькурятор.
// Версия с алиасом.
func callByAlias(m string, i int, c calcFunc) {
	fmt.Println("--- callByAlias ---")
	fmt.Println("Message:", m)
	// Версия с алиасом не игнорирует первый результат,
	// а выводит оба.
	fmt.Println(c(i))
	fmt.Println()
}

// вычислить и вернуть функцию-калькулятор.
// Версия с алиасом.
func getCalculatorAlias(b bool) calcFunc {
	if b {
		return f1
	} else {
		return f2
	}
}

// Возвращает функцию-калькулятор, которая определена как literal value
func getCalculatorLiteral() calcFunc {
	// это анонимная функция
	var c calcFunc = func(i int) (int, int) {
		return i + 100, i + 200
	}
	return c
}

// Возвращает функцию-калькулятор, которая определена как literal value.
// Версия 2 - без промежуточных переменных
func getCalculatorLiteral2(b bool) calcFunc {
	if b {
		return func(i int) (r1 int, r2 int) {
			r1 = i * 311
			r2 = i * 257
			return r1, r2
		}
	} else {
		return func(i int) (int, int) {
			return i * 77, i * 88
		}
	}
}

var m2exception, m2extended int = 1, 1

// Фабрика калькуляторов.
// m1 и m2 - значения, на которые будет умножаться входное значение,
// передаваемое калькулятору, чтобы вычислить два результата, которые
// возвращает калькулятор.
func calculatorFactory(m1, m2 int) calcFunc {
	// Внутри фабрики может быть своя логика, которая работает с
	// входными параметрами, которые затем могут быть переданы в
	// функцию, которая создаётся и возвращается фабрикой.
	// В данном случае - если m2 меньше, чем m2, то ставим m2 = m2exception,
	// и тогда калькулятор в качестве второго множителя получит m2exception,
	// и второй результат калькулятора будет вычисляться с использованием
	// глобальной переменной m2exception - при этом код в следующем if будет
	// исполняться только в момент создания функции в фабрике.
	// То есть m2exception вычисляется только один раз - при создании closure-функции
	// (и при условии, что m2 < m1).
	// А вот m2extended - вычисляется каждый раз при вызове closure-функции, потому что
	// ссылка на m2extended - внутри closure. Эффект - см. примеры ниже, где
	// используется calculatorFactory
	if m2 < m1 {
		m2 = m2exception
	}
	return func(i int) (int, int) {
		return i * m1, i * m2 * m2extended
	}
}

// Closing on value param.
// Поскольку при передаче параметров по значению - параметры копируются,
// можно зафиксировать значение m2extended на момент создания функции, передав
// m2extended в параметр m2ext
func calculatorFactoryOnValue(m1, m2, m2ext int) calcFunc {
	return func(i int) (int, int) {
		return i * m1, i * m2 * m2ext
	}
}

// Closing on pointer.
// Поскольку при передаче параметров по указателю - параметры не копируются,
// если передать ссылку на m2extended в m2ext, то изменения в m2extended будут влиять
// на результат вычисления всякий раз, когда функция будет вызываться
func calculatorFactoryOnPointer(m1, m2 int, m2ext *int) calcFunc {
	return func(i int) (int, int) {
		return i * m1, i * m2 * *m2ext
	}
}

func main() {
	fmt.Println("Function Types")

	// calculator - это переменная типа "функция"
	// с сигнатурой (int) (int, int)
	var calculator func(int) (int, int)
	fmt.Println("testFunc is nil? ", calculator == nil)

	// статическое присвоение значений калькулятору
	calculator = f1
	fmt.Println(calculator(10))

	calculator = f2
	fmt.Println(calculator(10))
	fmt.Println("calculator is nil? ", calculator == nil)

	// передаём параметр типа "функция" другой функции
	// калькулятор уже вызывается там
	callCalculator("Calculator f1", 11, f1)
	callCalculator("Calculator f2", 11, f2)

	// динамическое присвоение калькулатору значений
	// вариант 1 - присвоить переменной типа "функция" значение
	// и затем вызвать эту переменную
	calculator = getCalculator(false)
	fmt.Println(calculator(22))

	// вариант 2 - вызвать сразу функцию, которую получили
	fmt.Println(getCalculator(false)(22))

	// вариант 3 - передать полученное значение-функцию
	// в другую функцию и вызвать там
	callCalculator("func as a result", 33, getCalculator(true))

	// вариант 4 - с алиасом
	callByAlias("C алиасом.", 17, f2)

	// вариант 5 - вычислить функцию, вернуть алиасом
	// и тут же вызвать
	fmt.Println("getCalculatorAlias:")
	fmt.Println(getCalculatorAlias(true)(55))

	// вариант 6 - с возвратом литерала как функции
	fmt.Println("getCalculatorLiteral:")
	fmt.Println(getCalculatorLiteral()(1))

	// вариант 7 - с возвратом литерала как функции - closure
	fmt.Println("getCalculatorLiteral2: два раза")
	fmt.Println(getCalculatorLiteral2(true)(77))
	fmt.Println(getCalculatorLiteral2(false)(77))

	// вариант 8 - функция определена прямо внутри вызова
	// другой функции.
	fmt.Println("callByAlias with Literal function:")
	callByAlias("Literal func:", 17, func(i int) (int, int) {
		return i + 17, i + 18
	})

	// Вариант 9 - калькулятор создаётся через factory
	fmt.Println("callByAlias with Factory:")
	callByAlias("Factory", 17, calculatorFactory(11, 22))
	callByAlias("Factory", 17, calculatorFactory(22, 11))

	// Вариант 10 - результат от factory присвоим переменной
	// и затем вызовем эту переменную-функцию.
	fmt.Println("Call to calculator created by Factory")
	calculator = calculatorFactory(77, 55)
	fmt.Println(calculator(2))

	m2exception = 2 // не влияет
	m2extended = 10 // влияет
	fmt.Println(calculator(2))

	fmt.Println("Call to calculator created by FactoryOnValue")
	m2extended = 20 // влияет - передаётся в фабрику by value
	calculator = calculatorFactoryOnValue(1, 2, m2extended)
	fmt.Println(calculator(2))

	m2extended = 10 // не влияет - в функции копия предыдущего значения
	fmt.Println(calculator(2))

	fmt.Println("Call to calculator created by FactoryOnPointer")
	m2extended = 5 // влияет - передаётся в фабрику by address
	calculator = calculatorFactoryOnPointer(1, 2, &m2extended)
	fmt.Println(calculator(2))

	m2extended = 10 // влияет - передаётся в фабрику by address
	fmt.Println(calculator(2))

}
