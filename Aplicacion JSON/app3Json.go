//unmarshall
//herramientas: rawgit.com    json-to-go
//json: [{"Nombre":"James","Apellido":"Bond","Edad":32},{"Nombre":"MoneyPenny","Apellido":"Bons","Edad":30}]
//a tipo struct
package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

func main() {
	s := `[{"Nombre":"James","Apellido":"Bond","Edad":32},{"Nombre":"MoneyPenny","Apellido":"Bons","Edad":30}]`
	bs := []byte(s)

	//impresion de los tipos de datos de s y bs
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var personas []persona

	err := json.Unmarshal(bs, &personas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Toda la data:", personas)

	//impresion de la data
	for i, v := range personas {
		fmt.Println("\nPersona número:", i+1)
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
	}
}
