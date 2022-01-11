//funcion con slices -- suma de sus elementos
package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(ii...)
	fmt.Println("foo - La suma de los elementos del slice es: ", n)

	//envia a la funcion el slice entero
	n2 := bar(ii)
	fmt.Println("bar: ", n2)

}
func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

//funcion que recibe los elementos el slice
func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
