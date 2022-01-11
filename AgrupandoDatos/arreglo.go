//Arreglo: un bloque constructor
package main

import "fmt"

func main() {
	//crear arreglo
	var x [5]int
	x[3] = 39
	//imprime los elamentos del arreglo
	fmt.Println(x)
	//imprime el tipo de dato del arreglo
	fmt.Printf("%T\n", x)

	fmt.Println("La longitud del arreglo: ", len(x))
}
