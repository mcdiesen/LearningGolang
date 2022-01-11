//funcion que retorna una funcion

package main

import "fmt"

func main() {
	f := foo()
	fmt.Printf("%T\n", f)
}
func foo() func() int {
	return func() int { //foo retorna una funcion que retorna un entero
		return 42
	}
}
