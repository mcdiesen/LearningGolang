//closoure

package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()
	//a llama a la funcion clousure que mantiene el valor
	//de x
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	//b llama a la funcion clousure que mantiene el valor
	//de x
	fmt.Println(b())
	fmt.Println(b())

	//imprime 1,2,3,1,2
}
func incrementor() func() int {
	var x int
	//funcion anonima que mantiene una copia del entorno
	//las variables persisten en cada llamada
	return func() int { //closoure
		x++
		return x
	}
}
