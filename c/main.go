package c

import "fmt"

var C string

func init() {
	C = "hello"
	fmt.Println("c init", C)
}
