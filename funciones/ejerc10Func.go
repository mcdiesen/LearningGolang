//closoure
//encerar el scope de una variable en un bloque de código
package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f()) //copia de la variable x mantiene el valor
	//de la variable (entorno)
	fmt.Println(f())
	fmt.Println(f())
	g := foo() //copia de la variable x mantiene su valor en la variable g
	fmt.Println(g())
	fmt.Println(g())

}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
