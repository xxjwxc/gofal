package main

import (
	"fmt"

	"github.com/xxjwxc/gofal/fractional"
)

func call() {
	tmp := fractional.Model(7, 12)
	tmp1 := fractional.Model(1, 12)
	fmt.Println(tmp.Add(*tmp1))

	tmp = fractional.Model(1, 4)
	tmp1 = fractional.Model(1, 4)
	fmt.Println(tmp.Sub(*tmp1).Verdict())

	tmp = fractional.Model(3, 4)
	tmp1 = fractional.Model(2, 3)
	fmt.Println(tmp.Mul(*tmp1))

	tmp = fractional.Model(3, 4)
	tmp1 = fractional.Model(2, 3)
	fmt.Println(tmp.Div(*tmp1))

	tmp = fractional.Model(1, 3)
	fmt.Println(tmp.Verdict())

	tmp.Add(*fractional.Model(1)).Mul(*tmp)
	fmt.Println(tmp)
}

func main() {
	call()
}
