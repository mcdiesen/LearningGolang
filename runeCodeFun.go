package main

import "fmt"

//imprime el código ascii
//ciclos anidados
func main() {

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t %#U\n", i)
		}
	}
}
