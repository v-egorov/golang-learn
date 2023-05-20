package main

import (
	"fmt"
	"strconv"
)

func ifStatement() {
	kayakPrice := 375.75
	fmt.Printf("Price: %v\n", kayakPrice)

	if kayakPrice > 500 {
		fmt.Printf("Price is grater than 500\n")
	} else if kayakPrice < 300 {
		fmt.Printf("Price is less than 300\n")
	} else {
		fmt.Printf("No match\n")
	}

	priceString := "354-"
	if kayakPrice, err := strconv.Atoi(priceString); err == nil {
		fmt.Printf("Price: %v\n", kayakPrice)
	} else {
		fmt.Printf("Error: %v\n", err)
	}
}

func forLoop() {
	counter := 0
	for {
		fmt.Printf("Counter: %d\n", counter)
		if counter++; counter > 3 {
			break
		}
	}

	fmt.Println()
	counter = 0
	for counter <= 3 {
		fmt.
			Printf("Counter: %d\n", counter)
		counter++
	}

	fmt.Println()
	for counter = 0; counter <= 3; counter++ {
		if counter == 1 {
			continue
		}
		fmt.Println("Counter:", counter)
	}
}

func rangeStatement() {
	product := "string"
	for index, character := range product {
		fmt.Println("Index:", index, "char: ", string(character))
	}

	products := []string{"s1", "s2", "s3"}
	for _, p := range products {
		fmt.Println(p)
	}
}

func switchStatement() {
	product := "Kayak"
	for i, c := range product {
		switch c {
		case 'K':
			fmt.Printf("Uppercase K at %d\n", i)
			fmt.Printf("UC\n")
			fallthrough
		case 'k':
			fmt.Printf("k at %d position\n", i)
			// fallthrough
		case 'y':
			fmt.Printf("y at %d position\n", i)
		}
	}
}

func main() {
	ifStatement()
	forLoop()
	rangeStatement()
	switchStatement()
}
