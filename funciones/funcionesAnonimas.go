//funciones anonimas
package main

import "fmt"

func main() {
	foo()
	func(x int) {
		fmt.Println("El significado de la vida es:", x)
	}(42) //funcion autoejecutando
}
func foo() {
	fmt.Println("foo se ejecuto")
}
