//uso de operadores
package main

import "fmt"

//declaracion de constantes
const (
	//declaracion de una constante sin tipo
	x = 42
	//declaracion de una constante con tipo
	z int = 43
)

func main() {
	a := (42 == 42)
	b := (42 <= 42)
	c := (42 >= 42)
	d := (42 != 42)
	e := (42 < 42)
	f := (42 > 42)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(x, z)

}
