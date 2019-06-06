package main

import (
	"fmt"
	"youtangai/try_init/a"
)

func init() {
	fmt.Println("main init", a.A)
}

func main() {
	fmt.Println("main main", a.A)
}
