//retornar una función
package main

import "fmt"

func main() {

	fmt.Println(bar())
}

//funcion que retorna otra funcion
func bar() func() int {
	return func() int {
		return 1492
	}
}
