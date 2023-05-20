package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

type StructTest struct {
	f1     string
	f2, f3 int
}

type smallStruct struct {
	f1 string
	f2 int
}

func simpleStruct() {
	s1 := StructTest{
		f1: "Поле f1",
		f2: 11,
		f3: 10,
	}

	fmt.Println("Structs.")
	fmt.Println(s1)
	fmt.Println(s1.f1)

	s1.f3 = 33
	fmt.Println(s1)
}

// Не инициализированные явно поля будут иметь zero value
func zeroValues() {
	s2 := StructTest{
		f1: "F1",
		f2: 0,
		f3: 0,
	}
	fmt.Println(s2)
}

// Вывоз new и использование & вместе с литеральным значением - эквивалентны
func newAndLiteral() {
	var s3 = new(StructTest)
	var s4 = &StructTest{}
	s3.f1 = "S3"
	s4.f1 = "S4"
	fmt.Println(*s3, *s4)
}

// Использование позиционного присваивания значений
func positionalAssigment() {
	s5 := StructTest{"S5", 5, 10}
	fmt.Println(s5)
}

// Embedded struct field - без имени поля, только тип
// Без имён могут быть не только другие struct, но и примитивные типы
func embeddedFields() {
	type StructExample2 struct {
		// embedded field
		StructTest
		// обычное поле
		i1 int
		// а это опять embedded
		int
	}

	s6 := StructExample2{
		StructTest: StructTest{"S6", 77, 88},
		i1:         100,
		int:        200,
	}
	fmt.Println(s6)
	s6.StructTest.f1 = "Embedded"
	s6.int = 300
	fmt.Println(s6)
}

// В структуре может быть только одно embedded поле одного типа
func altFields() {
	type StructExample3 struct {
		StructTest            // embedded StructExample
		s1         StructTest // alt StructExample
		int                   // embedded int
		i          int        // alt int
	}

	v := StructExample3{
		StructTest: StructTest{
			f1: "f1-1",
			f2: 101,
		},
		s1: StructTest{
			f1: "f1-2",
			f2: 102,
		},
	}

	v.int = 111
	v.i = 103

	fmt.Println(v)
}

// Struct values are comparable if all their fields can be compared.
func compareStructs() {
	v1 := StructTest{
		f1: "100",
		f2: 100,
	}

	v2 := StructTest{
		f1: "102",
		f2: 100,
	}

	v3 := StructTest{
		f1: "100",
		f2: 100,
	}

	fmt.Println("v1 == v2", v1 == v2)
	fmt.Println("v1 == v3", v1 == v3)
}

// анонимная структура как параметр - имени и порядок полей
// должны совпадать с типом передаваемого параметра
func anonymousStruct(v struct {
	f1 string
	f2 int
}) {
	fmt.Println(v)
}

// Так гораздо нагляднее и понятнее, если сравнивать с
// анонимной структурой в предыдущем примере.
// Но зато в качестве параметра можно передавать только
// тип smallStruct (либо приведённый к нему)
func nonAnonymousStruct(v smallStruct) {
	pc, _, _, _ := runtime.Caller(1)
	fmt.Println(runtime.FuncForPC(pc).Name())
	fmt.Println(v)
}

// вывозы с использованием анонимного типа структуры vs неанонимный тип
func anonAndNonAnonStructs() {
	s := smallStruct{f1: "f1", f2: 100}
	anonymousStruct(s)

	type smallStruct2 struct {
		f1 string
		f2 int
	} // эта структура совпадает по названиям полей, типам и порядку полей
	// со структурой smallStruct

	s2 := smallStruct2{f1: "s2-f1", f2: 101}
	// здесь можно вызывать без приведения типов
	anonymousStruct(s2)

	nonAnonymousStruct(s)
	// а здесь нужно привести тип к нужному
	nonAnonymousStruct(smallStruct(s2))
}

// Создание анонимной структуры, инициализация значениями на месте
// и использование в качестве параметра
func anonStructLiteralAsArgument() {
	var builder strings.Builder
	// В функцию Encode передаём literal value анонимной структуры.
	// Поля, которые должны попасть в json - должны начинаться с заглавной буквы.
	json.NewEncoder(&builder).Encode(struct {
		Fld1         string
		Fld2         int
		privateField int
	}{
		Fld1:         "Test fld1 value",
		Fld2:         101,
		privateField: 102,
	})
	fmt.Println(builder.String())
}

// При присвоении значений элементам array, slice, map, которые состоят из структур -
// тип структуры указывать не обязательно.
func omitStructType() {
	slice := []StructTest{
		{
			f1: "Test-f1",
			f2: 101,
		},
		{
			f1: "Test-f1-1",
			f2: 102,
		},
	}
	fmt.Println(slice)

	m := map[string]StructTest{
		"key1": {
			f1: "map-val1",
			f2: 101,
		},
		"key2": {
			f1: "map-val2",
			f2: 102,
			f3: 105,
		},
	}
	fmt.Println(m)
}

// Struct Pointer Convenience Syntax
func structPointer() {
	v1 := StructTest{
		f1: "f1-val",
		f2: 101,
	}

	// v2 - указатель на ту же область памяти, где находится v1
	v2 := &v1

	fmt.Println(v1, v2)
	fmt.Println(v1, *v2)
	fmt.Println(&v1, (*v2).f1)
	// При использовании указателей необязательно всякий раз использовать
	// синтаксис (*varName).fieldName - можно просто varName.fieldName
	fmt.Println(v2.f1)
}

// Вызов функции с параметром = указатель на literal sctruct value
func pointerToLiteralValue() {
	fmt.Println(literalFuncDemo(
		&StructTest{
			f1: "Pointer to literal struct value",
			f2: 55,
		}))
}

// Принимаем на вход pointer, и его же затем возвращем.
// Это позволяет вызывать функцию и передавать ей pointer на литеральное
// значение (без использования промежуточной переменной).
func literalFuncDemo(v *StructTest) StructTest {
	v.f2 *= 10
	return *v
}

// Конструктор структур
func newStructTest(f1 string, f2, f3 int) *StructTest {
	return &StructTest{
		f1: f1,
		// Порядок полей - не обязательно как объявлен в типе
		f3: f3,
		f2: f2,
	}
}

func constructStruct() {
	v := newStructTest("constructor", 101, 102)
	fmt.Println(*v, v.f3)

	// Slice из указателей на структуры, создаётся из вызовов конструкторов
	s := []*StructTest{
		newStructTest("s1", 10, 20),
		newStructTest("s2", 30, 40),
	}
	fmt.Println(s[0].f1, s[1].f2)
}

// Pointer как поле структуры
type StructWithPointer struct {
	sp1 int
	// указатель на структуру, и при этом это embedded field
	*StructTest
}

// Конструктор для StructWithPointer
func newStructWithPointer(sp1 int, st *StructTest) *StructWithPointer {
	return &StructWithPointer{sp1, st}
}

// Обратите внимание на то, как можно использовать (ссылаться на) структуру *StructTest
// внутри переменной sp - не нужно использовать синтаксис *StructTest, достаточно StructTest
func demoStructWithPointer() {
	sp := newStructWithPointer(100, newStructTest("demoStructWithPointer", 101, 102))

	fmt.Println(sp)
	fmt.Println(sp.StructTest)
	fmt.Println(sp.StructTest.f1)
}

// Паника с zero-values для неинициализированной структуры
func panicWithZeroValues() {
	var sp StructWithPointer
	// Структура sp инициализирована zero-values, для указателя это nil.
	// Печать целиком сработает, и для sp.StructTest будет выведено nil
	fmt.Println("Здесь будет паника, если обратиться к полю f1 вложенной структуры StructTest - потому что sp.StructTest == nil")
	fmt.Println(sp)
	// А здесь будет паника, потому что мы попытаемся обратиться к полю
	// указателя на структуру, а указатель == nil
	// fmt.Println(sp.StructTest.f1)
}

// Как избежать паники, связанной с zero-values в структурах
func noPanic() {
	// А для этого нужно обязательно поля pointer инициализировать
	var sp StructWithPointer = StructWithPointer{StructTest: &StructTest{}}
	fmt.Println(sp)
	fmt.Println("Здесь паники не будет, потому что sp.StructTest != nil - мы иницализировали эту стркутуру")
	fmt.Println("sp.StructTest.f1: ", sp.StructTest.f1)
}

func main() {
	simpleStruct()
	zeroValues()
	newAndLiteral()
	positionalAssigment()
	embeddedFields()
	altFields()
	compareStructs()
	anonAndNonAnonStructs()
	anonStructLiteralAsArgument()
	omitStructType()
	structPointer()
	pointerToLiteralValue()
	constructStruct()
	demoStructWithPointer()
	panicWithZeroValues()
	noPanic()
}
