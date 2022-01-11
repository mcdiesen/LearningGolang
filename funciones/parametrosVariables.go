//operadores variables ...
package main

import "fmt"

func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("El valor total almacenado en la variable es", x)
}

//recibe una cantidad variable de parámetros
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	suma := 0
	for i, v := range x {
		suma += v
		fmt.Println("El valor en el indice", i, "suma", v, "al total, quedando", suma)
	}
	fmt.Println("El total es: ", suma)
	return suma
}

//firma de la funcion
//func (r receptor) identificador (parametros) (retornos(s)){codigo}
