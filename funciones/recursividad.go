//funciones recursivas
//usar con cuidado la recursividad
//calculo del factorial de un número 4! = 4*3!...
package main

import "fmt"

func main() {
	f := factorial(5)
	fmt.Println(f)

	f2 := cicloFact(5)
	fmt.Println(f2)
}

//calculo del factorial utilizando funcion recursiva
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//calculo del factorial utilizando ciclo for
func cicloFact(i int) int {
	total := 1
	for ; i > 0; i-- {
		total *= i
	}
	return total
}
