//Logprintln

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("sin_archivo.txt")

	if err != nil {
		log.Println("Un error ocurrio", err)
	}
	defer f2.Close()

	fmt.Println("Revisa el archivo log.txt en el directorio.")
}
