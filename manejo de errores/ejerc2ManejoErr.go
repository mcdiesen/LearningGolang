//errorf
package main

import (
	"encoding/json"
	"errors"
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

	bs, err := aJSON(p1)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}

//aJSON tambien necesita retornar un error
func aJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)

	if err != nil {
		return []byte{}, errors.New(fmt.Sprint("Hubo un error en aJSON %v", err))
	}
	return bs, nil
}
