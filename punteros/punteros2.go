//uso de punteros

package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x antes", x)
	fmt.Println("x antes", &x)
	foo(&x)
	fmt.Println("x despues", x)
	fmt.Println("x despues", &x)
}

func foo(y *int) {
	fmt.Println("Y antes", y)
	fmt.Println("Y antes", *y)
	*y = 42 //modificará el valor de la variable X
	fmt.Println("Y despues", y)
	fmt.Println("Y despues", *y)
}
