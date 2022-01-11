//sort -- ordenar slice, colecciones definidas por el usuario

package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 99, 88, 16, 56, 1} //desordenados
	xs := []string{"McLean", "Charles", "Dr.", "Golang dev"}
	sort.Ints(xi)   //ordena el slice
	fmt.Println(xi) //impresion de los valores del slice

	sort.Strings(xs) //ordenar el slice de string
	fmt.Println(xs)

}
