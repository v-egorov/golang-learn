package main

import "fmt"

// Это функция, на вход принимает указатель на структуру.
func printContent(demoStruct *DemoStruct) {
	fmt.Println("f1: ", demoStruct.f1, "f2: ", demoStruct.f2)
}

// А это метод структуры.
// (demoStruct *DemoStruct) - называется receiever, printContent - имя метода.
// receiever demoStruct внутри функции ничем не отличается от других параметров
// Методы могут быть объявлены для любого типа, объявленного в текущем пакете.
// Ресиверы могут быть pointer receiver - пример ниже printContent, или
// value reciever - valueRecieverFunc. Отличаются способом передачи ресивера - по значению
// или по ссылке.
func (demoStruct *DemoStruct) printContent() {
	fmt.Println("f1: ", demoStruct.f1, "f2: ", demoStruct.f2)
}

func (d DemoStruct) valueRecieverFunc() {
	fmt.Println("valueRecieverFunc: ", d.f1)
}

// Метод с параметром, возвращающий результат
func (d *DemoStruct) addF2(i int) int {
	return d.f2 + i
}

// Go не поддерживает methods overloading внутри одной структуры, но
// внутри разных типов структур могут быть методы с одинаковым именем.
func (a *AnotherStruct) printContent() {
	fmt.Println("af1:", a.af1, "af2:", a.af2, "af3:", a.af3)
}

// Метод для типа-алиаса
func (a TypeAlias) methodForAlias() map[int]string {
	result := make(map[int]string)
	for _, v := range a {
		result[v.f2] = v.f1
	}
	return result
}

func withFunction() {
	for _, d := range demoStructs {
		printContent(d)
	}
}

func withMethod() {
	for _, d := range demoStructs {
		// так вызываются методы
		d.printContent()
	}
}

func noMethos() {
	fmt.Println(demoStructs)
	for _, d := range demoStructs {
		fmt.Println(d.f1, d.f2)
	}
}

// Имплементация интерфейса для структур.
// Если интерфейс имплементирован методами - value recievers, то переменной типа "инетерфейс"
// можно будет присваиваить как значения других переменных типов, имплементирующих интерфейс,
// так и ссылки на переменные таких типов. В зависимости от того, использовано ли значение или
// ссылка - в метод, импелементирующий метод интерфейса также придёт либо значение (и тогда
// в методе будет использована копия ориганальной переменной), либо ссылка. Примеры ниже.
func (d DemoStruct) getString() string {
	return d.f1
}

func (d DemoStruct) getSum(addValue int) int {
	return d.f2 + addValue
}

func (a AnotherStruct) getString() string {
	return a.af1 + " : " + a.af2
}

func (a AnotherStruct) getSum(i int) int {
	return a.af3 + i
}

// Для импленемтации интерфейса нужно, чтобы все методы были либо value, либо pointer recievers,
// смешивать нельзя.
// В случае, когда имплементация интерфейса делается через pointer recievers - важно попимать,
// что интенрфейс импелементирует не тип struct, а именно *struct - в примере ниже
// интерфейс DemoInterface имплементирован типом *PointerRecieverStruct, а не PointerRecieverStruct.
// Такая импелементация гарантирует, что вызываться методы будут с параметром-ссылкой, то есть
// копирования значения параметра в новую переменную локально внутри имплементации метода
// не произойдёт.
func (p *PointerRecieverStruct) getString() string {
	return p.ps
}

func (p *PointerRecieverStruct) getSum(_ int) int {
	return 0
}

// Для структуры AnotherStruct опредеён метод getString2 - которого нет в DemoInterface.
// Но это не мешает AnotherStruct имплементировать интерфейс DemoInterface.
// Таким образом, структура может имплементировать несколько интерфейсов.
func (a AnotherStruct) getString2() string {
	return a.af2
}

// Использование интерфеса для создания экземпляров
func interfaceAsAnInstance() {
	// В slice можно положить экземпляры, которые реализуют один интерфейс,
	// но при этом эти экземпляры могут быть разных типов - главное, чтобы эти типы
	// реализовывали нужный интерфейс.
	// Переменные типа interface имеют два типа: static и dynamic.
	// Static type - это тип интерфейса (DemoInterface в данном случае).
	// Dynamic type - это тип присвоенного переменной значения, которое имплементирует интерфейс.
	s := []DemoInterface{
		DemoStruct{
			f1: "DemoStruct - f1.1",
			f2: 1,
		},
		DemoStruct{
			f1: "DemoStruct - f1.2",
			f2: 2,
		},
		AnotherStruct{
			af1: "AnotherStruct - af1.1",
			af2: "AnotherStruct - af2.1",
			af3: 1,
		},
	}

	for _, v := range s {
		// Внутри цикла мы имеем дело только с типом static - т.е. v имеет тип DemoInterface
		fmt.Println("Get String: ", v.getString(), "Get Sum: ", v.getSum(77))
	}
}

// Методы, объявленные как value receivers - работают с копией структуры
func valueMethodReceiver() {
	d := DemoStruct{
		f1: "demo-1",
		f2: 1,
	}

	// При создании экземпляра i создаётся копия структуры d
	var i DemoInterface = d

	// Можно присвоить и &d - то тогда в i2 будет не скопированы значения, в передана ссылка
	var i2 DemoInterface = &d

	d.f1 = "changed f1"
	// значение в d - изменилось
	fmt.Println(d)
	// в i - осталсь оригинальная копия
	fmt.Println(i)
	// а в i2 - ссылка на d, поэтому обращение к i2 вернет
	// то же изменённое значение, что и в d
	fmt.Println(i2)
}

// Методы, объявленные как pointer recievers - работают со ссылкой на оригинал
func pointerMethodReciever() {
	p := PointerRecieverStruct{
		ps: "Pointer Reciever Struct",
	}

	// Вот тут можно будет использовать только &p - потому что методы объявлены как pointer recievers
	var i DemoInterface = &p
	p.ps = "Changed Pointer Reciever Struct"
	fmt.Println(p)
	fmt.Println(i)
}

// Интерфейсы равны в том случае, если у них одинаковый динамичкский тип и все другие их поля равны.
func compareInterfaces() {
	var v1 DemoInterface = &DemoStruct{
		f1: "f1",
		f2: 1,
	}

	var v2 DemoInterface = &DemoStruct{
		f1: "f1",
		f2: 1,
	}

	fmt.Println("v1 == v1", v1 == v2) // Будет false, потому что v1 и v2 - указатели.
	// А указатели равны только в том случае, когда они указывают на одну и ту же область памяти.

	var v3 DemoInterface = AnotherStruct{
		af1: "1",
		af2: "2",
		af3: 3,
	}

	var v4 DemoInterface = AnotherStruct{
		af1: "1",
		af2: "2",
		af3: 3,
	}

	fmt.Println("v3 == v4", v3 == v4) // А здесь true - не укзатели, и все поля равны

	// При попытке сравнить несравнимые типы - получим run-time error. Например, слайсы - non-comparable,
	// поэтому если в структуре будет поле слайс, то попытка сравнить такие структуры приведет к RTE - panic
}

// Type Assertion - приведение от более общего типа (интерфейса) к частному (стурктуре).
// Type assertion может быть применён только к интерфейсу.
// Не путать с Type Conversion - это другое.
func typeAssertion() {
	demoSlice := []DemoInterface{
		DemoStruct{
			f1: "s1",
			f2: 1,
		},
		DemoStruct{
			f1: "s2",
			f2: 2,
		},
		/* будет паника, если добавить сюда этот элемент слайса и затем сделать type assertion
		без проверки типов.
		AnotherStruct{
			af1: "s3.1",
			af2: "s3.2",
			af3: 3,
		},*/
	}

	for _, v := range demoSlice {
		// Здесь манипулируем интерфейсом
		fmt.Println(v.getString())

		// А здесь делаем type assertion - то есть приведение
		// от более общего статического типа (интерфейса) к более частному, динамическому типу
		// (структуре, которая этот интерфейс имплементирует)
		fmt.Println(v.(DemoStruct).f1, v.(DemoStruct).f2)
	}

	demo2 := []DemoInterface{
		DemoStruct{
			f1: "-- s1 --",
			f2: 1,
		},
		AnotherStruct{
			af1: "-- s2 --",
			af2: "s22",
			af3: 2,
		},
	}

	for _, v := range demo2 {
		// Тут паники не будет - проверяем успешность type assertion, в этом случае
		// type assertion возвращает 2 значения (comma ok idiom) - динамический тип и bool.
		if d, ok := v.(DemoStruct); ok {
			fmt.Println(d.f1, d.f2)
		} else {
			fmt.Println(v.getString())
		}

		// Внутри switch можно использовать конструкцию v.(type) - более наглядный вариант, чем if
		// v.(type) называется Special type assertion
		switch vt := v.(type) {
		case DemoStruct:
			fmt.Println("Dynamic type: DemoStruct", vt.f1)
		case AnotherStruct:
			fmt.Println("Dynamic type: AnotherStruct", vt.af1)
		}

	}
}

// Пустой интерфейс (интерфейс без методов) может быть представлен (совместим с) любым типом,
// включая встроенные в язык типы.
func emptyInterface() {
	fmt.Println("Empty Interface")

	// В слайсе data могут лежать элементы любых типов, потому что он объявлен как слайс
	// пустых интерфейсов, с которым совместимы все типы.
	data := []interface{}{
		100,
		true,
		"Это строка",
		DemoStruct{
			f1: "1",
			f2: 1,
		},
		AnotherStruct{
			af1: "2",
			af2: "2",
			af3: 2,
		},
	}

	for _, v := range data {
		switch value := v.(type) {
		case string:
			fmt.Println(value)
		case int, bool:
			fmt.Println(value)
		case DemoStruct:
			fmt.Println(value.getString())
		case AnotherStruct:
			fmt.Println(value.getString2())
		}

	}

	// можем передать slice
	emptyInterfaceAsParameter(data...)

	// или любое количество параметров
	emptyInterfaceAsParameter(100,
		200,
		"Test string",
		&DemoStruct{
			f1: "1",
			f2: 1,
		},
		DemoStruct{
			f1: "2",
			f2: 2,
		},
		&AnotherStruct{
			af1: "3",
			af2: "3",
			af3: 3,
		},
	)
}

// В items может быть передан как слайс, внутри которго могут лежать значения
// любых типов, так и любое количество атомарных параметров любых типов - см примеры выше.
func emptyInterfaceAsParameter(items ...interface{}) {
	fmt.Println("emptyInterfaceAsParameter")
	for _, item := range items {
		switch value := item.(type) {
		case int, bool, string:
			fmt.Println("встроенный тип ", value)
		// Внимание на порядок следования case - сначала более узкий тип, затем более широкий
		case DemoStruct:
			fmt.Println("DemoStruct:", value.addF2(101))
		case *DemoStruct:
			fmt.Println("*DemoStruct:", value.f1)
		case DemoInterface:
			fmt.Println("DemoInterface", value.getString())
		}
	}
}

func main() {
	fmt.Println("Methods and Intergaces")
	noMethos()
	withFunction()
	withMethod()

	// Вызов метода
	fmt.Println(demoStructs[1].addF2(55))

	// Вызов метода с тем же именем, но в другом типе
	anotherStructs[1].printContent()

	v := DemoStruct{
		f1: "Test f1",
		f2: 105,
	}
	// Метод, где receiever - это указатель, может быть вызван и для переменной,
	// которая не является указателем. Обратный процесс также рабоатет -
	// если receiever объявлен как value, можно метод вызвать и для пойнтера.
	v.printContent()
	// Работает и так - но смысла нет
	(&v).printContent()

	// Необычная особенность - если метод объявлен с value reciever,
	// то вызвать его можно как в примере ниже (не путать со статическими
	// методами, например, в C# или Java - это другое):
	DemoStruct.valueRecieverFunc(DemoStruct{f1: "Demo value reciever"})
	DemoStruct.valueRecieverFunc(v)
	// Можно даже и вот так - по опять же зачем?
	(*DemoStruct).printContent(&v)
	// если можно вот так:
	v.printContent()

	ta := TypeAlias{
		&DemoStruct{
			f1: "f1-11",
			f2: 11,
		},
		{
			f1: "f1-22",
			f2: 22,
		},
	}
	// Вызов метода типа-алиаса
	fmt.Println(ta.methodForAlias())
	// Вызов этого же метода, но чуть по-другому - с приведением типа
	fmt.Println(TypeAlias(demoStructs).methodForAlias())

	// Интерфейсы для создания экземпляров slice
	interfaceAsAnInstance()
	// вызовем метод getString2 - чтобы убрать warning, что метод нигде не используется
	fmt.Println(anotherStructs[1].getString2())

	// Интерфейсы можно использовать как переменные, параметры функций
	// и результаты функций.
	// Методы не могут определены с интерфейсами в качестве recievers

	valueMethodReceiver()
	pointerMethodReciever()
	compareInterfaces()
	typeAssertion()
	emptyInterface()
}
