package main

import (
	"fmt"
	"reflect"
)

func arrays() {
	var names [3]string

	names[0] = "str1"
	names[1] = "str2"
	names[2] = "str3"

	fmt.Println(names)

	names2 := [3]string{"str4", "str5"}
	fmt.Println(names2)

	names3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(names3)
}

func copyArray() {
	fmt.Println(" -- copyArray ---")
	arr1 := [...]int{1, 2, 3}
	arr2 := arr1
	fmt.Println(arr1, arr2)

	arr2[0] = 7
	fmt.Println(arr1, arr2)
}

func arayPointer() {
	fmt.Println("--- arrayPointer ---")
	arr1 := [...]int{1, 2, 3}
	arr2 := &arr1
	fmt.Println(arr1, arr2)
	fmt.Println(arr1, *arr2)

	arr2[0] = 7
	fmt.Println(arr1, *arr2)
}

func slices() {
	fmt.Println("--- Slices ---")
	s := make([]string, 3)
	s[0] = "s0"
	s[1] = "s1"
	fmt.Println(s)

	s2 := []string{"s2-0", "s2-1", "s2-3"}
	fmt.Println(s2)
	fmt.Println(s2[1])
	s2 = append(s2, "test")
	fmt.Println(s2)
	fmt.Println(*&s2[0])
}

func twoSlices() {
	fmt.Println("--- 2 slices ---")
	s1 := make([]string, 3, 5)
	s1[0] = "0"
	s1[1] = "1"
	s1[2] = "2"

	// same array b/c capacity allow to use it
	s2 := append(s1, "3", "4")

	// different array
	// s2 := append(s1, "3", "4", "5")

	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)

	s2[0] = "7"
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
}

func appendSlices() {
	fmt.Println("--- appendSlices ---")
	s1 := []string{"0", "1"}
	s2 := []string{"2", "3", "4"}
	fmt.Println(append(s1, s2...))
}

func sliceFromSlice() {
	fmt.Println("--- sliceFromSlice ---")
	s := []string{"0", "1", "2", "3", "4"}
	s1 := s[1:]
	s2 := s1[2:3:4]
	fmt.Println(s, s1, s2)

	s2[0] = "99"
	fmt.Println(s, s1, s2)

	s2 = append(s2, "100", "101")
	s2[0] = "77"
	fmt.Println(s, s1, s2)
}

func sliceDelete() {
	fmt.Println("--- sliceDelete ---")
	s := []string{"0", "1", "2", "3", "4"}
	fmt.Println(s)

	s = append(s[:2], s[3:]...)
	fmt.Println(s)
}

func compareSlices() {
	fmt.Println("--- compareSlices ---")
	s1 := []string{"0", "1", "2"}
	//s2 := s1
	s2 := []string{"0", "1", "2", ""}

	fmt.Println(reflect.DeepEqual(s1, s2))
}

func slice2array() {
	fmt.Println("--- slice2array ---")
	s := []string{"0", "1", "2", "3"}
	arrayPointer := (*[3]string)(s)
	a := arrayPointer
	fmt.Println(s, arrayPointer, *a)
}

func maps() {
	fmt.Println("--- maps ---")
	m := make(map[string]int, 5)
	m["01"] = 10
	m["1-"] = 11
	m["2*"] = 12

	fmt.Println(m)

	m2 := map[string]int{
		"--": 101,
		"++": 102,
	}
	fmt.Println(m2)
	fmt.Println(m2["++"])

	fmt.Println(m2["0"])
	if value, ok := m2["0"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("no value")
	}

	delete(m2, "--")
	fmt.Println(m2)

	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func runes() {
	fmt.Println("--- runes ---")
	var r []rune = []rune("€test")
	var s1 string = string(r[0])
	fmt.Println(r)
	fmt.Println(s1)
}

func main() {
	arrays()
	copyArray()
	arayPointer()
	slices()
	twoSlices()
	appendSlices()
	sliceFromSlice()
	sliceDelete()
	compareSlices()
	slice2array()
	maps()
	runes()
}
