package main

import (
	"fmt"
	"math"
	"strconv"
)

func convFromStr() {
	val1 := "1"
	val2 := "121"

	if bool1, err := strconv.ParseBool(val1); err == nil {
		fmt.Println("Parsed value: ", bool1)
	} else {
		fmt.Println("Conversion error: ", err)
	}

	if int1, err := strconv.Atoi(val2); err == nil {
		var intResult int = int1
		fmt.Println("Atoi: ", intResult)
		fmt.Printf("FormatInt: %s\n", strconv.FormatInt(int64(int1), 36))
	} else {
		fmt.Println("Atoi conversion error: ", err)
	}
}

func con2Str() {
	val1 := 161
	val2 := 3456.3579

	str1 := strconv.FormatInt(int64(val1), 16)
	str2 := strconv.Itoa(val1)
	str3 := strconv.FormatFloat(val2, 'E', -1, 64)

	fmt.Printf("FormatInt: %s\n", str1)
	fmt.Printf("Itoa: %s\n", str2)
	fmt.Printf("FornatFloat: %s\n", str3)
}

func overflow() {
	var intVal = math.MaxInt64
	var floatVal = math.MaxFloat64

	fmt.Println(intVal, intVal+1, intVal*2, intVal+(intVal-1))
	fmt.Println(floatVal, floatVal+1e291, floatVal*2)
	fmt.Println(math.IsInf((floatVal * 2), 0))
}

func main() {
	overflow()
	convFromStr()
	con2Str()
}
