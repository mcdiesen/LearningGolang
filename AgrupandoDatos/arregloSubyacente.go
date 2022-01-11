//arreglos subyacentes
package main

import "fmt"

func main() {
	//append
	x := []int(1, 2, 3, 4, 5)
	fmt.Println(x)

	//append para agregar nuevo slice
	y := append(x[:2], x[3:]...) //un nuevo arreglo es asignado
	fmt.Println(y)

}
