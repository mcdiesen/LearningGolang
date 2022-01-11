//añadiendo a un slice
package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 39}
	y := []int{333, 444, 555}
	fmt.Println(x)

	//append permite añadir elementos al slice
	x = append(x, 44, 55, 66)
	fmt.Println(x)

	//otra forma de añador elementos
	x = append(x, y...)
	fmt.Println(x)
	//otra forma de agr
	fmt.Println()

}
