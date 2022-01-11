//marshall
//asignar los datos de struct a formato json

package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	//para asignarlo a JSON los identificadoes deben iniciar en mayuscula
	Nombre   string
	Apellido string
	Edad     int
}

func main() {

	p1 := persona{
		Nombre:   "James",
		Apellido: "Bond",
		Edad:     32,
	}

	p2 := persona{
		Nombre:   "MoneyPenny",
		Apellido: "Bons",
		Edad:     30,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	personas := []persona{
		p1,
		p2,
	}
	fmt.Println(personas)
	//pasarlo a JSON

	bs, err := json.Marshal(personas)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(string(bs))

}
