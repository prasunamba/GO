package module1

import "fmt"

type mystructt struct {
	name string
	age  int
}

func Swap() {
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println("before ", a, b, c)
	a, b, c = b, c, a
	fmt.Println("after ", a, b, c)
}
