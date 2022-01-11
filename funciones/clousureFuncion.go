//clousure --> clausura

package main

import "fmt"

//var x int

func main() {
	var x int
	fmt.Println(x)
	x += 1
	fmt.Println(x)

	{ //encerando la variable y <-- no es visible fuera del scope
		y := 42
		fmt.Println(y)
	}

	foo()
}
func foo() {
	fmt.Println("Hola")
	//x += 1
}
