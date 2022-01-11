//punteros--> direccion en memoria de un valor almacenado en la memoria RAM
//punteros es una dirección de memoria
package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) //imprime la diracción de a

	//tipo de a
	fmt.Printf("%T\n", a)  //imprime el tipo de a
	fmt.Printf("%T\n", &a) //indica que es un puntero a memoria almacenado un int

	var b *int = &a
	fmt.Println(b) //la direccion de memoria de a. es decir b apunta a a

	fmt.Println(*b) //accede a la direccion de b, pero obtiene el valor
	*b = 43
	fmt.Println(*b)
	fmt.Println(a) //el valor de a tambien se ha modificado
}
