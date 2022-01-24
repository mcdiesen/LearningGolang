//fatalln
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()

	_, err := os.Open("sin_archivo.txt")

	if err != nil {
		log.Fatal(err) //log por pantalla
		//panic(err)
	}
}
func foo() {
	fmt.Println("Cuando os.Exit() es llamada, las funciones diferidas no corren")
}

/*
... las funciones de Fatal llaman a os.Exit() despues de escribir el mensaje del log ...
*/
