package main

import (
	"fmt"
)

func main() {

	var a int8 = -1
	var b int16 = 32767
	var c int32 = 2147483647
	var d int64 = 9223372036854775807

	fmt.Println("int8:", a)
	fmt.Println("int16:", b)
	fmt.Println("int32:", c)
	fmt.Println("int64:", d)

	message := "Hello World"

	fmt.Printf("type --> message: %T\n", message)
}
