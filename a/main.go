package a

import (
	"fmt"
	"youtangai/try_init/b"
)

var A string

func init() {
	A = b.B
	fmt.Println("a init", A)
}
