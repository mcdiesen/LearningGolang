//Sort ordenar slice de int y slice de string

package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"Una", "Persona", "que", "nunca", "ha", "comido", "un", "error", "nunca", "intento"}

	//impresion sin haber ordenado los elementos de los slices
	fmt.Println(xi)
	fmt.Println(xs)

	//impresion del slice ordenado
	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)
}
