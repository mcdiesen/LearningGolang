//asignar funcion a una variable

package main

import "fmt"

var x int = 7 //declaracion de la variable y asignación
var g func() = func() {
	fmt.Println("g desde afuera")
}

func main() {

	f := func() {
		for i := 0; i <= 2; i++ {
			fmt.Println(i)
		}
	}
	f()

	fmt.Printf("%T\n", f) //imprime el tipo de la variable f
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	g()
	fmt.Printf("%T\n", g)

}
