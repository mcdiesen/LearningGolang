//Iterar sobre el slice - range
//range devuelve indice y valor del slice
package main

import "fmt"

func main() {
	x := []int{2, 4, 39}
	fmt.Println(x)
	fmt.Println("La capacidad del slice: ", cap(x))
	//numero de elementos
	//fmt.Println(x[0])
	//fmt.Println(x[1])
	//fmt.Println(x[2])
	//fmt.Println(x[3])
	fmt.Println("Indice-valor")
	for i, v := range x {
		fmt.Println(i, v)
	}
	//dividiendo el slice z[1:4]
	for i := 0; i < 4; i++ {
		fmt.Println(i, x[1])
	}
}
