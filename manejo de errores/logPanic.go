//log.Panic
//intentará abrir un archivo que no existe
package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("sin_archivo.txt")

	if err != nil {
		log.Panic(err)
		//Panic(err)
	}

}
