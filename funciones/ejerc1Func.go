//crea dos funciones que retornan valores
package main

import "fmt"

func main() {
	n := foo()
	//bar retorna dos valores
	x, s := bar()
	fmt.Println(n, x, s)
}

func foo() int {
	return 42
}
func bar() (int, string) {
	return 1492, "Descubrimiento de américa"
}
