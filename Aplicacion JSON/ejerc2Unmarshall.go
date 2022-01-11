//Unmarshall

package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre string `json:"Nombre,omitempty"`
	Edad   int    `json:"Edad,omitempty"`
	Dichos string `json:"Dichos,omitempty"`
}

func main() {
	s := `[{"Nombre":"James","Edad":32},{"Nombre":"MoneyPenny","Edad":30},{"Dichos":"La mamá de quien"}]`
	//variables de tipo personas
	var personas []Persona
	err := json.Unmarshal([]byte(s), &personas)
	if err != nil {
		fmt.Println(err)
	}
	//imprimiria en formato struct
	fmt.Println(personas)
	for i, persona := range personas {
		fmt.Println("\nPersona #", i+1)
		fmt.Println(persona.Nombre, persona.Edad)
	}
	for _, dicho := range personas {
		fmt.Println("\t\t", dicho)
	}
}
