package main

type DemoStruct struct {
	f1 string
	f2 int
}

var demoStructs = []*DemoStruct{
	{"f1-1", 100},
	{"f1-2", 101},
}

type AnotherStruct struct {
	af1 string
	af2 string
	af3 int
}

var anotherStructs = []*AnotherStruct{
	{
		af1: "af1-1",
		af2: "af2-1",
		af3: 1,
	},
	{
		af1: "af2-2",
		af2: "af2-2",
		af3: 2,
	},
}

// Алиас типа, для него тоже можно определить методы
type TypeAlias []*DemoStruct

// Интерфейс, который будет реализован для наших структур
type DemoInterface interface {
	getString() string
	getSum(addValue int) int
}

// Для этого типа будем имплементировать pointer recievers
type PointerRecieverStruct struct {
	ps string
}
