//defer statement
//invoca a una funcion cuya ejecución
//cuando llega al final de su cuerpo
package main

import "fmt"

func main() {
	defer foo() //imprimira primero bar
	bar()
}
func foo() {
	fmt.Println("foo")
}
func bar() {
	fmt.Println("bar")
}
