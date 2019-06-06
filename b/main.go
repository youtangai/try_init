package b

import (
	"fmt"
	"youtangai/try_init/c"
)

var B string

func init() {
	B = c.C
	fmt.Println("b init", B)
}
