//Control de flujo
// gobyexamples --> ejemplos de Go
package main

import "fmt"

func main() {
	//ciclos anidados
	for i := 0; i <= 100; i++ {
		fmt.Printf("El ciclo externo \n:", i)
		for j := 0; j <= 3; j++ {
			fmt.Printf("el ciclo interno\n", j)
		}
	}
}
