//condicional Switch
package main

import "fmt"

func main() {
	switch exp {
	case 2 == 3:
		//primer caso
		fmt.Println("")
	case 3 == 3:
		//segundo caso
		fmt.Println("")
	case 4 == 5:
		//tercer caso
		fmt.Println("")
		fallthrough
	default:
		//condicion por defecto
		fmt.Println("Imprimiendo desde Default")
	}
}
