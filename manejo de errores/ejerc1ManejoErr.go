//chequeo y manejo de error
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type persona struct {
	Nombre   string
	Apellido string
	Frases   []string
}

func main() {
	p1 := persona{
		Nombre:   "Charles",
		Apellido: "McLean",
		Frases:   []string{"tu puedes", "todo sacrificio tiene recomensa", "la mama de quien"},
	}
	bs, err := json.Marshal(p1)
	//Manejo de errores en el programa
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}
