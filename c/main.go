package c

import (
	"fmt"
	"time"
)

var C string

func init() {
	time.Sleep(10 * time.Second)
	C = "hello"
	fmt.Println("c init", C)
}
