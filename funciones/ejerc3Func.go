//funcion defer
package main

import "fmt"

func main() {

	defer foo()
	fmt.Println("Hola mundo")
}
func foo() {
	//funcion aninima con defer
	defer func() {
		fmt.Println("foo diferida corrio")
	}()
	fmt.Println("imprime Defer")
}
