//callback pasando una funcion como argumento
//callback --> programacion funcional - no recomendado en Golang
package main

import "fmt"

func main() {
	//lista de numeros enteros
	//que sean sumados a través de una función
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := suma(ii...)
	fmt.Println("La suma es total del slice es: ", s)
	s1 := sumaPares(suma, ii...)
	fmt.Println("El total de pares es: ", s1)
	s2 := sumaImpares(suma, ii...)
	fmt.Println("La suma de los impares es: ", s2)

}
func suma(xi ...int) int {
	//fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

//funcion que suma los numeros pares
func sumaPares(f func(xi ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}

//funcion que suma los numeros impares
func sumaImpares(f func(xi ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 != 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}
