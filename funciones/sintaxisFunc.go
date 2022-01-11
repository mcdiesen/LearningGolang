//go programming specification
package main

import "fmt"

func main() {
	foo()
	bar("James Bond")
	woo("MoneyPenny")
	x, y := saludar("Charles ", "McLean")
	fmt.Println(x)
	fmt.Println(y)
}

//forma básica de una funcion
//no recibe parametros y no retorna nada
func foo() {
	fmt.Println("Hola desde foo")
}

func bar(s string) {
	fmt.Println("Hola ", s)
}
func woo(st string) string {
	return fmt.Sprint("Hola desde woo, ", st)
}
func saludar(n, a string) (string, bool) {
	p := fmt.Sprint(n, " ", a, " ", `dice "hola"`)
	q := true
	return p, q
}

//func (r receptor) identificador (parametros) retorno(s){codigo} <-- firma
//de la funcion
